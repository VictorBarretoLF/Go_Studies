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
	TempoExecucao float64
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

	fmt.Println(data)

	agora := time.Now()
	url := os.Args[1]

	get, err := http.Get(url)
	if err != nil {
		fmt.Println("Ocorreu um erro ao executar o get(url)")
		panic(err)
	}

	decorrido := time.Since(agora).Seconds()
	status := get.StatusCode
	fmt.Printf("Status: [%d] Tempo de carga: [%f]\n", status, decorrido)
}
