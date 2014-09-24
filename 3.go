package main

import "fmt"
import "math"

func sumOfDivisors(n int) int {
	sum := 1
	bound := int(math.Sqrt(float64(n)))
	for i:= 2; i<= bound + 1; i++ {
		if (n % i == 0) {
			sum = sum + i + n/i
		}
	}
	return sum
}

func main() {

	fmt.Printf("result is: %d",sumOfDivisors(13195))
}