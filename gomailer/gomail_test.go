package gomailer_test

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/mail"
	"strings"
	"testing"

	gomail "gopkg.in/gomail.v2"

	"github.com/jinzhu/configor"
	"github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
)

var Mailer *mailer.Mailer

var Config = struct {
	SendRealEmail bool
	Address       string `env:"SMTP_Address"`
	Port          int    `env:"SMTP_Port"`
	User          string `env:"SMTP_User"`
	Password      string `env:"SMTP_Password"`
}{}

var Box bytes.Buffer

func init() {
	configor.Load(&Config)

	if Config.SendRealEmail {
		d := gomail.NewDialer(Config.Address, Config.Port, Config.User, Config.Password)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		sender, err := d.Dial()
		if err != nil {
			panic(fmt.Sprintf("Got error %v when dail mail server: %#v", err, Config))
		}

		Mailer = mailer.New(&mailer.Config{
			Sender: gomailer.New(&gomailer.Config{Sender: sender}),
		})
	} else {
		sender := gomail.SendFunc(gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
			Box.WriteString("From: " + from)
			Box.WriteString("To: " + strings.Join(to, ", "))
			_, err := msg.WriteTo(&Box)
			return err
		}))

		Mailer = mailer.New(&mailer.Config{
			Sender: gomailer.New(&gomailer.Config{Sender: sender}),
		})
	}
}

func TestSendEmail(t *testing.T) {
	err := Mailer.Send(mailer.Email{
		TO:   []mail.Address{{Address: "jinzhu@example.org"}},
		From: &mail.Address{Address: "from@example.org"},
		Text: "text email",
	})

	if err != nil {
		t.Errorf("No error should raise when send email")
	}
	t.Errorf("hello world")
}
