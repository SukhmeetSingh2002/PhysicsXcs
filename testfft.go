package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"sync"
	"time"
)

type Result struct {
	Frequency int
	Amplitude float64
}

func f_t(t float64) float64 {
	return math.Sin(2*math.Pi*t) + 5*math.Sin(4*math.Pi*t) + 2*math.Sin(6*math.Pi*t) + 3*math.Sin(8*math.Pi*t)
}
func analyzeSignal() []Result {
	const T float64 = 1
	const N int = 8
	const t float64 = T / float64(N)
	const f float64 = 1 / t
	const sensitivity float64 = 1e-5

	var n_nyquist int = int(math.Floor(f / 2))
	var samples []complex128 = make([]complex128, N)
	for i := 0; i < N; i++ {
		samples[i] = complex(f_t(float64(i)*t), 0)
	}
	var p_fft_normal []complex128 = FFT(samples, N)
	fmt.Printf("---------FFT Analysis for frequencies less than %v-----------\n", n_nyquist)

	p_fft_normal = p_fft_normal[:(n_nyquist + 1)]
	// create slice for stuct Result
	var results []Result
	for ind, val := range p_fft_normal {
		if math.Abs(real(val)) < sensitivity {
			p_fft_normal[ind] = complex(0, imag(val))
		}
		if math.Abs(imag(p_fft_normal[ind])) < sensitivity {
			p_fft_normal[ind] = complex(real(p_fft_normal[ind]), 0)
		}
		p_fft_normal[ind] = p_fft_normal[ind] * complex(2, 0)
		if cmplx.Abs(p_fft_normal[ind]) > sensitivity {
			var Amplitude float64 = cmplx.Abs(p_fft_normal[ind]) / float64(N)
			var res Result = Result{ind, Amplitude}
			results = append(results, res)
		}
	}
	return results
}

func timeFFTImpl() {
	fmt.Println("---------FFT-----------")
	var p_1 []complex128 = []complex128{}
	for i := 0; i < 1024; i++ {
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

	if checkSame(&p_fft_normal, &p_fft_parallel, 1e-5) {
		fmt.Println("Both FFTs are same")
	} else {
		panic("Both FFTs are not same")
	}

	fmt.Println("---------Inverse FFT-----------")
	var invertFFTSignal []complex128 = InverseFFT(p_fft_normal, n)
	if checkSame(&p_1, &invertFFTSignal, 1e-5) {
		fmt.Println("Inverted signal is same as original")
	} else {
		panic("Inverted signal is not same as original")
	}
}
