package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account2 struct {
	Name     string `json:"name"`
	Password string `json:"pass_word"`
	Money    float64 `json:"money"`
}

type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Extra   []interface{}
	Level   map[string]interface{}
	Account Account2
}

func main() {

	extra := []interface{}{123, "hello world"}

	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	level := make(map[string]interface{})

	level["bash"] = 100
	level["python"] = "垃圾"
	level["golang"] = nil

	account := Account2{
		Name:     "rzp",
		Password: "boom1234",
		Money:    4.50,
	}

	user := User{
		Name:    "ruizhipeng",
		Age:     25,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		Extra:   extra,
		Level:   level,
		Account: account,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(rs))

}
