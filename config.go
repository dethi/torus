package main

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Config struct {
	DatabasePath  string `toml:"database_path"`
	DataPath      string `toml:"data_path"`
	HtpasswdPath  string `toml:"htpasswd_path"`
	DownloadToken uint   `toml:"download_token"`
	TorrentPort   uint   `toml:"torrent_port"`
	WebPort       uint   `toml:"web_port"`

	Mailgun MailgunConfig `toml:"mailgun"`
}

type MailgunConfig struct {
	Domain    string `toml:"domain"`
	SecretKey string `toml:"secret_key"`
	PublicKey string `toml:"public_key"`
	Email     string `toml:"email"`
}

var cfg Config

// Load TOML config
func LoadConfig(pathname string) error {
	if _, err := toml.DecodeFile(pathname, &cfg); err != nil {
		return errors.Wrap(err, "load config failed")
	}
	return nil
}
