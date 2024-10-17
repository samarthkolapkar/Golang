package main

import "fmt"

func main() {
	var x int
	fmt.Print("Enter the number between 1-7 to know the day!!:")
	fmt.Scan(&x)
	switch x {
	case 1:
		fmt.Println("Todays is Monday!!")
	case 2:
		fmt.Println("today is Tuesday!!")
	case 3:
		fmt.Println("Today is Wednesday!!")
	case 4:
		fmt.Println("Today is Thursday!!")
	case 5:
		fmt.Println("Today is Friday!!")

	case 6:
		fmt.Println("Today is Saturday!!")

	case 7:
		fmt.Println("Today is Sunday!!")

	}

}
