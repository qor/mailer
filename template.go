package mailer

// Template email template
type Template struct {
	Name   string
	Layout string
	Data   interface{}
}

// Render render template
func (mailer Mailer) Render(template Template) Email {
	return Email{}
}
