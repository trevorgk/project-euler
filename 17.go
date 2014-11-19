package main

import (
	"fmt"
	"strconv"
	"github.com/trevorgk/project-euler/eulerlib"
)
func euler_17(n int) {
	s := strconv.FormatInt(int64(n), 10)
	fmt.Println(s[0])
	fmt.Println(eulerlib.WordCount(s, false))
}
