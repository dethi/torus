package mailgun

import mgclient "github.com/mailgun/mailgun-go"

type Config struct {
	Domain    string `toml:"domain"`
	SecretKey string `toml:"secret_key"`
	PublicKey string `toml:"public_key"`
	Email     string `toml:"email"`
}

type MailService struct {
	email string
	mg    mgclient.Mailgun
}

func New(cfg Config) *MailService {
	return &MailService{
		email: cfg.Email,
		mg:    mgclient.NewMailgun(cfg.Domain, cfg.SecretKey, cfg.PublicKey),
	}
}
