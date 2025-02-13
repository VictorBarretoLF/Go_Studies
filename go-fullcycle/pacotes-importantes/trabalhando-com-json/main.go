package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo int
}

type ContaV2 struct {
	IdTransacao int `json:"i"`
	Numero int `json:"-"` // o '-' é para ignorar
	Saldo int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res)) // {"Numero":1,"Saldo":100}

	enconder := json.NewEncoder(os.Stdout)
	err = enconder.Encode(conta) // {"Numero":1,"Saldo":100}
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo) // 200

	jsonPuro = []byte(`{"n": 3, "s": 300}`)
	var contaY Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaY.Saldo) // 0

	jsonPuro = []byte(`{"i": 1, "n": 3, "s": 300}`)
	var contaV2 ContaV2
	err = json.Unmarshal(jsonPuro, &contaV2)
	if err != nil {
		panic(err)
	}
	println(contaV2.Saldo, contaV2.Numero, contaV2.IdTransacao) // 300 0 1
}