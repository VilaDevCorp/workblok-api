package svc

import (
	"workblok/db"
	"workblok/svc/auth"
	"workblok/svc/block"
	"workblok/svc/user"
	"workblok/svc/verificationCode"
)

type Service struct {
	Block            block.BlockSvc
	User             user.UserSvc
	Auth             auth.AuthSvc
	VerificationCode verificationCode.VerificationCodeSvc
}

var svc Service

func Get() *Service {
	return &svc
}

func Setup() {
	client := db.GetClient()
	svc = Service{
		Block: &block.BlockSvcImpl{DB: client},
		User:  &user.UserSvcImpl{DB: client},
		Auth: &auth.AuthSvcImpl{
			DB:                  client,
			VerificationCodeSvc: &verificationCode.VerificationCodeSvcImpl{DB: client},
		},
		VerificationCode: &verificationCode.VerificationCodeSvcImpl{DB: client},
	}
}
