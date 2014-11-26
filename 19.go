package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
	"time"
)

func euler_19() {

	if _ , err := eulerlib.DateIsA(time.Monday, 31, 2, 2000); err != nil{
		fmt.Println(err)
	}

	sundaysOnTheFirst,// day, month,
		year := 0,// 1, time.January,
		1900

	for i := 0;;i++ {
		if (year > 2000) {
			break;
		}


	}

	fmt.Println("Sundays on the first:", sundaysOnTheFirst)
}
