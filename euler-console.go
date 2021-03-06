package main

import (
	"flag"
	"fmt"
	"time"
	"strconv"
)

func main() {
	
	probPtr := flag.Int("prob",0, "Problem to solve")
	flag.Parse()

	if *probPtr == 0 {
		fmt.Println("euler-console is a tool for running my own project euler solutions")
		fmt.Println("Usage:")
		fmt.Println("\t project-euler -prob=[num] [arg1] ...")
		return
	}
	
	fmt.Println("project euler, problem", *probPtr)
	fmt.Println("arguments:", flag.Args())
	fmt.Println("----------------------")
	start := time.Now()

	switch *probPtr {
	case 1:
		euler_1(IntArg(flag.Arg(0)), IntArg(flag.Arg(1)), IntArg(flag.Arg(2)))
	case 2:
		euler_2(IntArg(flag.Arg(0)))
	case 3:
		euler_3(IntArg(flag.Arg(0)))
	case 4:
		euler_4(IntArg(flag.Arg(0)))
	case 5:
		euler_5(IntArg(flag.Arg(0)))
	case 6:
		euler_6(IntArg(flag.Arg(0)))
	case 7:
		euler_7(IntArg(flag.Arg(0)))
	case 8:
		s := "73167176531330624919225119674426574742355349194934" +
				"96983520312774506326239578318016984801869478851843" +
				"85861560789112949495459501737958331952853208805511" +
				"12540698747158523863050715693290963295227443043557" +
				"66896648950445244523161731856403098711121722383113" +
				"62229893423380308135336276614282806444486645238749" +
				"30358907296290491560440772390713810515859307960866" +
				"70172427121883998797908792274921901699720888093776" +
				"65727333001053367881220235421809751254540594752243" +
				"52584907711670556013604839586446706324415722155397" +
				"53697817977846174064955149290862569321978468622482" +
				"83972241375657056057490261407972968652414535100474" +
				"82166370484403199890008895243450658541227588666881" +
				"16427171479924442928230863465674813919123162824586" +
				"17866458359124566529476545682848912883142607690042" +
				"24219022671055626321111109370544217506941658960408" +
				"07198403850962455444362981230987879927244284909188" +
				"84580156166097919133875499200524063689912560717606" +
				"05886116467109405077541002256983155200055935729725" +
				"71636269561882670428252483600823257530420752963450"
		euler_8(s)
	case 9:
		euler_9(IntArg(flag.Arg(0)))
	case 10:
		euler_10(IntArg(flag.Arg(0)))
	case 11:
		euler_11()
	case 12:
		euler_12(500)
	case 14:
		euler_14()
	case 16:
		euler_16(IntArg(flag.Arg(0)))
	case 17:
		euler_17()
	case 18:
		euler_18()
	case 19:
		euler_19()
	case 20:
		euler_20(IntArg(flag.Arg(0)))
	case 21:
		euler_21()

	default:
		fmt.Println("solution not implemented")
		return
	}

 	elapsed := time.Since(start)
	fmt.Println("----------------------")
 	fmt.Println("Solution took", elapsed)
}

func IntArg(s string) int {
	k, _ := strconv.Atoi(s)

	return k
}

func Float64Arg(s string) float64 {
	k, _ := strconv.ParseFloat(s, 64)

	return k
}
