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
	SendRealEmail bool   `env:"DEBUG"`
	Address       string `env:"SMTP_Address"`
	Port          int    `env:"SMTP_Port"`
	User          string `env:"SMTP_User"`
	Password      string `env:"SMTP_Password"`
	DefaultTo     string `env:"SMTP_TO" default:"jinzhu@example.org"`
	DefaultFrom   string `env:"SMTP_From" default:"from@example.org"`
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
			Box.WriteString(fmt.Sprintf("From: %v\n", from))
			Box.WriteString(fmt.Sprintf("To: %v\n", strings.Join(to, ", ")))
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
		TO:          []mail.Address{{Address: Config.DefaultTo}},
		From:        &mail.Address{Address: Config.DefaultFrom},
		Text:        "text email",
		HTML:        "html email <img src='cid:logo.png'/>",
		Attachments: []mailer.Attachment{{FileName: "gomail.go"}, {FileName: "../test/logo.png", Inline: true}},
	})

	if err != nil {
		t.Errorf("No error should raise when send email")
	}

	fmt.Println(Box.String())
}
