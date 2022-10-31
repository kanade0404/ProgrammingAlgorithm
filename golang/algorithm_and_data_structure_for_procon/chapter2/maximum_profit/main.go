package main

import (
	"fmt"
	"math"
)

func Solve(ns []int) int {
	var max, min = -2000000000, ns[0]
	nsLen := len(ns)
	for i := 1; i < nsLen; i++ {
		fmt.Println(math.Max(float64(max), float64(ns[i]-min)))
		max = int(math.Max(float64(max), float64(ns[i]-min)))
		fmt.Println(math.Min(float64(min), float64(ns[i])))
		min = int(math.Min(float64(min), float64(ns[i])))
	}
	return max
}
