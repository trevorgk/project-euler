package main

import (
	"fmt"
	"github.com/trevorgk/project-euler/eulerlib"
)


func euler_7(max int) {
	
	num := 1
	for i := 1; i<=max; i++ {
		for len(eulerlib.GetDivisors(num)) != 2 {
			num++
		}
		fmt.Println(num, "is the", i, "prime number")
		num++
	}
}