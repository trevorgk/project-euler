package main

import (
	"fmt"
)

//	find the smallest positive number that is evenly
//	divisible by all of the numbers from 1 to n
func euler_5(n int) {
	val := 1
OuterLoop:
	for {
		val++
		
		for i := 1; i <= n; i++ {
			if (val % i != 0) {
				continue OuterLoop;
			}
		}
		
		fmt.Println(val,"is the one!")
		return;
	}
}
