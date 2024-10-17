package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type student struct {
	Roll_no int    `json:"Roll_number"`
	Name    string `json:"Name,omitempty"`
	Marks   int    `json:"-"`
}

func main() {
	s := student{Roll_no: 1, Name: "Samarth", Marks: 56}
	//marshal- converting the structure data to the JSON format
	sbytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json Data :", string(sbytes))
	var student1 student
	filename := "example_1.json"
	//RawMessage-used when we want to delay in json encoding or precomputing
	s2raw, err1 := ioutil.ReadFile(filename)
	fmt.Println(err1)
	//unmarshal- converting the JSON data to the structures
	err2 := json.Unmarshal(s2raw, &student1)
	fmt.Println(err2)
	fmt.Println(student1)

}
