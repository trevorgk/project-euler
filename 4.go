package main

import (
	"fmt"
	"strconv"
)

type Tuple struct {
	x int
	y int
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
func main() {
	factors := make([]Tuple,0)
	max := Tuple{0,0}
	
	for i := 999;i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if (IsPalindrome(strconv.Itoa(product))) {
				//fmt.Println(product, "is a palindrome comprised of", i, "and", j)
				factors = append(factors, Tuple{i,j})
				if (i + j) > (max.x + max.y) {
					max = Tuple {i, j}
				}
			}
		}
	}
	fmt.Println("highest pair are:", max.x, "and", max.y,"which result in palindrome:", max.x * max.y)
}