package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account4 struct {
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
	Money    float64 `json:"money,string"`
}

func main() {
	account := Account4{
		Email:    "rzp@abc.com",
		Password: "boom1234",
		Money:    23.50,
	}
	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)
	fmt.Println(string(rs))
}
