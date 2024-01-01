package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"numero"` //possui tags, possoescrever dessa forma no json
	Balance int `json:"saldo" validate:"gt=0"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	result, err := json.Marshal(account) // guardando na variável
	if err != nil {
		panic(err)
	}

	println(string(result))

	err = json.NewEncoder((os.Stdout)).Encode(account) //dump //inprimindo no terminal
	if err != nil {
		panic(err)
	}

	jsonData := []byte(`{"Number": 2, "Balance": 200}`)
	var accountJson Account
	err = json.Unmarshal(jsonData, &accountJson)
	if err != nil {
		panic(err)
	}
	println(accountJson.Balance)
}
