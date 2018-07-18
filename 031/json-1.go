package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account struct {
	Email    string
	Password string
	Money    float64
}

func main() {
	account := Account{
		Email:    "ruizhipeng001@gmail.com",
		Password: "booom1234",
		Money:    10.25,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
}
