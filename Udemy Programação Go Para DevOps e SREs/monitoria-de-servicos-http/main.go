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
	dataFalha     string
}

const (
	LISTA_SERVIDORES_CSV_LOCALIZACAO          = "lista-servidores.csv"
	LISTA_SERVIDORES_DOWNTIME_CSV_LOCALIZACAO = "lista-servidores-downtime.csv"
)

func criarListaServidores(serverList *os.File) []Server {
	csvReader := csv.NewReader(serverList)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

func checkServers(servidores []Server) []Server {
	var downServers []Server

	for _, servidor := range servidores {
		agora := time.Now()
		resp, err := http.Get(servidor.ServerURL)

		if err != nil {
			fmt.Printf("Server %s is down %s\n", servidor.ServerName, err.Error())
			servidor.status = 0
			servidor.dataFalha = agora.Format("02/01/2006 15:20:30")
			downServers = append(downServers, servidor)
			continue
		}
		servidor.status = resp.StatusCode

		if servidor.status != 200 {
			servidor.dataFalha = agora.Format("02/01/2006 15:20:30")
			downServers = append(downServers, servidor)
		}

		servidor.tempoExecucao = time.Since(agora).Seconds()
		fmt.Printf("Status: [%d] Tempo de carga: [%f] Server: [%s]\n", servidor.status, servidor.tempoExecucao, servidor.ServerURL)

		defer resp.Body.Close()
	}

	return downServers
}

func openFiles(serverListFile string, downTimeFile string) (*os.File, *os.File) {
	serverList, err := os.OpenFile(serverListFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	downTimeList, err := os.OpenFile(downTimeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return serverList, downTimeList
}

func generateDowntime(downTimeList *os.File, downServers []Server) {
	csvWriter := csv.NewWriter(downTimeList)
	for _, servidor := range downServers {
		line := []string{servidor.ServerName, servidor.ServerURL, servidor.dataFalha, fmt.Sprintf("f", servidor.tempoExecucao), fmt.Sprintf("d", servidor.status)}
		csvWriter.Write(line)
	}
	csvWriter.Flush()
}

func main() {
	serverList, downTimeList := openFiles(LISTA_SERVIDORES_CSV_LOCALIZACAO, LISTA_SERVIDORES_DOWNTIME_CSV_LOCALIZACAO)

	defer serverList.Close()
	defer downTimeList.Close()

	servidores := criarListaServidores(serverList)
	downServers := checkServers(servidores)
	generateDowntime(downTimeList, downServers)
}
