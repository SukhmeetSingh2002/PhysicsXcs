package main

import (
	"fmt"
	"math"
)

func checkSame(p1 *[]complex128, p2 *[]complex128, sensitivity float64) bool {
	if len(*p1) != len(*p2) {
		return false
	}
	for i := range *p1 {
		var real_normal float64 = real((*p1)[i])
		var imag_normal float64 = imag((*p1)[i])
		var real_parallel float64 = real((*p2)[i])
		var imag_parallel float64 = imag((*p2)[i])

		if math.Abs(real_normal-real_parallel) > sensitivity || math.Abs(imag_normal-imag_parallel) > sensitivity {
			fmt.Printf("Index: %d, Normal: %v, Parallel: %v\n", i, (*p1)[i], (*p2)[i])
			return false
		}
	}
	return true
}
