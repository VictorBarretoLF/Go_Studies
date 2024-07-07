package main

import "alertas-email/email"

func main() {
	email.SendMail([]string{"victorbarretolins@gmail.com"},
		"Servidor fora do ar",
		"Google",
		"Erro ao conectar no servidor",
		"06/20/2024 14:30",
		"./email/template.html",
	)
}
