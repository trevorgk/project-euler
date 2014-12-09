package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
	"time"
)

//	Unlike vanilla solution below, this func uses go's very handy built in - LOC reduced greatly
func euler_19() {

	sundaysOnTheFirst, date, week := 0, time.Date(1900, time.January, 7, 0, 0, 0, 0, time.UTC), time.Hour * 24 * 7

	for ; date.Year() <= 2000; date = date.Add(week) {
		if date.Day() == 1 && date.Year() >= 1901 {
			sundaysOnTheFirst++
		}
		//fmt.Println("date is:", date)
	}
		
	fmt.Println("Sundays on the first:", sundaysOnTheFirst)
}

//	solution implemented using almost no built-ins (except constants) 
func euler_19_vanilla() {

	sundaysOnTheFirst, 
		day, month, year := 0, 7, time.January, 1900

	for {
		day += 7

		if (year > 2000) {
			break;
		}
		
		daysInMonth := eulerlib.DaysInMonth(month, year, false)
		
		if (day / daysInMonth) > 0 {
			day -= daysInMonth
			if (day == 1 && year >= 1901) {
				sundaysOnTheFirst++
			}
			
			if (month == time.December) {
				month = time.January
				year++
			} else {
				month++
			}
			//fmt.Printf("date is %v/%v/%v\n", day, month, year)
		}
	}

	fmt.Println("Sundays on the first:", sundaysOnTheFirst)
}
