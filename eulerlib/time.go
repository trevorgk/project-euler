package eulerlib

import (
	"fmt"
	"time"
	"errors"
)

func IsLeapYear(year int) bool {
	return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}

func DaysInMonth(month time.Month, year int, print bool) int {
	days := 31
	switch (month){
		case time.September:
		case time.April:
		case time.June:
		case time.November:
			days = 30
		case time.February:
			days = 28
			if IsLeapYear(year) {
				days = 29
			}
	}
	if print {
		fmt.Println("days in", month, year, "is", days)
	}
	return days

}

//	Could have deleted this as it doesn't do it's role but interesting as a template for error handling
func DateIsA(weekday time.Weekday, day int, month time.Month, year int) (bool, error)  {
	daysInMonth := DaysInMonth(month, year, false)

	if (day > daysInMonth){
		return false, errors.New("days in month exceeded")
	}
	return true, nil
}
