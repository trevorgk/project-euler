package main

import (
	"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
)

type Triplet struct {
	x int
	y int
	z int
}

// find a Pythagorean triplet for which a + b + c = n	
func euler_9(n int) {

	triplets := make([]Triplet,0)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			for k := 1; k <= n - i - j; k++ {
				if (i + j + k) == n {
					triplet := Triplet {i, j, k}
					triplets = append(triplets, triplet)
					//fmt.Println("appending triplet", triplet)
				}
			}	
		}	
	}
	
	for i := 0; i < len(triplets); i++ {
		triplet := triplets[i]
		if (eulerlib.Exponent(triplet.x, 2) + eulerlib.Exponent(triplet.y, 2)) == eulerlib.Exponent(triplet.z,2) {
			fmt.Println(triplet, "is the triplet")		
			return	
		}
	}
	
	fmt.Println("triplet does not exist")

}
