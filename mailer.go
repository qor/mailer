package mailer

import "github.com/qor/assetfs"

// SenderInterface sender's interface
type SenderInterface interface {
	Send(Email) error
}

// Mailer mailer struct
type Mailer struct {
	*Config
}

// Config mailer config
type Config struct {
	DefaultEmailTemplate *Email
	AssetFS              assetfs.Interface
	Sender               SenderInterface
}

// New initialize mailer
func New(config *Config) *Mailer {
	if config == nil {
		config = &Config{}
	}

	if config.AssetFS == nil {
		config.AssetFS = assetfs.AssetFS().NameSpace("mailer")
	}

	config.AssetFS.RegisterPath("app/views/mailers")

	return &Mailer{config}
}

// Send send email
func (mailer Mailer) Send(email Email, templates ...Template) error {
	email = mailer.DefaultEmailTemplate.Merge(email)

	if len(templates) == 0 {
		return mailer.Sender.Send(email)
	}

	for _, template := range templates {
		return mailer.Sender.Send(mailer.Render(template).Merge(email))
	}
}
