package main

import "alertas-email/slack"

func main() {
	slack.SendSlack("Alerta de servidor down eita!!!")

	// email.SendMail([]string{"victorbarretolins@gmail.com"},
	// 	"Servidor fora do ar",
	// 	"Google",
	// 	"Erro ao conectar no servidor",
	// 	"06/20/2024 14:30",
	// 	"./email/template.html",
	// )

}
