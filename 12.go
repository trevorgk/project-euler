package main

import (
        "github.com/trevorgk/project-euler/eulerlib"
        "fmt"
)

//	Highly divisible triangular number
//	What is the value of the first triangle number
//	to have over n divisors
func euler_12(numDivisors int) {
	
	i, sum := 0, 0
	
	for {
		sum += i
		divisors := eulerlib.GetDivisors(sum)
		
		if len(divisors) > numDivisors {
			fmt.Println(sum,":",divisors)
			break;
		}
		i++
	}
}
