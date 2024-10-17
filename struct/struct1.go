package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
	DOB   string
}

func main() {
	var p person
	var p2 person
	p = person{fname: "Samarth", lname: "Kolapkar", age: 22, DOB: "20/02/2002"}
	p1 := person{fname: "Dinesh", lname: "Lokare", age: 23, DOB: "12/12/2001"}
	fmt.Printf("Enter the first name:")
	fmt.Scan(&p2.fname)
	fmt.Printf("Enter the last name:")
	fmt.Scan(&p2.lname)
	fmt.Print("Enter the age :")
	fmt.Scan(&p2.age)
	fmt.Print("Enter the date of birth(DD/MM/YYYY):")
	fmt.Scan(&p2.DOB)
	fmt.Println("First Name\t\tLast name\t\tAge\t\tDOB")
	fmt.Printf("%s\t\t\t%s\t\t%d\t\t%s\n", p.fname, p.lname, p.age, p.DOB)
	fmt.Printf("%s\t\t\t%s\t\t\t%d\t\t%s\n", p1.fname, p1.lname, p1.age, p1.DOB)
	fmt.Printf("%s\t\t\t%s\t\t\t%d\t\t%s\n", p2.fname, p2.lname, p2.age, p2.DOB)
}
