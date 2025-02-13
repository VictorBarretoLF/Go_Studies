package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	fmt.Println("Server is running on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	json.NewEncoder(w).Encode(cep)
}

func BuscaCep(cep string) (*ViaCep, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var viaCepResponse ViaCep
	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		return nil, err
	}

	return &viaCepResponse, nil
}