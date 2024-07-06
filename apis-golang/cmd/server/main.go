package main

import "procom/configs"

func main(){
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(configs.DBDriver)
	println(configs.DBHost)
}