package mailer

import (
	"io"
	"net/mail"
)

// Email email struct
type Email struct {
	TO, CC, BCC   []mail.Address
	From, ReplyTo *mail.Address
	Subject       string
	Headers       mail.Header
	Attachments   []Attachment
	Template      string // template name
	Layout        string // application
}

// Attachment attachment struct
type Attachment struct {
	FileName string
	Inline   bool
	MimeType string
	Content  io.Reader
}

func (Email) Merge(Email) Email {
	return Email{}
}
