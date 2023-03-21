package email_sender

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

func GetOtpBodyMessage(otpCode string, title string) ([]byte, error) {
	filePath := filepath.Join("template", "email", "opt.html")
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Title string
		Otp   string
	}{
		Title: title,
		Otp:   otpCode,
	})

	return body.Bytes(), nil
}
