package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type User2 struct {
	Username interface{} `json:"username"`
	Password string      `json:"password"`
}

var jsonString2 string = `{
	"username":123456999989,
	"password":"boom123"
}`

func Decode2(r io.Reader) (u *User2, err error) {
	u = new(User2)
	if err = json.NewDecoder(r).Decode(u); err != nil {
		return
	}
	switch t := u.Username.(type) {
	case string:
		u.Username = t
	case float64:
		u.Username = int64(t)
	}
	return
}

func main() {
	user, err := Decode2(strings.NewReader(jsonString2))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", user)
}
