package mailer

import (
	"io"
	"net/mail"
)

type Interface interface {
	Send(Email) error
}

type Mailer struct {
	*Config
}

type Config struct {
	DefaultFrom    *mail.Address
	DefaultReplyTo *mail.Address
	DefaultTo      []mail.Address
	Sender         Interface
}

func New(config *Config) *Mailer {
	if config == nil {
		config = &Config{}
	}

	return &Mailer{config}
}

func (mailer Mailer) Send(email Email) error {
	if len(email.TO) == 0 {
		email.TO = mailer.Config.DefaultTo
	}

	if email.From == nil {
		email.From = mailer.Config.DefaultFrom
	}

	if email.ReplyTo == nil {
		email.ReplyTo = mailer.Config.DefaultReplyTo
	}

	return mailer.Sender.Send(email)
}

type Email struct {
	TO, CC, BCC   []mail.Address
	From, ReplyTo *mail.Address
	Subject       string
	Headers       mail.Header
	Attachments   []Attachment
	Template      string // template name
}

type Attachment struct {
	FileName string
	Inline   bool
	MimeType string
	Content  io.Reader
}
