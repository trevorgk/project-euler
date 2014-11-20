package main

import "fmt"
import "math"


func getDivisors(n int) []int {
	bound := int(math.Sqrt(float64(n)))
	divisors := make([]int,0)
	divisors = append(divisors, 1)
	for i:= 2; i <= bound + 1; i++ {
		if (n % i == 0) {
			divisors = append(divisors, i, n/i)
		}
	}
	divisors = append(divisors, n)
	return divisors
}

func main () {
	
	
	num := 1
	for i := 1; i<=10001; i++ {
		for len(getDivisors(num)) != 2 {
			num++
		}
		fmt.Println(num, "is the", i, "prime number")
		num++
	}
}