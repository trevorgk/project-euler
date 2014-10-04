package eulerlib

import (
	"fmt"
	"math"
)


func RuneToInt(r rune) int {
	return int(r - '0')
}

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

//	fibonacci - nested function solution
func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y - x
	}
}

func LargestPrimeDivisor(n int) int {
	divisors := GetDivisors(n)
	largestPrime := 0
	for _, value := range divisors {
		if len(GetDivisors(value)) == 2 {
			if largestPrime < value {
				largestPrime = value
			}
		}
	}
	return largestPrime
}

func GetDivisors(n int) []int {
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


func IsPalindrome(s string) bool {
	reverse := Reverse(s)
	return reverse == s
}

func Reverse(s string) string  {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Exponent (num int, n int) int {
	i := 1;
	result := num
	for i < n {
		result *= num
		i++
	}
	return result
}
