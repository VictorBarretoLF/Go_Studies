package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
	"golang.org/x/crypto/bcrypt"
)

func Secret(user, realm string) string {
	fmt.Println("User:", user, "Realm:", realm)
	if user == "victor" {
		// Instead of returning a plain text password, we return a bcrypt hash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("teste"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
			return ""
		}
		return string(hashedPassword)
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	porta := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)

	http.HandleFunc("/", authenticator.Wrap(
		func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
			fmt.Print("User", &r.Username)
			http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
		},
	))

	fmt.Printf("Subindo servidor na porta %s!\n", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil))
}
