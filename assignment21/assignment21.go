/*Write a program to accept a list of holidays (date in any of the supported formats). Store this
list in an internal array. After the user confirms entering of the holiday list, accept a date from
the user, and confirm whether it's a working day or not. (All Saturdays and Sundays are
implicitly considered as holidays).*/
package main

import (
	"fmt"
	"time"
)

type calender struct {
	holidays []string
}

//function will check the date enter by user is holiday or not for saturday and sunday
func weekdays(check_date1 time.Time) bool {
	weekday := check_date1.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

//this function will check that date in the holiday list if it is there then it will return true
func check_holiday(holiday []string, check_date string) bool {
	for i := range holiday {
		// weekday := check_date.Weekday()
		if check_date == holiday[i] {
			return true
		}
	}
	return false
}

func main() {
	var holidays []string
	var holiday_date string
	var check_date string
	fmt.Println("Enter the holiday dates in fomrat (YYYY-MM-DD)!!")
	fmt.Println("After entering all the holidays enter stop to stop entering!!")

	for {
		fmt.Scan(&holiday_date)
		if holiday_date == "stop" {
			break
		}
		holidays = append(holidays, holiday_date)
	}
	fmt.Println(holidays)
	fmt.Printf("Enter the date you want to check for holidays:")
	fmt.Scan(&check_date)
	check_date1, _ := time.Parse("2006-01-02", check_date)
	if check_holiday(holidays, check_date) || weekdays(check_date1) {
		fmt.Println("It's a holiday")
	} else {
		fmt.Println("Not a holiday")
	}

}
