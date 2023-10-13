package svc

import (
	"workblok/db"
	"workblok/svc/auth"
	"workblok/svc/block"
	"workblok/svc/user"
	"workblok/svc/verificationCode"
)

type Service struct {
	Block            block.Svc
	User             user.Svc
	Auth             auth.Svc
	VerificationCode verificationCode.Svc
}

var svc Service

func Get() *Service {
	return &svc
}

func Setup() {
	client := db.GetClient()
	svc = Service{
		Block:            &block.Store{DB: client},
		User:             &user.Store{DB: client},
		Auth:             &auth.Store{DB: client},
		VerificationCode: &verificationCode.Store{DB: client},
	}
}
