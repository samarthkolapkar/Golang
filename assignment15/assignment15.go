package main

import (
	"fmt"
	"time"
)

// func isTimezone_valid(timezone string) bool {
// 	return false
// }
//This function will return the IST timezone in specific format
func IST_time(currentTime time.Time, timezone string) (string, error) {
	location, _ := time.LoadLocation(timezone)
	// if !isTimezone_valid(timezone) {
	// 	return "", fmt.Errorf("The %s time zone is not supported!!", timezone)
	// }
	istime := currentTime.In(location)
	// fmt.Println(location1)
	// fmt.Println(location1.Format("2 Jan 2006 15:04:05"))
	formated_ist := ""
	formated_ist += istime.Format("2 Jan 2006") + "\n"
	formated_ist += istime.Format("Jan 2 2006") + "\n"
	formated_ist += istime.Format("2006-01-02") + "\n"
	formated_ist += istime.Format("2006-02-02 T 15:04:05Z")

	return formated_ist, nil
}

func main() {
	currentTime := time.Now().UTC()
	timezone := "Asia/Kolkata"
	istime, err := IST_time(currentTime, timezone)
	if err != nil {
		fmt.Println("Error while loading the IST!!")
	}
	fmt.Println(istime)
}
