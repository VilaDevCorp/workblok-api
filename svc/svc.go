package svc

import (
	"appname/db"
	"appname/svc/activity"
	"appname/svc/auth"
	"appname/svc/task"
	"appname/svc/template"
	"appname/svc/user"
	"appname/svc/verificationCode"
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
