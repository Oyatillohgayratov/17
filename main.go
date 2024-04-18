package main

import (
	"encoding/json"
	"homework/config"
	"homework/methods"
	"log"
	"os"
)

func main() {
	file, err := os.Open("employees.json")
	if err != nil {
		log.Fatal(err)
	}

	var persons []config.User

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&persons)
	if err != nil {
		log.Fatal(err)
	}
	
	persons = methods.Age(persons)
	persons = methods.Add(persons)


	if err != nil {
		log.Fatal(err)
    }
	f, err := os.OpenFile("employees_new.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	err = encoder.Encode(persons)
	if err != nil {

		log.Fatal(err)
	}




}
