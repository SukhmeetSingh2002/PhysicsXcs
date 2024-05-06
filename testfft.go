package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"sync"
	"time"
)

type Result struct {
	Frequency float64
	Amplitude float64
	Phase     float64
}

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

func verifyDecomposition(results []Result, N int, T float64, tolerance float64) (bool, []float64) {
	var reconstructedSignal []float64 = make([]float64, N)
	for i := 0; i < N; i++ {
		t := float64(i) * T / float64(N)
		reconstructedSignal[i] = decomposedSignal(results, t)
	}

	var maxError float64 = 0
	var errors []float64 = make([]float64, N)
	for i := 0; i < N; i++ {
		originalValue := f_t(float64(i) * T / float64(N))
		error := math.Abs(originalValue - reconstructedSignal[i])
		// fmt.Println("Original: ", originalValue, "Reconstructed: ", reconstructedSignal[i], "Error: ", error)
		errors[i] = error
		if error > maxError {
			maxError = error
		}
	}

	return maxError <= tolerance, errors
}

func analyzeSignal() []Result {
	const T float64 = 9
	const N int = 8192
	const t float64 = T / float64(N)
	const f float64 = 1 / t
	const df float64 = f / float64(N)
	const sensitivity float64 = 1e-5

	var n_nyquist int = int(math.Floor(f / 2))
	var samples []complex128 = make([]complex128, N)
	for i := 0; i < N; i++ {
		samples[i] = complex(f_t(float64(i)*t), 0)
	}
	var p_fft_normal []complex128 = FFT(samples, N)

	fmt.Printf("---------FFT Analysis for frequencies less than %v-----------\n", n_nyquist)
	p_fft_normal = p_fft_normal[:(n_nyquist)]
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
			var phase float64 = cmplx.Phase(p_fft_normal[ind])
			var res Result = Result{float64(ind) * df, Amplitude, phase}
			results = append(results, res)
		}
	}
	if len(results) == 0 {
		fmt.Println("No significant frequency found")
		return results
	}
	for _, val := range results {
		fmt.Printf("Frequency: %v, Amplitude: %v, Phase: %v\n", val.Frequency, val.Amplitude, val.Phase)
	}
	verification, errors := verifyDecomposition(results, N, T, 1e-4)
	analyzeError(errors)
	if verification {
		fmt.Println("Decomposition is correct.")
	} else {
		fmt.Println("Decomposition is erroneous.")
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
