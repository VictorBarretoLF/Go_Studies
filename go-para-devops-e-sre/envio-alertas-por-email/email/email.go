package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendMail(to []string, subject string, servidor string, erro string, horario string, templatePath string) {
	from := "victorbarretolins@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		panic("GMAIL_PASSWORD NOT FOUND")
		os.Exit(1)
	}

	smtHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtHost)

	t, _ := template.ParseFiles(templatePath)

	var body bytes.Buffer

	mineHeaders := "MINE-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mineHeaders)))

	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  servidor,
		Error:   erro,
		Horario: horario,
	})

	err := smtp.SendMail(smtHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Printf("Erro ao enviar email: %s", err)
		os.Exit(0)
	}
	fmt.Println("Email enviado com sucesso!")

}
