package mailgun

import (
	"github.com/dethi/torus"
	"github.com/pkg/errors"
)

func (ms *MailService) SendMail(m torus.Mail) error {
	msg := ms.mg.NewMessage(ms.email, m.Title, m.Body, m.To)
	_, _, err := ms.mg.Send(msg)
	return errors.Wrapf(err, "send mail `%v` to %v", m.Title, m.To)
}
