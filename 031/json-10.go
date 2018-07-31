package main

import (
	"encoding/json"
	"fmt"
)

//混合结构
var jsonString5 string = `{
	"things":[
		{
			"name":"RZP",
			"age":25
		},
		{
			"city":"Tianjin",
			"country":"China"
		},
		{
			"name":"PT",
			"age":26
		},
		{
			"city":"Heilongjiang",	
			"country":"China"
		}
	]
}`

type Mixed struct {
	Name    string `json:"name"`
	Age2    int    `json:"age"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func Decode5(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]Mixed
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("+v\n", data["things"])
	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{Name: item.Name, Age: item.Age2})
		} else {
			places = append(places, Place{City: item.City, Country: item.Country})
		}
		return
	}
	return
}

func main() {
	personsB, placesB := Decode5([]byte(jsonString5))
	fmt.Printf("%+v\n", personsB)
	fmt.Printf("%%+v\n", placesB)
	var data2 map[string][]Mixed
	var jsonStr2 []byte
	err := json.Unmarshal(jsonStr2, &data2)
	if err != nil {

	}
}

//golang会初始化没有匹配的json和抛弃没有匹配的json，匹配到的内容会进行赋值 没有匹配到的就是零值
