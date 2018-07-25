package main

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful-swagger12"
)

var jsonString4 string = `{
	"things": [
		{
			"name":"rzp",
			"age":"26"
		},
		{
			"city":"Tianjin",
			"country":"China"
		},
		{
			"name":"pt",
			"age":"27"
		},
		{
			"city":"Heilongjiang",
			"country":"China"
		}
	]
}`

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Place struct {
	City string `json:"city"`
	Country string `json:"country"`
}

func Decode4(jsonStr []byte) (persons []Person,places []Place)  {
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonStr,&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range data["things"]{
		item := data["things"][i]
		if item["name"]!=nil{
			persons = addPerson(persons)
		}
	}
	{

	}
}
