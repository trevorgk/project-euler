package main

import "fmt"
import "math"

func largestPrimeDivisor(n int) int {
	divisors := getDivisors(n)
	largestPrime := 0
	for _, value := range divisors {
		if len(getDivisors(value)) == 2 {
			if largestPrime < value {
				largestPrime = value
			}
		}
	}
	return largestPrime
}

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

func main() {

	fmt.Printf("result is: %d", largestPrimeDivisor(600851475143))
}