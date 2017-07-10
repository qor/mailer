package mailer

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
	Sender               SenderInterface
}

// New initialize mailer
func New(config *Config) *Mailer {
	if config == nil {
		config = &Config{}
	}

	return &Mailer{config}
}

// Send send email
func (mailer Mailer) Send(email Email) error {
	email = mailer.DefaultEmailTemplate.Merge(email)

	// render email template
	if email.Text == "" && email.HTML == "" && email.Template != "" {
		// render templates with funcs maps
	}

	return mailer.Sender.Send(email)
}
