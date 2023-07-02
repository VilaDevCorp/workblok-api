package verificationCode

import (
	"context"
	"fmt"
	"math/rand"
	"sensei/ent"
	"sensei/ent/predicate"
	"sensei/ent/user"
	"sensei/ent/verificationcode"
	"sensei/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Svc interface {
	Create(ctx context.Context, form CreateForm) (*ent.VerificationCode, error)
	UseCode(ctx context.Context, form UseForm) utils.HttpResponse
	Get(ctx context.Context, verificationCodeId uuid.UUID) (*ent.VerificationCode, error)
	Delete(ctx context.Context, verificationCodeId uuid.UUID) error
}

type Store struct {
	DB *ent.Client
}

func (s *Store) Create(ctx context.Context, form CreateForm) (*ent.VerificationCode, error) {
	expireDate := time.Now().Add(time.Minute * 15)
	fmt.Println(expireDate)
	code := rand.Intn((1000000))
	codeStr := fmt.Sprintf("%06d", code)
	userId, err := s.DB.User.Query().Where(user.MailEQ(form.Mail)).FirstID(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(userId)
	clientTx, err := s.DB.Tx(ctx)
	prevValidCodes, err := clientTx.VerificationCode.Query().Where(verificationcode.And(verificationcode.HasUserWith(user.ID(userId)), verificationcode.Valid(true))).All(ctx)
	for _, prevCode := range prevValidCodes {
		clientTx.VerificationCode.UpdateOneID(prevCode.ID).SetValid(false).Save(ctx)
	}
	if err != nil {
		clientTx.Rollback()
		return nil, err
	}
	verif, err := clientTx.VerificationCode.Create().SetCode(codeStr).SetUserID(userId).SetType(form.Type).SetExpireDate(expireDate).SetValid(true).Save(ctx)
	if err != nil {
		clientTx.Rollback()
		return nil, err
	}
	clientTx.Commit()
	return verif, err

}

func (s *Store) Get(ctx context.Context, verificationCodeId uuid.UUID) (*ent.VerificationCode, error) {
	return s.DB.VerificationCode.Get(ctx, verificationCodeId)
}

func (s *Store) UseCode(ctx context.Context, form UseForm) utils.HttpResponse {
	clientTx, err := s.DB.Tx(ctx)
	query := clientTx.VerificationCode.Query()
	var conditions []predicate.VerificationCode
	var userCondition []predicate.User
	userCondition = append(userCondition, user.Mail(form.Mail))
	conditions = append(conditions, verificationcode.Valid(true))
	conditions = append(conditions, verificationcode.HasUserWith(userCondition...))
	conditions = append(conditions, verificationcode.TypeEQ(form.Type))
	verificationCode, err := query.Where(verificationcode.And(conditions...)).First(ctx)
	if err != nil {
		return utils.NotFoundEntity(form.Mail)
	}
	if verificationCode.ExpireDate.Before(time.Now()) {
		return utils.ExpiredCode()
	}
	if form.Code != verificationCode.Code {
		return utils.VerificationCodeNotMatch()
	}
	if form.Type == utils.VALIDATION_TYPE {
		if err != nil {
			return utils.InternalError(err)
		}
		userId, err := clientTx.User.Query().Where(user.MailEQ(form.Mail)).FirstID(ctx)
		if err != nil {
			return utils.InternalError(err)
		}
		_, err = clientTx.User.UpdateOneID(userId).SetMailValid(true).Save(ctx)
		if err != nil {
			return utils.InternalError(err)
		}
		_, err = clientTx.VerificationCode.UpdateOneID(verificationCode.ID).SetValid(false).Save(ctx)
		if err != nil {
			clientTx.Rollback()
			return utils.InternalError(err)
		}
		clientTx.Commit()
	}
	if form.Type == utils.RECOVER_TYPE {
		if err != nil {
			return utils.InternalError(err)
		}
		userId, err := clientTx.User.Query().Where(user.MailEQ(form.Mail)).FirstID(ctx)
		if err != nil {
			return utils.InternalError(err)
		}
		fmt.Println(userId)
		bytesPass, err := bcrypt.GenerateFromPassword([]byte(form.NewPass), 14)
		if err != nil {
			return utils.InternalError(err)
		}
		fmt.Println(bytesPass)
		fmt.Println(string(bytesPass[:]))

		_, err = clientTx.User.UpdateOneID(userId).SetPassword(string(bytesPass[:])).Save(ctx)
		if err != nil {
			return utils.InternalError(err)
		}
		_, err = clientTx.VerificationCode.UpdateOneID(verificationCode.ID).SetValid(false).Save(ctx)
		if err != nil {
			clientTx.Rollback()
			return utils.InternalError(err)
		}
		clientTx.Commit()
	}
	return utils.OkOperation(nil)
}

func (s *Store) Delete(ctx context.Context, verificationCodeId uuid.UUID) error {
	err := s.DB.VerificationCode.DeleteOneID(verificationCodeId).Exec(ctx)
	return err
}
