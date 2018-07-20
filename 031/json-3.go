package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account3 struct {
	Gmail    string  `json:"gmail"`
	Password string  `json:password`
	Money    float64 `json:"money"`
}

var jsonString string = `{
	"gmail":"rzp@123.com",
	"password":"123boom",
	"money":23.45
}`

func main() {

	account := Account3{}

	err := json.Unmarshal([]byte(jsonString), &account)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v/n", account)
}
