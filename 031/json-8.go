package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type User3 struct {
	Username json.RawMessage `json:"username"`
	password string          `json:"password"`

	Email string
	Phone int64
}

var jsonString3 string = `{
	"username":"123@abc.com",
	"password":"123"
}`

func Decode3(r io.Reader) (u *User3, err error) {
	u = new(User3)
	if err = json.NewDecoder(r).Decode(u); err != nil {
		return
	}
	var email string
	if err = json.Unmarshal(u.Username, &email); err == nil {
		u.Email = email
		return
	}
	var phone int64
	if err = json.Unmarshal(u.Username, &phone); err == nil {
		u.Phone = phone
	}
	return
}

func main() {
	user, err := Decode3(strings.NewReader(jsonString3))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", user)
}
