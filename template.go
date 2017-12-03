package mailer

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Template email template
type Template struct {
	Name    string
	Layout  string
	Data    interface{}
	Request *http.Request
	Writer  http.ResponseWriter
	funcMap template.FuncMap
}

// Funcs set template's funcs
func (tmpl Template) Funcs(funcMap template.FuncMap) Template {
	tmpl.funcMap = funcMap
	return tmpl
}

// Render render template
func (mailer Mailer) Render(t Template) (*Email, error) {
	var email Email

	if t.funcMap == nil {
		t.funcMap = template.FuncMap{}
	}

	if _, ok := t.funcMap["root_url"]; !ok {
		t.funcMap["root_url"] = func() string {
			if t.Request != nil && t.Request.URL != nil {
				var newURL url.URL
				newURL.Host = t.Request.URL.Host
				newURL.Scheme = t.Request.URL.Scheme
				return newURL.String()
			}
			return ""
		}
	}

	if t.Layout != "" {
		result, err := mailer.Config.Render.Layout(t.Layout+".text").Funcs(t.funcMap).Render(t.Name+".text", t.Data, t.Request, t.Writer)
		if err != nil {
			return nil, errors.Wrap(err, "mail/render layout (txt) failed")
		}
		email.Text = string(result)

		result, err = mailer.Config.Render.Layout(t.Layout+".html").Funcs(t.funcMap).Render(t.Name+".html", t.Data, t.Request, t.Writer)
		if err != nil {
			return nil, errors.Wrap(err, "mail/render layout (html) failed")
		}
		email.HTML = string(result)

		if email.HTML == "" {
			result, err := mailer.Config.Render.Layout(t.Layout).Funcs(t.funcMap).Render(t.Name, t.Data, t.Request, t.Writer)
			if err != nil {
				return nil, errors.Wrap(err, "mail/render layout (fallback) failed")
			}
			email.HTML = string(result)
		}
	} else {
		result, err := mailer.Config.Render.Funcs(t.funcMap).Render(t.Name+".text", t.Data, t.Request, t.Writer)
		if err != nil {
			return nil, errors.Wrap(err, "mail/render (txt) failed")
		}
		email.Text = string(result)

		result, err = mailer.Config.Render.Funcs(t.funcMap).Render(t.Name+".html", t.Data, t.Request, t.Writer)
		if err != nil {
			return nil, errors.Wrap(err, "mail/render (html) failed")
		}
		email.HTML = string(result)

		if email.HTML == "" {
			result, err := mailer.Config.Render.Funcs(t.funcMap).Render(t.Name, t.Data, t.Request, t.Writer)
			if err != nil {
				return nil, errors.Wrap(err, "mail/render (fallback) failed")
			}
			email.HTML = string(result)
		}
	}

	return &email, nil
}
