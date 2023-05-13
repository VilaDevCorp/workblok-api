package auth

import (
	"context"
	"sensei/ent"
	"sensei/ent/user"
	"sensei/utils"

	"golang.org/x/crypto/bcrypt"
)

type Svc interface {
	Login(ctx context.Context, form LoginForm) (utils.HttpResponse, *string)
}

type Store struct {
	DB *ent.Client
}

func hashPassword(password string) (result []byte, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (s *Store) Login(ctx context.Context, form LoginForm) (utils.HttpResponse, *string) {
	if form.Username == "" || form.Password == "" {
		res := utils.BadRequest(form, nil)
		return res, nil
	}
	loginUser, err := s.DB.User.Query().Where(user.UsernameEQ(form.Username)).Only(ctx)
	if err != nil {
		res := utils.NotFoundEntity(form.Username)
		return res, nil
	}

	if !checkPassword(loginUser.Password, form.Password) {
		res := utils.Unauthorized("Incorrect password")
		return res, nil

	}
	csrfToken, err := utils.GenerateRandomToken(64)
	if err != nil {
		res := utils.InternalError(err)
		return res, nil
	}
	// hash csrf
	hashedCsrfToken, err := utils.HashAndSalt(csrfToken)
	if err != nil {
		res := utils.InternalError(err)
		return res, nil
	}

	tokenString, err := utils.GenerateJWT(loginUser.ID.String(), loginUser.Mail, loginUser.Username, csrfToken)
	if err != nil {
		res := utils.InternalError(err)
		return res, nil
	}
	result := utils.LoginResult{Csrf: hashedCsrfToken}
	res := utils.OkLogged(result)
	return res, &tokenString
}
