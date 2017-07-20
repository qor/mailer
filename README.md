# Mailer

Mail solution

## Usage

### Initailize Mailer

Mailer will have multiple sender adaptors, here is how to use [gomail](https://github.com/go-gomail/gomail) to send emails

```go
dailer := gomail.NewDialer(Config.Address, Config.Port, Config.User, Config.Password)
sender, err := dailer.Dial()

Mailer := mailer.New(&mailer.Config{
	Sender: gomailer.New(&gomailer.Config{Sender: sender}),
})
```

### Sending Emails

```go
Mailer.Send(mailer.Email{
	TO:          []mail.Address{{Address: "jinzhu@example.org", Name: "jinzhu"}},
	From:        &mail.Address{Address: "jinzhu@example.org"},
	Subject: "subject",
	Text:        "text email",
	HTML:        "html email <img src='cid:logo.png'/>",
	Attachments: []mailer.Attachment{{FileName: "gomail.go"}, {FileName: "../test/logo.png", Inline: true}},
})
```

### Sending Emails with templates

Mailer is using [Render](github.com/qor/render) to render email templates and layouts, please refer it for How-To.

Emails could have HTML and text version, it will use template `hello.html.tmpl` with layout `application.html.tmpl` as the HTML version, `hello.text.tmpl` with layout `application.text.tmpl` as the text version, if we don't have the layout file, we will use template's content for the email, if we don't have the template, we will skip that email version.

Templates and layouts are located in `app/views/mailers`, which could be customized it with Mailer's AssetFS.

```go
Mailer.Send(
	mailer.Email{
		TO:      []mail.Address{{Address: Config.DefaultTo}},
		From:    &mail.Address{Address: Config.DefaultFrom},
		Subject: "hello",
	},
	mailer.Template{Name: "hello", Layout: "application", Data: currentUser},
)
```
