package email_sender

import (
	"bytes"
	"fmt"
	"html/template"
)

func GetOptBodyMessage(optCode string, title string) ([]byte, error) {
	t, err := template.ParseFiles("template/email/opt.html")
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Title string
		Opt   string
	}{
		Title: title,
		Opt:   optCode,
	})

	return body.Bytes(), nil
}
