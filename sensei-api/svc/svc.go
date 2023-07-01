package svc

import (
	"sensei/db"
	"sensei/svc/activity"
	"sensei/svc/auth"
	"sensei/svc/task"
	"sensei/svc/template"
	"sensei/svc/user"
	"sensei/svc/verificationCode"
)

type Service struct {
	Activity         activity.Svc
	Task             task.Svc
	User             user.Svc
	Auth             auth.Svc
	VerificationCode verificationCode.Svc
	Template         template.Svc
}

var svc Service

func Get() *Service {
	return &svc
}

func Setup() {
	client := db.GetClient()
	svc = Service{
		Activity:         &activity.Store{DB: client},
		Task:             &task.Store{DB: client},
		User:             &user.Store{DB: client},
		Auth:             &auth.Store{DB: client},
		VerificationCode: &verificationCode.Store{DB: client},
		Template:         &template.Store{DB: client},
	}
}
