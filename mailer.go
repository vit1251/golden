package main

import (
	"github.com/vit1251/golden/pkg/mailer"
)

func Mailer() {

//	m := mailer.NewMailer()
	m := mailer.NewMailerCompat()

	m.Check()

}
