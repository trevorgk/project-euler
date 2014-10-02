package main

import "fmt"

func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += (i * i)
	}
	fmt.Println("sum of squares is", result)
	return result
}

func SquareOfSums (n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	result = (result * result)
	fmt.Println("square of sums is", result)
	
	return result
}

func main () {
	fmt.Println("difference is", SquareOfSums(100) - SumOfSquares(100))
}