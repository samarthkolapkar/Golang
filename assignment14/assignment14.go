package main

import (
	"fmt"
	"time"
)

//This function is used to format the date
func formated_time(currentTime time.Time) string {
	//"2 Jan 2006" is used as a reference
	formattedDate := ""
	formattedDate += "Formatted date:\n"
	formattedDate += currentTime.Format("2 Jan 2006") + "\n"
	formattedDate += currentTime.Format("Jan 2 2006") + "\n"
	formattedDate += currentTime.Format("2006-01-02") + "\n"
	formattedDate += currentTime.Format("2006-02-02T15:04:05Z") + "\n"
	formattedDate += currentTime.Format("Monday, 2006-01-02") + "\n"
	return formattedDate
}
func main() {
	// Get current UTC time
	currentTime := time.Now().UTC()

	formatedDate := formated_time(currentTime)

	fmt.Println(formatedDate)

}
