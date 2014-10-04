package main

import (
	"fmt"
	"github.com/trevorgk/project-euler/eulerlib"
)

//	Find the difference between the sum of the squares 
//	of the first n natural numbers and the 
//	square of the sum
func euler_6 (n int) {
	fmt.Println("difference is", eulerlib.SquareOfSums(n) - eulerlib.SumOfSquares(n))
}
