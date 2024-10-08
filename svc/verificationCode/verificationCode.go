package verificationCode

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"workblok/conf"
	"workblok/customerrors"
	"workblok/ent"
	"workblok/ent/predicate"
	"workblok/ent/user"
	"workblok/ent/verificationcode"
	"workblok/svc/mail"
	"workblok/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type VerificationCodeSvc interface {
	Create(
		ctx context.Context,
		form CreateForm,
		transaction *ent.Tx,
	) error
	UseCode(
		ctx context.Context,
		form UseForm,
	) error
	Get(
		ctx context.Context,
		verificationCodeId uuid.UUID,
	) (*ent.VerificationCode, error)
	Delete(
		ctx context.Context,
		verificationCodeId uuid.UUID,
	) error
}

type VerificationCodeSvcImpl struct {
	DB   *ent.Client
	Mail mail.MailSvc
}

func (s *VerificationCodeSvcImpl) Create(
	ctx context.Context,
	form CreateForm,
	transaction *ent.Tx,
) error {
	expireDate := time.Now().Add(time.Minute * 15)
	code := rand.Intn((1000000))
	codeStr := fmt.Sprintf("%06d", code)
	// If transaction is not nil, it means that the transaction is being managed by another function, and we should not commit or rollback it
	var clientTx *ent.Tx
	var externalTx bool
	var err error

	if transaction != nil {
		clientTx = transaction
		externalTx = true
	} else {
		clientTx, err = s.DB.Tx(ctx)
		if err != nil {
			return err
		}
	}

	user, err := clientTx.User.Query().
		Where(user.EmailEQ(form.Email)).
		First(ctx)
	if err != nil {
		//If the user is not found, we should rollback the transaction and return no error to avoid username enumeration
		if !externalTx {
			clientTx.Rollback()
		}
		return nil
	}

	verificationCode, err := clientTx.VerificationCode.Create().
		SetCode(codeStr).
		SetUserID(user.ID).
		SetType(form.Type).
		SetExpireDate(expireDate).
		Save(ctx)
	if err != nil {
		if !externalTx {
			clientTx.Rollback()
		}
		return err
	}
	hostUrl := conf.Get().Dev.FrontUrl
	if conf.Get().Env == "prod" {
		hostUrl = conf.Get().Prod.FrontUrl
	}
	if conf.Get().Mail.Enabled {
		var mailSubject, mailBody string
		if form.Type == utils.VALIDATION_TYPE {
			mailSubject = conf.Get().Mail.MailTypes.AccountActivation.Subject
			mailBody = conf.Get().Mail.MailTypes.AccountActivation.Body
		}
		if form.Type == utils.RECOVER_TYPE {
			mailSubject = conf.Get().Mail.MailTypes.PasswordReset.Subject
			mailBody = conf.Get().Mail.MailTypes.PasswordReset.Body

		}
		err = s.Mail.SendMail(
			user.Email,
			mailSubject,
			fmt.Sprintf(
				mailBody,
				hostUrl,
				user.Username,
				verificationCode.Code,
			),
		)

		if err != nil {
			if !externalTx {
				clientTx.Rollback()
			}
			return err
		}
	}
	if !externalTx {
		clientTx.Commit()
	}
	return nil
}

func (s *VerificationCodeSvcImpl) Get(
	ctx context.Context,
	verificationCodeId uuid.UUID,
) (*ent.VerificationCode, error) {
	verifCode, err := s.DB.VerificationCode.Get(
		ctx,
		verificationCodeId,
	)
	if err != nil {
		return nil, &customerrors.NotFoundError{
			Resource: "Verification code: " + verificationCodeId.String(),
		}
	}
	return verifCode, nil
}

func (s *VerificationCodeSvcImpl) validateAccount(
	clientTx *ent.Tx,
	ctx context.Context,
	email string,
) error {
	userId, err := clientTx.User.Query().
		Where(user.EmailEQ(email)).
		FirstID(ctx)
	if err != nil {
		return &customerrors.NotFoundError{
			Resource: "User: " + email,
		}
	}
	_, err = clientTx.User.UpdateOneID(userId).
		SetEmailValid(true).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *VerificationCodeSvcImpl) resetPassword(
	clientTx *ent.Tx,
	ctx context.Context,
	email string,
	newPass string,
) error {
	userId, err := clientTx.User.Query().
		Where(user.EmailEQ(email)).
		FirstID(ctx)
	if err != nil {
		return &customerrors.NotFoundError{
			Resource: "User: " + email,
		}
	}
	bytesPass, err := bcrypt.GenerateFromPassword(
		[]byte(newPass),
		14,
	)
	if err != nil {
		return err
	}
	_, err = clientTx.User.UpdateOneID(userId).
		SetPassword(string(bytesPass[:])).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil

}

func (s *VerificationCodeSvcImpl) UseCode(
	ctx context.Context,
	form UseForm,
) error {
	clientTx, err := s.DB.Tx(ctx)
	if err != nil {
		utils.InternalError(err)
	}
	query := clientTx.VerificationCode.Query()
	var conditions []predicate.VerificationCode
	var userCondition []predicate.User
	userCondition = append(
		userCondition,
		user.Email(form.Email),
	)
	conditions = append(
		conditions,
		verificationcode.HasUserWith(
			userCondition...),
	)
	conditions = append(
		conditions,
		verificationcode.TypeEQ(form.Type),
	)
	verificationCode, err := query.Where(verificationcode.And(conditions...)).
		Order(ent.Desc(verificationcode.FieldCreationDate)).
		First(ctx)
	if err != nil {
		clientTx.Rollback()
		return &customerrors.NotFoundError{
			Resource: "Verification code for user: " + form.Email,
		}
	}
	if verificationCode.Code == form.Code &&
		!verificationCode.Valid {
		clientTx.Rollback()
		return &customerrors.AlreadyUsedValidationCodeError{}
	}
	if verificationCode.Code == form.Code &&
		verificationCode.ExpireDate.Before(
			time.Now(),
		) {
		clientTx.Rollback()
		return &customerrors.ExpiredValidationCodeError{}
	}
	if form.Code != verificationCode.Code {
		clientTx.Rollback()
		return &customerrors.IncorrectValidationCodeError{}
	}
	if form.Type == utils.VALIDATION_TYPE {
		err = s.validateAccount(
			clientTx,
			ctx,
			form.Email,
		)
		if err != nil {
			clientTx.Rollback()
			return err
		}
	}
	if form.Type == utils.RECOVER_TYPE {
		err = s.resetPassword(
			clientTx,
			ctx,
			form.Email,
			form.NewPass,
		)
		if err != nil {
			clientTx.Rollback()
			return err
		}
	}
	_, err = clientTx.VerificationCode.UpdateOneID(verificationCode.ID).
		SetValid(false).
		Save(ctx)
	if err != nil {
		clientTx.Rollback()
		return err
	}
	clientTx.Commit()
	return nil
}

func (s *VerificationCodeSvcImpl) Delete(
	ctx context.Context,
	verificationCodeId uuid.UUID,
) error {
	verificationCode := s.DB.VerificationCode.Query().
		Where(verificationcode.IDEQ(verificationCodeId)).
		FirstX(ctx)
	if verificationCode == nil {
		return &customerrors.NotFoundError{
			Resource: "Verification code: " + verificationCodeId.String(),
		}
	}
	err := s.DB.VerificationCode.DeleteOneID(verificationCodeId).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
