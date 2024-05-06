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
		var diff complex128 = (*p1)[i] - (*p2)[i]
		if math.Abs(real(diff)) > sensitivity || math.Abs(imag(diff)) > sensitivity {
			fmt.Printf("Index: %d, Normal: %v, Parallel: %v\n", i, (*p1)[i], (*p2)[i])
			return false
		}
	}
	return true
}
