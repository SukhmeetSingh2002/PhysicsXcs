package main

import (
	"fmt"
	"sync"
	"time"
)

  


func main() {
	fmt.Println("---------FFT-----------")
	var p_1 []complex128 = []complex128{}
	for i := 1; i < 5; i++ {
		p_1 = append(p_1, complex(float64(i), 0))
	}
	var n int = len(p_1)

	var p_fft_normal []complex128 = make([]complex128, n)
	var p_fft_parallel []complex128 = make([]complex128, n)
	t0 := time.Now()
	p_fft_normal = FFT(p_1, n)
	fmt.Printf("Time taken by normal FFT: %v\n", time.Since(t0))

	var wg sync.WaitGroup
	t0 = time.Now()
	wg.Add(1)
	p_fft_parallel = FFTParallel(p_1, n, &wg)
	wg.Wait()
	fmt.Printf("Time taken by parallel FFT: %v\n", time.Since(t0))

	
	fmt.Printf("Result from normal FFT: %v\n",p_fft_normal)
	fmt.Printf("Result from parallel FFT: %v\n",p_fft_parallel)

	if checkSame(&p_fft_normal, &p_fft_parallel, 1e-5) {
		fmt.Println("Both FFTs are same")
	} else {
		panic("Both FFTs are not same")
	}

	fmt.Println("---------Inverse FFT-----------")
	var invertFFTSignal []complex128 = InverseFFT(p_fft_normal, n)
	fmt.Printf("Inverse FFT: %v\n", invertFFTSignal)
	if checkSame(&p_1, &invertFFTSignal, 1e-5) {
		fmt.Println("Inverted signal is same as original")
	} else {
		panic("Inverted signal is not same as original")
	}


}
