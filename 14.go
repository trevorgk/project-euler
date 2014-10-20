package main

import (
        "github.com/trevorgk/project-euler/eulerlib"
        "fmt"
)

//	Longest Collatz sequence
//	Which starting number, under one million, 
//	produces the longest chain?
func euler_14() {
	maxLength, maxStartingNum := 0,0
	for i := 2; i < 1000000; i++ {
	
		length := eulerlib.Collatz(i, false)
		
		if length > maxLength {
			maxLength = length
			maxStartingNum = i
		}
	}
	
	fmt.Println("Max starting number is: ", maxStartingNum)
}