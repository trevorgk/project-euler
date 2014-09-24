package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y - x
	}
}

func main() {

	f := fibonacci()
	result := 0
	for i := f(); i < 4000000; i = f() {
		if i%2 == 0 {
			result += i
		}
	}
	fmt.Printf("result is: %d",result)
}