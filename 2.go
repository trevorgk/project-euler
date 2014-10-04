package main

import (	
	"fmt"
	"github.com/trevorgk/project-euler/eulerlib"
)
//	find the sum of the even-valued terms of fibonacci sequence where values
//	do next exceed n
func euler_2(n int) {

	f := eulerlib.Fibonacci()
	result := 0
	for i := f(); i < n; i = f() {
		if i%2 == 0 {
			result += i
		}
	}
	fmt.Println("result is",result)
}
