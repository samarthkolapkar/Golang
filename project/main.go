package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func query_city(p1 []Person, city string) {
	for _, data := range p1 {
		if data.City == city {
			fmt.Printf("Name: %s\n", data.Name)
			// fmt.Printf("Age: %s", data.Age)
		}
	}
}
func validcity(p1 []Person, city string) bool {
	for _, p := range p1 {
		if city == p.City {
			return true
		}
	}
	return false
}
func query_age(p1 []Person, age int) {
	for _, p := range p1 {
		if p.Age == age {
			fmt.Printf("Name: %s\t", p.Name)
			fmt.Printf("City: %s\n", p.City)
		}
	}

}
func unmarshal_JSON(data []byte) (p1 []Person) {
	//var p1 []Person
	err := json.Unmarshal(data, &p1)
	if err != nil {
		log.Fatal(err)
	}

	//query_to_json(p1)

	return
}

func readJSON() (data []byte) {
	data, err1 := ioutil.ReadFile("data.json")
	if err1 != nil {
		fmt.Println("Error While reading the JSON file!!")
	}
	return

}

func main() {
	var age int
	var choice int
	var city string
	data := readJSON()
	stored := unmarshal_JSON(data)
	fmt.Println("1:Find the name of person by name of city\n2:Find the name of person by age")
	fmt.Printf("Enter the choice of operation you want to perform:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("Enter the age to find the names and city where they live in JSON:")
		fmt.Scan(&age)
		query_age(stored, age)
	case 2:
		fmt.Printf("Enter the name of city to find the person name:")
		fmt.Scan(&city)
		if validcity(stored, city) {
			query_city(stored, city)
		} else {
			fmt.Println("City is not valid please enter the valid city!!")
		}
	}

	// fmt.Println(string(file))
	// query_to_json(string(file))

}
