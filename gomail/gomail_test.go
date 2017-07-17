package gomail

import (
	"testing"

	"github.com/qor/mailer"
)

var Mailer *mailer.Mailer

func init() {
	sender := New(&Config{Sender: dailer})

	Mailer = mailer.New(&mailer.Config{
		Sender: sender,
	})
}

func TestSendEmail(t *testing.T) {
	t.Errorf("hello world")
}
