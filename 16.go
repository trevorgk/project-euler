package main

import (
	"fmt"
	"math"
	//"math/big"
	//"strconv"
)

//	What is the sum of the digits of the number 2^n
func euler_16(n float64) {
	pow := math.Pow(2,n)
	fmt.Println("2**i yields", pow)//, ([]byte(pow)))
}
