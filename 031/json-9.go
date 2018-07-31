package main

import (
	"encoding/json"
	"fmt"
)

var jsonString4 string = `{
	"things": [
		{
			"name":"rzp",
			"age":26
		},
		{
			"city":"Tianjin",
			"country":"China"
		},
		{
			"name":"pt",
			"age":27
		},
		{
			"city":"Heilongjiang",
			"country":"China"
		}
	]
}`

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func addPerson(persons []Person, item map[string]interface{}) []Person {
	name := item["name"].(string)
	age := item["age"].(float64)
	person := Person{name, int(age)}
	persons = append(persons, person)
	return persons
}

func addPlace(places []Place, item map[string]interface{}) []Place {
	city := item["city"].(string)
	country := item["country"].(string)
	place := Place{city, country}
	places = append(places, place)
	return places
}

func Decode4(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range data["things"] {
		item := data["things"][i]
		if item["name"] != nil {
			persons = addPerson(persons, item)
		} else {
			places = addPlace(places, item)
		}
	}
	return
}

func main() {
	personsA, placesA := Decode4([]byte(jsonString4))
	fmt.Printf("%+v\n", personsA)
	fmt.Printf("%+v\n", placesA)

}
