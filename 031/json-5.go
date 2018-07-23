package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account5 struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Money    float64 `json:"money,string"`
}

var JsonString string = `{
	"email":"rzp@gmail.com",
	"password":"123",
	"money":"100.5"
}`

func main() {
	account := Account5{}

	err := json.Unmarshal([]byte(JsonString), &account)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", account)
}
