package torus

import "time"

type Mail struct {
	Datetime time.Time
	To       string
	From     string
	Title    string
	Body     string
}

type MailService interface {
	SendMail(m Mail) error
	ReceiveMail() <-chan Mail
}
