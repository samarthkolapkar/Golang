/*Write a program to accept two dates and print the count of week-end days (Consider
Saturdays and Sundays as week-ends).*/

package main

import (
	"fmt"
	"time"
)

//This function will calculate the number of weekends in between dates that user has enter
func count_number_of_weekends(date1_parse, date2_parse time.Time) int {
	weekend_count := 0

	for date1_parse.Before(date2_parse) || date1_parse.After(date1_parse) {
		//store the weekdays form the startin date
		weekday := date1_parse.Weekday()
		// fmt.Println(weekday)
		if weekday == time.Saturday || weekday == time.Sunday {
			weekend_count += 1
		}
		date1_parse = date1_parse.AddDate(0, 0, 1)

	}
	return weekend_count
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
	weekend_count := count_number_of_weekends(date1_parse, date2_parse)
	fmt.Printf("Weekend count is %d\n", weekend_count)
}
