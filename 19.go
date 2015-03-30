package main

import (
	"time"
	"fmt"
	//"github.com/trevorgk/project-euler/eulerlib"
)

/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

 */
func euler_19 () {

	result := 0;

	date := time.Date(1901,time.January,1, 0, 0, 0, 0, time.UTC)
	for date.Weekday() != time.Sunday {
		date = date.Add(time.Hour * 24)
	}

	for ; date.Year() < 2001; date = date.Add(time.Hour * 24 * 7) {
		if date.Weekday() == time.Sunday && date.Day() == 1 {
			result++
			//fmt.Println(date);
		}
	}
	fmt.Println("There were", result, "Sundays on the first day of the month in the twentieth century")
	return
}
