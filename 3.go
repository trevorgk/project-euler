package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
)

//	Find largest prime factor of number n
func euler_3(n int) {
	fmt.Printf("result is: %d", eulerlib.LargestPrimeDivisor(n))
}
