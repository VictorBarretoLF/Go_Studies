package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type PontosDeAcesso struct {
	ID          string `json:"id"`
	Nome        string `json:"nome"`
	Descricao   string `json:"descricao"`
	Localizacao string `json:"localizacao"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
}

func connectDB() *pgx.Conn {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if port == "" {
		port = "5432"
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	router := gin.Default()
	router.GET("/pontosDeAcesso", func(ctx *gin.Context) {
		conn := connectDB()
		defer conn.Close(context.Background())
		// pontosDeAcesso := []PontosDeAcesso{
		// 	{
		// 		ID:          "1",
		// 		Nome:        "Entrada Principal",
		// 		Descricao:   "A entrada principal do evento",
		// 		Localizacao: "Portão 1, Lado Norte",
		// 	},
		// 	{
		// 		ID:          "2",
		// 		Nome:        "Saída de Emergência Sul",
		// 		Descricao:   "Saída de emergência no lado sul do evento",
		// 		Localizacao: "Sul, Perto do Estacionamento",
		// 	},
		// }

		// ctx.IndentedJSON(http.StatusOK, pontosDeAcesso)
	})

	router.Run("10.155.1.109:8080")
}
