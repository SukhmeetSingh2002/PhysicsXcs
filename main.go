package main

import (
	"fmt"
	// "sync"
	// "time"
)

func main() {
	timeFFTImpl()

	var analysisResults []Result = analyzeSignal()
	for _, val := range analysisResults {
		fmt.Printf("Frequency: %d, Amplitude: %f\n", val.Frequency, val.Amplitude)
	}

}
