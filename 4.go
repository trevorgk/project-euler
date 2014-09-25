package main

import (
	"fmt"
	"strconv"
)

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
	for i := 999;i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if (IsPalindrome(strconv.Itoa(product))) {
				fmt.Println(product, "is a palindrome comprised of", i, "and", j)
				return
			}
		}
	}
}