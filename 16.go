package main

import (
	"fmt"
    "github.com/trevorgk/project-euler/eulerlib"
	//"math"
	//"math/big"
	//"strconv"
)

//	What is the sum of the digits of the number 2^n
func euler_16(n int) {
	pow := eulerlib.BinaryExponent(2, n)
	fmt.Printf("2**%v yields %v\n", uint64(n), pow)//, ([]byte(pow)))
	//fmt.Println(pow)//, ([]byte(pow)))
}
