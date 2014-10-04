package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
)

//	Find the sum of all the primes below n
func euler_10(n int) {
	sum := 0
	for i := 2; i < n; i++ {
		if eulerlib.IsPrime(i) {
			//fmt.Println(i, "is prime")
			sum += i
		}
	}
	
	fmt.Println("Sum of all primes below", n, "is", sum)
}
