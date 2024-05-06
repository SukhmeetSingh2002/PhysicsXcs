package main

import (
	"fmt"
	"math"
)

type signal func(float64) float64

func f_t(t float64) float64 {
	return math.Sin(2 * math.Pi * t)
}

func decomposedSignal(result []Result, t float64) float64 {
	var signal float64 = 0
	for _, res := range result {
		signal += res.Amplitude * math.Cos(2*math.Pi*res.Frequency*t+res.Phase)
	}
	return signal
}

func main() {
	timeFFTImpl()

	// ref: https://youtu.be/mkGsMWi_j4Q?si=fOWOdx6yhox0-cse
	const N_sample int = 1024
	const T float64 = 9


	var results []Result = analyzeSignal(f_t, T, N_sample)
	for _, val := range results {
		fmt.Printf("Frequency: %v, Amplitude: %v, Phase: %v\n", val.Frequency, val.Amplitude, val.Phase)
	}
	verification, errors := verifyDecomposition(results, N_sample, T, 1e-4)
	analyzeError(errors)
	if verification {
		fmt.Println("Decomposition is correct.")
	} else {
		fmt.Println("Decomposition is erroneous.")
	}
}
