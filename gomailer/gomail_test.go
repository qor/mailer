package gomailer_test

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"testing"

	gomail "gopkg.in/gomail.v2"

	"github.com/jinzhu/configor"
	"github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
)

var Mailer *mailer.Mailer

var Config = struct {
	Address  string `env:"SMTP_Address"`
	Port     int    `env:"SMTP_Port"`
	User     string `env:"SMTP_User"`
	Password string `env:"SMTP_Password"`
}{}

func init() {
	configor.Load(&Config)

	d := gomail.NewDialer(Config.Address, Config.Port, Config.User, Config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	dailer, err := d.Dial()
	if err != nil {
		panic(fmt.Sprintf("Got error %v when dail mail server: %#v", err, Config))
	}

	sender := gomailer.New(&gomailer.Config{Sender: dailer})

	Mailer = mailer.New(&mailer.Config{
		Sender: sender,
	})
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
