package main

import "fmt"

//	Find the sum of all the multiples of x or y below n
func euler_1(x int, y int, n int) {

	result := 0
	for i := 1; i < (n - 1); i++ {
		if i % x == 0 || i % y == 0 {
			//fmt.Println("iteration",i)
			result += i
		}
	}
	fmt.Println("sum of all the multiples of", x, "or", y, "below", n, "is", result)
}
