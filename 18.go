package main

import (
	//"github.com/trevorgk/project-euler/eulerlib"
	"fmt"
)

//Find the maximum total from top to bottom of the triangle below
func euler_18 () {
	grid := [][]int {
		[]int{			  				  75							},	//	0
		[]int{							95, 64							},
		[]int{						  17, 47, 82						},
		[]int{					    18, 35, 87, 10						},
		[]int{					  20, 4, 82, 47, 65						},
		[]int{				    19, 1, 23, 75, 3, 34					},
		[]int{				  88, 2, 77, 73, 7, 63, 67					},
		[]int{			    99, 65,  4, 28,  6, 16, 70, 92				},
		[]int{			  41, 41, 26, 56, 83, 40, 80, 70, 33			},
		[]int{		    41, 48, 72, 33, 47, 32, 37, 16, 94, 29			},
		[]int{		  53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14		},
		[]int{	    70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57		},
		[]int{	  91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48	},
		[]int{	63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31	},
		[]int{ 4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23	}};//14

	fmt.Println("total is", maxRecurse(grid));
}

func subTree(grid [][]int, pos int) (subTree [][]int){
	//	subtree length = length(grid) - 1
	subTree = make([][]int, len(grid) - 1)

	for i := 0; i < len(grid) - 1; i++ {
		subTree[i] = make([]int, i + 1);

		for j := 0; j <= i; j++ {
			//fmt.Printf("subTree[%d][%d] = %d\n", i, j, grid[i+1][j] )
			subTree[i][j] = grid[i+1][j + pos]
		}
	}

	return
}
	
func getAvailablePositions(pos int) (left, right int){
	return pos, pos+1;
}

//	deprecated in favour of recursive version
func maxTraverse(grid [][]int) (total int){
	total = 0;
	for i, pos := 0,0; i < len(grid) - 1; i++{
		value := grid[i][pos];
		total += value;
		//fmt.Printf("grid[%d][%d] = %d [%d]\n", i, pos, value, total )
		left, right := getAvailablePositions(pos);
		//fmt.Println(grid[i+1][left], grid[i+1][right]);
		pos = left;
		if grid[i+1][left] < grid[i+1][right] {
			pos = right;
		}
	}
	return
}

func maxRecurse(grid [][]int) (total int) {
	if len(grid) == 1 {
		total = grid[0][0];
	} else {
		//	create secondary slices
		var treeLeft = subTree(grid, 0)
		var treeRight = subTree(grid, 1)
		//fmt.Println("treeLeft", treeLeft);
		//fmt.Println("treeRight", treeRight);

		if maxRecurse(treeLeft) >= maxRecurse(treeRight) {
			total = grid[0][0] + maxRecurse(treeLeft);
		} else {
			total = grid[0][0] + maxRecurse(treeRight);
		}

	}
	//fmt.Println("len(grid)", len(grid), "total:", total);

	return;
}

