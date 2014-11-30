package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
	"time"
)

func euler_19() {

	//if _ , err := eulerlib.DateIsA(time.Monday, 31, 2, 2000); err != nil{
	//	fmt.Println(err)
	//}

	sundaysOnTheFirst, 
		day, month, year := 0, 7, time.January, 1900

	
	for {
		day += 7
		if (year > 2000) {
			break;
		}
		
		daysInMonth := eulerlib.DaysInMonth(month, year)
		
		if (day % daysInMonth) > 0 {
			day -= daysInMonth
			if (day == 1) {
				sundaysOnTheFirst++
				fmt.Println(day, month, year, "is a Sunday")
			}
			
			month += 1
			if (month == time.December) {
				month = time.January
				year++
			} else {
				month++
			}
		}
	}

	fmt.Println("Sundays on the first:", sundaysOnTheFirst)
}
