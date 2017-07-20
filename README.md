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

Emails could have HTML and text version, when sending emails,
It will look up template `hello.html.tmpl` and layout `application.html.tmpl` from view paths, and render it as HTML version's content, and use template `hello.text.tmpl` and layout `application.text.tmpl` as text version's content.
If we haven't find the layout file, we will only render template's as the content, and if we don't have the template, we will just skip that version, for example, if `hello.text.tmpl` doesn't exist, we will only send the HTML version.

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

Templates and layouts are located in `app/views/mailers`, which could be customized it with Mailer's AssetFS.

```go
import "github.com/qor/assetfs"

func main() {
	assetFS := assetfs.AssetFS().NameSpace("mailer")
	assetFS.RegisterPath("mailers/views")

	Mailer := mailer.New(&mailer.Config{
		Sender: gomailer.New(&gomailer.Config{Sender: sender}),
		AssetFS: assetFS,
	})
}
```
