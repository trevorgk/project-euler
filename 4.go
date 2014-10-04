package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
	"strconv"
)

type Tuple struct {
	x int
	y int
}

//	Find the largest palindrome made from the product of two n-digit numbers
func euler_4(n int) {
	factors := make([]Tuple,0)
	max := Tuple{0,0}

	start := eulerlib.Exponent(10,n)
	fmt.Println("starting at", start)
	for i := start;i >= 0; i-- {
		for j := i; j >= 0; j-- {
			product := i * j
			if (eulerlib.IsPalindrome(strconv.Itoa(product))) {
				factors = append(factors, Tuple{i,j})
				if (i + j) > (max.x + max.y) {
					max = Tuple {i, j}
				}
			}
		}
	}
	fmt.Println("highest pair are:", max.x, "and", max.y,"which result in palindrome:", max.x * max.y)
}
