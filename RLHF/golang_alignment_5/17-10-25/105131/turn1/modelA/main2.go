// emailtemplate/emailtemplate.go

package emailtemplate

import (
	"fmt"
	"html/template"
	"io"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

// SendEmail renders an email template and sends it to the writer.
func SendEmail(w io.Writer, templateName string, data interface{}) error {
	t := templates.Lookup(templateName)
	if t == nil {
		return fmt.Errorf("template %s not found", templateName)
	}

	return t.Execute(w, data)
}
