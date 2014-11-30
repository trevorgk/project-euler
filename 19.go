package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
	"time"
)

func euler_19() {

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
		//runcount := 0
		//if (runcount < 25 && year >= 1901) {
		//	runcount++
		//	fmt.Printf("date is %v/%v/%v\n", day, month, year)
		//}
	}

	fmt.Println("Sundays on the first:", sundaysOnTheFirst)
}
