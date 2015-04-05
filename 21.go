package main

import(
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
)

func euler_21() {
	//a, b := 220, 284
	//fmt.Println("sum of divisors of", a, "is", eulerlib.SumOfDivisors(a))
	//fmt.Println("sum of divisors of", b, "is", eulerlib.SumOfDivisors(b))
	//fmt.Println("amicable:", eulerlib.AreAmicablePairs(a,b))

	//amicablePairs := make([]int,0)
	size := 10000
	sumOfDivisors := make([]int,size)
	for i := 0; i < size; i++ {
		sumOfDivisors[i] = eulerlib.SumOfDivisors(i);
	}

	total := 0;
	for i := 0; i < size; i++ {
		sum := sumOfDivisors[i];
		if sum >= size {
			//	have to calculate as it's outside the array bounds
			if eulerlib.SumOfDivisors(sum) == sum {
				total += sum
				fmt.Println(sum, "is amicable")
			}
		} else {
			//	can use the precalculated result from the array
			if sumOfDivisors[i] == sum {
				total += sum
				fmt.Println(sum, "is amicable")
			}
		}
	}

	fmt.Println("total is: ", total);
}
