package auth

import (
	"context"
	"time"
	"workblok/customerrors"
	"workblok/ent"
	"workblok/ent/session"
	"workblok/ent/user"
	"workblok/schema"
	"workblok/svc/verificationCode"
	"workblok/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthSvc interface {
	Login(ctx context.Context, form LoginForm) (*LoginResult, error)
	SignUp(ctx context.Context, form SignUpForm) (*ent.User, error)
	RefreshToken(
		ctx context.Context,
		refreshToken string,
		userAgent string,
	) (authToken string, newRefreshToken string, newSessionId uuid.UUID, err error)
	Logout(
		ctx context.Context,
		refreshTokenID uuid.UUID,
	) error
}

type LoginResult struct {
	AuthJWTToken    string
	RefreshJWTToken string
	SessionId       uuid.UUID
}

type AuthSvcImpl struct {
	DB                  *ent.Client
	VerificationCodeSvc verificationCode.VerificationCodeSvc
}

func checkPassword(
	hashPassword, password string,
) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (s *AuthSvcImpl) Login(
	ctx context.Context,
	form LoginForm,
) (*LoginResult, error) {
	if form.Email == "" ||
		form.Password == "" {
		return nil, &customerrors.EmptyFormFieldsError{}
	}
	loginUser, err := s.DB.User.Query().
		Where(user.EmailEQ(form.Email)).
		Only(ctx)
	if err != nil {
		return nil, &customerrors.InvalidCredentialsError{}
	}

	if !checkPassword(
		loginUser.Password,
		form.Password,
	) {
		return nil, &customerrors.InvalidCredentialsError{}
	}
	if !loginUser.EmailValid {
		return nil, &customerrors.NotValidatedAccountError{}
	}
	clientTx, err := s.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}
	session, err := clientTx.Session.Create().
		SetUserID(loginUser.ID).
		Save(ctx)
	if err != nil {
		clientTx.Rollback()
		return nil, err
	}
	//authToken
	authTokenString, err := utils.GenerateJWT(
		loginUser.ID,
		loginUser.Email,
		loginUser.Username,
		time.Minute*10,
		form.RememberMe,
		session.ID,
		form.UserAgent,
	)
	if err != nil {
		clientTx.Rollback()
		return nil, err
	}

	maxAge := time.Minute * 30
	if form.RememberMe {
		maxAge = time.Hour * 24 * 30
	}
	//refreshToken
	refreshTokenString, err := utils.GenerateJWT(
		loginUser.ID,
		loginUser.Email,
		loginUser.Username,
		maxAge,
		form.RememberMe,
		session.ID,
		form.UserAgent,
	)
	if err != nil {
		clientTx.Rollback()
		return nil, err
	}
	err = clientTx.Commit()
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		AuthJWTToken:    authTokenString,
		RefreshJWTToken: refreshTokenString,
		SessionId:       session.ID,
	}, nil
}

func (s *AuthSvcImpl) SignUp(
	ctx context.Context,
	form SignUpForm,
) (createdUser *ent.User, err error) {
	if form.UserName == "" || form.Email == "" ||
		form.Password == "" {
		return nil, &customerrors.EmptyFormFieldsError{}
	}
	bytesPass, err := bcrypt.GenerateFromPassword(
		[]byte(form.Password),
		14,
	)

	if err != nil {
		return nil, err
	}

	alreadyExistUser, _ := s.DB.User.Query().
		Where(user.UsernameEQ(form.UserName)).
		First(ctx)
	if alreadyExistUser != nil {
		return nil, &customerrors.UsernameAlreadyInUseError{}
	}
	alreadyExistMail, _ := s.DB.User.Query().
		Where(user.EmailEQ(form.Email)).
		First(ctx)
	if alreadyExistMail != nil {
		return nil, &customerrors.EmailAlreadyInUseError{}
	}
	clientTx, err := s.DB.Tx(ctx)
	if err != nil {
		return nil, err
	}
	user, err := clientTx.User.Create().
		SetUsername(form.UserName).
		SetPassword(string(bytesPass[:])).
		SetEmail(form.Email).
		SetConfig(&schema.Config{}).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	err = s.VerificationCodeSvc.Create(
		ctx,
		verificationCode.CreateForm{
			Email: user.Email,
			Type:  utils.VALIDATION_TYPE,
		},
		clientTx,
	)
	if err != nil {
		clientTx.Rollback()
		return nil, err
	}
	err = clientTx.Commit()
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *AuthSvcImpl) RefreshToken(
	ctx context.Context,
	refreshToken string,
	userAgent string,
) (newAuthToken string, newRefreshToken string, newSessionId uuid.UUID, err error) {

	refreshTokenClaims, err := utils.ValidateToken(
		refreshToken,
	)
	if err != nil {
		return "", "", uuid.Nil, &customerrors.InvalidTokenError{}
	}

	if refreshTokenClaims.UserAgent != userAgent {
		s.DB.Session.DeleteOneID(refreshTokenClaims.SessionId).
			Exec(ctx)
		return "", "", uuid.Nil, &customerrors.UserAgentNotMatchError{}
	}

	clientTx, err := s.DB.Tx(ctx)
	if err != nil {
		return "", "", uuid.Nil, err
	}

	//We get the session object
	sessionObj, err := clientTx.Session.Query().
		Where(session.IDEQ(refreshTokenClaims.SessionId)).
		Only(ctx)

	//If err!= nil (session not found) or it was removed time ago, we remove all the sessions
	//(danger of stolen token )
	if err != nil || sessionObj.DeleteDate != nil &&
		sessionObj.DeleteDate.Before(time.Now().Add(-time.Second*10)) {
		clientTx.Rollback()
		//We remove all the sessions (hard delete) linked to the user
		//Posibility of a token stolen
		s.DB.Session.Delete().
			Where(session.HasUserWith(user.IDEQ(refreshTokenClaims.UserId))).
			Exec(ctx)
		return "", "", uuid.Nil, &customerrors.TokenAlreadyUsedError{}
	}
	// If the session has been deleted less than 10 seconds ago only conflict error
	//(the grace period to avoid conflicts at refreshing tokens)
	err = clientTx.Session.Update().
		Where(session.IDEQ(sessionObj.ID)).
		SetDeleteDate(time.Now()).
		Exec(ctx)

	if err != nil {
		clientTx.Rollback()
		return "", "", uuid.Nil, err
	}

	//We remove all the sessions that have been created more than 30 days ago (to reduce the size of the table)
	clientTx.Session.Delete().
		Where(session.And(session.HasUserWith(user.IDEQ(refreshTokenClaims.UserId))), session.CreationDateLT(time.Now().Add(-time.Hour*24*30))).
		Exec(ctx)

	authTokenString, err := utils.GenerateJWT(
		refreshTokenClaims.UserId,
		refreshTokenClaims.Mail,
		refreshTokenClaims.Username,
		time.Minute*10,
		refreshTokenClaims.RememberMe,
		uuid.UUID{},
		"",
	)
	if err != nil {
		return "", "", uuid.Nil, err
	}

	refreshTokenDbObj, err := clientTx.Session.Create().
		SetUserID(refreshTokenClaims.UserId).
		Save(ctx)

	if err != nil {
		clientTx.Rollback()
		return "", "", uuid.Nil, err
	}
	maxAge := time.Minute * 30
	//If the original refresh token was created with rememberMe,
	//the new refresh token will also be created with rememberMe
	if refreshTokenClaims.RememberMe {
		maxAge = time.Hour * 24 * 30
	}
	refreshTokenString, err := utils.GenerateJWT(
		refreshTokenClaims.UserId,
		refreshTokenClaims.Mail,
		refreshTokenClaims.Username,
		maxAge,
		refreshTokenClaims.RememberMe,
		refreshTokenDbObj.ID,
		userAgent,
	)
	if err != nil {
		clientTx.Rollback()
		return "", "", uuid.Nil, err
	}

	err = clientTx.Commit()
	if err != nil {
		return "", "", uuid.Nil, err
	}

	return authTokenString, refreshTokenString, refreshTokenDbObj.ID, nil
}

func (s *AuthSvcImpl) Logout(
	ctx context.Context,
	sessionId uuid.UUID,
) error {
	err := s.DB.Session.DeleteOneID(sessionId).
		Exec(ctx)
	return err
}
