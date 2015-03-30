package eulerlib

import (
	"fmt"
	"math"
	"math/big"
)

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

func GetDivisors(n int, includeSelf bool) []int {
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

func SumOfDivisors(n int) int {
	sum, divisors := 0, GetDivisors(n);
	for i := 0; i < len(divisors); i++ {
		sum += divisors[i];
	}
	return sum;
}

func AreAmicablePairs(a int, b int) bool {
	return SumOfDivisors(a) == SumOfDivisors(b);
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

//	see http://simple.wikipedia.org/wiki/Exponentiation_by_squaring
func BinaryExponent(x uint64, n int) uint64 {

	fmt.Println("BinaryExponent", x, "**", n)
	if n == 1 {
		return x
	}
	
	if n % 2 == 0 {
		return BinaryExponent(x * x, n/2)
	}

	return x * BinaryExponent(x * x, (n-1)/2)
}

func FactorialInt(n int64) *big.Int {
	return Factorial(big.NewInt(n))
}

func Factorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n == one {
		//fmt.Printf("1")
		return one
	}
	//fmt.Printf("%d * ", n)

	j := new(big.Int).Sub(n, one)
	i := new(big.Int).Mul(n, Factorial(j))
	return i
}
