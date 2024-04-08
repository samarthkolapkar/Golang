package main

import (
	"fmt"
	"strings"
)

const no_of_tickets uint = 50

type userdata struct {
	fname      string
	lname      string
	age        int
	emailID    string
	no_tickets uint
	moviename  string
}

func main() {
	greetuser()
	fname, lname, age, emailID, moviename, no_tickets := Userinput()
	isvalidName, isvalidEmail, isvalidAge, isvalidtickets := validate(fname, lname, age, emailID, no_tickets)
	if isvalidName && isvalidEmail && isvalidAge && isvalidtickets {
		booktickets(fname, lname, age, emailID, moviename, uint(no_tickets))
	} else {
		if !isvalidName {
			fmt.Print("Entered name is to short please re-enter the name!!\n")
		}
		if !isvalidEmail {
			fmt.Print("Entered Email address is not valid please enter it again!!\n")
		}
		if !isvalidAge {
			fmt.Print("Entered age is not valid enter a valid age!!!\n")
		}
	}

}
func greetuser() {
	fmt.Print("Welcome to Booking Aap!!\n")
	fmt.Print("We are here to book tickets for movies!!\n")
}
func Userinput() (string, string, int, string, string, int) {
	var fname string
	var lname string
	var age int
	var emailID string
	var moviename string
	var no_tickets int
	fmt.Print("Enter your first name:")
	fmt.Scan(&fname)
	fmt.Print("Enter your last name:")
	fmt.Scan(&lname)
	fmt.Print(("Enter your age:"))
	fmt.Scan(&age)
	fmt.Print("Enter your email id:")
	fmt.Scan(&emailID)
	fmt.Print("Enter the name of movie:")
	fmt.Scan(&moviename)
	fmt.Print("Enter the number of tickets you want to book!!:")
	fmt.Scan(&no_tickets)
	return fname, lname, age, emailID, moviename, no_tickets

}
func validate(fname string, lname string, age int, emailID string, no_tickets int) (bool, bool, bool, bool) {

	isvalidName := len(fname) >= 2 && len(lname) >= 2
	isvalidAge := age >= 18
	isvalidEmail := strings.Contains(emailID, "@")
	isvalidtickets := no_of_tickets >= uint(no_tickets)
	return isvalidName, isvalidEmail, isvalidAge, isvalidtickets

}

func booktickets(fname string, lname string, age int, moviename string, emailID string, no_tickets uint) {
	remaining_tickets := no_of_tickets - no_tickets
	fmt.Print("\n", remaining_tickets)
	var Userdata = userdata{
		fname:      fname,
		lname:      lname,
		age:        age,
		emailID:    emailID,
		no_tickets: no_tickets,
		moviename:  moviename,
	}
	fmt.Print(Userdata)

}
