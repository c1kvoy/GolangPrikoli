package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID       uint8  `json:"id"`
	Username string `json:"username"` // структурные теги чтобы сериализовать и десериализовать объеты
	Email    string `json:"mail"`     // оч удобно :)
}

func main() {
	json_data := []byte(`{
		"id":       8,
		"username": "Kolya",
		"mail":     "Koulyaa@mail"
	}`)

	var person Person
	err := json.Unmarshal(json_data, &person)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Println("Unmarshled data")
	fmt.Println(person)
	encoded, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err.Error(), "Error in encoding")
	}
	fmt.Println("Marshled data")
	fmt.Println(string(encoded))

}
