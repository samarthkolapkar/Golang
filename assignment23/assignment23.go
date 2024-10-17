/*Same as (22) to accept a list of holidays, and then prompt a user for two inputs: input date
as a first argument and a number of business days as a second argument. Number of
business days will be a positive or negative whole number. The output shall be the date
relative to an input date +/- the number of business days. Holidays must be excluded while
calculating the output date.*/

package main

import (
	"fmt"
	"time"
)

type calender struct {
	holidays []string
}

//This function will calculate the bussiness day
func calculate_businessdays(input_date1 time.Time, business_day int, holidays []string) time.Time {

	for i := 0; i < abs(business_day); {
		// Move the input date by one day forward or backward depending on the business_day sign
		input_date1 = input_date1.AddDate(0, 0, business_day/abs(business_day))

		// Check if the resulting date is a business day and not a holiday
		if isBusinessDay(input_date1.Weekday()) && !isHoliday(input_date1.Format("2006-01-02"), holidays) {
			i++
		}
	}
	return input_date1
}

//checks the date that user has enter in the holidays string array
func isHoliday(date string, holidays []string) bool {
	for _, holiday := range holidays {
		if date == holiday {
			return true
		}
	}
	return false
}

//if the saturday and sunday is there in between date then it will return that
func isBusinessDay(weekday time.Weekday) bool {
	return weekday != time.Saturday && weekday != time.Sunday
}

//if user want to check the previous business days in that case user will enter the negative number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	var holidays []string
	var holiday_date string
	var input_date string
	var business_day int
	fmt.Println("Enter the holiday dates in fomrat (YYYY-MM-DD)!!")
	fmt.Println("After entering all the holidays enter stop to stop entering!!")

	for {
		fmt.Scan(&holiday_date)
		if holiday_date == "stop" {
			break
		}
		holidays = append(holidays, holiday_date)
	}
	fmt.Printf("Enter the input date:")
	fmt.Scan(&input_date)
	input_date1, err := time.Parse("2006-01-02", input_date)
	if err != nil {
		fmt.Println("Error while parsing the date")
	}

	fmt.Printf("Enter the number of business day:")
	fmt.Scan(&business_day)

	resultDate := calculate_businessdays(input_date1, business_day, holidays)
	fmt.Println("Output Date:", resultDate)

}
