package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
    CEP         string `json:"cep"`
    Logradouro  string `json:"logradouro"`
    Complemento string `json:"complemento"`
    Unidade     string `json:"unidade"`
    Bairro      string `json:"bairro"`
    Localidade  string `json:"localidade"`
    UF          string `json:"uf"`
    Estado      string `json:"estado"`
    Regiao      string `json:"regiao"`
    IBGE        string `json:"ibge"`
    GIA         string `json:"gia"`
    DDD         string `json:"ddd"`
    SIAFI       string `json:"siafi"`
}

func main()  {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao realizar a requisição: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o parsing da resposta: %v\n", err)
		}

		fmt.Println(data)

		output := fmt.Sprintf(
			"CEP: %s\nLogradouro: %s\nComplemento: %s\nUnidade: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegiao: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
			data.CEP, data.Logradouro, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.UF, data.Estado, data.Regiao, data.IBGE, data.GIA, data.DDD, data.SIAFI,
		)

		file, err := os.Create("cidades.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
	}
}