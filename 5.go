package main

import "fmt"

func main() {
	val := 1
OuterLoop:
	for {
		val++
		
		for i := 1; i <= 20; i++ {
			if (val % i != 0) {
				continue OuterLoop;
			}
		}
		
		fmt.Println(val,"is the one!")
		return;
	}
}