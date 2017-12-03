package mailer

import (
	"github.com/pkg/errors"
	"github.com/qor/assetfs"
	"github.com/qor/render"
)

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
	*render.Render
}

// New initialize mailer
func New(config *Config) (*Mailer, error) {
	if config == nil {
		config = &Config{}
	}

	if config.AssetFS == nil {
		config.AssetFS = assetfs.AssetFS().NameSpace("mailer")
	}

	if err := config.AssetFS.RegisterPath("app/views/auth/mail"); err != nil {
		return nil, errors.Wrap(err, "Mailer: could not registerPath")
	}

	if config.Render == nil {
		config.Render = render.New(nil)
		config.Render.SetAssetFS(config.AssetFS)
	}

	return &Mailer{config}, nil
}

// Send send email
func (mailer Mailer) Send(email Email, templates ...Template) error {
	if mailer.DefaultEmailTemplate != nil {
		email = mailer.DefaultEmailTemplate.Merge(email)
	}

	if len(templates) == 0 {
		return mailer.Sender.Send(email)
	}

	for i, template := range templates {
		m, err := mailer.Render(template)
		if err != nil {
			return errors.Wrapf(err, "failed to render email template(%d): %s", i, template.Name)
		}
		if err := mailer.Sender.Send(m.Merge(email)); err != nil {
			return errors.Wrapf(err, "failed to send email template(%d): %s", i, template.Name)
		}
	}
	return nil
}
