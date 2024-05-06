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

func analyzeError(errors []float64) {
	var N int = len(errors)
	var maxError float64 = 0
	var minError float64 = math.MaxFloat64
	var sumError float64 = 0
	for _, val := range errors {
		if val > maxError {
			maxError = val
		}
		if val < minError {
			minError = val
		}
		sumError += val
	}
	var avgError float64 = sumError / float64(N)
	var stdDeviation float64 = 0
	for _, val := range errors {
		stdDeviation += math.Pow(val-avgError, 2)
	}
	stdDeviation = math.Sqrt(stdDeviation / float64(N))
	fmt.Printf("Max Error: %v, Min Error: %v, Average Error: %v, Standard Deviation: %v\n", maxError, minError, avgError, stdDeviation)
}
