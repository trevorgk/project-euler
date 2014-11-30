package main

import (
	"fmt"
	//"math"
	"math/big"
	//"strconv"
)

//	What is the sum of the digits of the number 2^n
func euler_16(n int) {
	pow := big.NewInt(1 << uint(n))
	fmt.Printf("2**%v yields %v\n",n, pow)//, ([]byte(pow)))
	//fmt.Println(pow)//, ([]byte(pow)))
}
