/*Same as (21) to accept a list of holidays, and then prompt a user for two dates (in the
supported format) and print the number of working days between two dates. Consider both
dates during the calculation.*/
package main

import (
	"fmt"
	"time"
)

//calculate the working days
func count_number_of_weekdays(date1_parse, date2_parse time.Time) int {
	weekdays_count := 0

	for date1_parse.Before(date2_parse) || date1_parse.After(date1_parse) {
		//store the weekdays form the startin date
		weekday := date1_parse.Weekday()
		// fmt.Println(weekday)
		if weekday != time.Saturday && weekday != time.Sunday {
			weekdays_count += 1
		}
		//increment the date by 1
		date1_parse = date1_parse.AddDate(0, 0, 1)

	}
	return weekdays_count
}

func main() {
	var first_date string
	var second_date string

	fmt.Println("Enter the date as per the format(DD-MM-YYYY)")
	fmt.Printf("Enter the first date:")
	fmt.Scan(&first_date)
	date1_parse, err := time.Parse("2006-01-02", first_date)
	if err != nil {
		fmt.Println("Error while parsing the date")
	}
	fmt.Printf("Enter the second date:")
	fmt.Scan(&second_date)
	date2_parse, err := time.Parse("2006-01-02", second_date)
	if err != nil {
		fmt.Println("Error while parsing the date")
	}
	weekdays_count := count_number_of_weekdays(date1_parse, date2_parse)
	fmt.Printf("Weekend count is %d\n", weekdays_count)
}
