package eulerlib

import (
	"fmt"
	"math"
	"unicode/utf8"
	"strings"
	"time"
	"errors"
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

func IsPrime (n int) bool {
	return len(GetDivisors(n)) == 2
}

func LargestPrimeDivisor(n int) int {
	divisors := GetDivisors(n)
	largestPrime := 0
	for _, value := range divisors {
		if IsPrime(value) {
			if largestPrime < value {
				largestPrime = value
			}
		}
	}
	return largestPrime
}

func GetDivisors(n int) []int {
	bound := int(math.Sqrt(float64(n + 1)))
	divisors := make([]int,0)
	divisors = append(divisors, 1)
	for i:= 2; i <= bound; i++ {
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


func Collatz(startingNum int, doPrint bool) int {
	num := startingNum
	length := 0
	
	if doPrint { 
		fmt.Print(num)
	}
	
	for num != 1 {
		length++
		if num % 2 == 0 {
			num = num/2
		} else {
			num = 3 * num + 1
		}
		if doPrint { 
			fmt.Print("->", num)
		}
		
	}
	if doPrint { 	
		fmt.Print("\n")
	}
	
	return length
}

func WordCount(verbiage string, includeWhitespace bool) int {
	if (includeWhitespace) {
		return utf8.RuneCountInString(verbiage)
	}
	sum, split := 0, strings.Split(verbiage, " ")

	for i := 0; i < len(split); i++ {
		sum += utf8.RuneCountInString(split[i])
	}

	return sum
}

func IsLeapYear(year int) bool {
	return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
}

func DaysInMonth(month time.Month, year int) int {
	switch (month){
	case time.September:
	case time.April:
	case time.June:
	case time.November:
		return 30
	case time.February:
		if IsLeapYear(year) {
			return 29
		}

		return 28
	}

	return 31
}

func DateIsA(weekday time.Weekday, day int, month time.Month, year int) (bool, error)  {
	daysInMonth := DaysInMonth(month, year)

	if (day > daysInMonth){
		return false, errors.New("days in month exceeded")
	}
	return true, nil
}
