package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	ServerName    string
	ServerURL     string
	tempoExecucao float64
	status        int
}

func criarListaServidores(data [][]string) []Server {
	var servidores []Server

	for i, line := range data {
		if i > 0 {
			servidor := Server{
				ServerName: line[0],
				ServerURL:  line[1],
			}
			servidores = append(servidores, servidor)
		}
	}

	return servidores
}

func checkServers(servidores []Server) {
	for _, servidor := range servidores {
		agora := time.Now()
		resp, err := http.Get(servidor.ServerURL)
		if err != nil {
			fmt.Println(err)
		}
		servidor.status = resp.StatusCode
		servidor.tempoExecucao = time.Since(agora).Seconds()

		fmt.Printf("Status: [%d] Tempo de carga: [%f] Server: [%s]\n", servidor.status, servidor.tempoExecucao, servidor.ServerURL)

		defer resp.Body.Close()
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	servidores := criarListaServidores(data)
	checkServers(servidores)

}
