package main

import (
	"fmt"
	"github.com/trevorgk/project-euler/eulerlib"
)

//	Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
func euler_8(s string) {
	runes := []rune(s)
	max := [] int {0,0,0,0,0}
	productMax := 0
	for i := 0; i < len(s) - 5; i++ {
		snippet := runes[i:i+5]
		//fmt.Println("snippet", snippet)
		productMax = max[0] * max[1] * max[2] * max[3] * max[4]
		productSnippet := eulerlib.RuneToInt(snippet[0]) * eulerlib.RuneToInt(snippet[1]) * eulerlib.RuneToInt(snippet[2]) * eulerlib.RuneToInt(snippet[3]) * eulerlib.RuneToInt(snippet[4])
		
		if productSnippet > productMax {
			for j := 0; j < 5; j++ {
				max[j] = eulerlib.RuneToInt(snippet[j])
			}
		}
	}
    fmt.Println("max snippet is", max, "which gives a product of", productMax)
}

