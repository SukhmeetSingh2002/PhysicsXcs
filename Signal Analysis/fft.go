package main

import (
	"math"
	"math/cmplx"
	"sync"
)

func FFTParallel(p []complex128, n int, wg *sync.WaitGroup) []complex128 {
	defer wg.Done()
	if n&(n-1) != 0 {
		panic("n must be a power of 2")
	}
	if n == 1 {
		return p
	}
	var u []complex128 = make([]complex128, n/2)
	var v []complex128 = make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		u[i] = p[2*i]
		v[i] = p[2*i+1]
	}
	var u_fft, v_fft []complex128
	var wg2 sync.WaitGroup
	wg2.Add(2)

	go func() {
		u_fft = FFTParallel(u, n/2, &wg2)

	}()
	go func() {
		v_fft = FFTParallel(v, n/2, &wg2)
	}()
	wg2.Wait()
	var p_fft []complex128 = make([]complex128, n)
	var factor float64 = 2 * math.Pi / float64(n)
	var twiddleStep complex128 = cmplx.Exp(complex(0, -factor))
	var twiddle complex128 = 1
	for i := 0; i < n/2; i++ {
		p_fft[i] = u_fft[i] + twiddle*v_fft[i]
		p_fft[i+n/2] = u_fft[i] - twiddle*v_fft[i]
		twiddle = twiddle * twiddleStep
	}

	return p_fft
}

func InverseFFT(p_fft []complex128, n int) []complex128 {
	if n&(n-1) != 0 {
		panic("n must be a power of 2")
	}
	if n == 1 {
		return p_fft
	}
	for ind, val := range p_fft {
		p_fft[ind] = cmplx.Conj(val)
	}
	var p []complex128 = FFT(p_fft, n)
	for ind, val := range p {
		p[ind] = cmplx.Conj(val) / complex(float64(n), 0)
	}
	return p

}

func FFT(p []complex128, n int) []complex128 {
	if n&(n-1) != 0 {
		panic("n must be a power of 2")
	}
	if n == 1 {
		return p
	}
	var u []complex128 = make([]complex128, n/2)
	var v []complex128 = make([]complex128, n/2)

	for i := 0; i < n/2; i++ {
		u[i] = p[2*i]
		v[i] = p[2*i+1]
	}
	var u_fft []complex128 = FFT(u, n/2)
	var v_fft []complex128 = FFT(v, n/2)
	var p_fft []complex128 = make([]complex128, n)

	var factor float64 = 2 * math.Pi / float64(n)
	var twiddle complex128 = complex(1, 0)
	var twiddleStep complex128 = cmplx.Exp(complex(0, -factor))
	for i := 0; i < n/2; i++ {
		p_fft[i] = u_fft[i] + twiddle*v_fft[i]
		p_fft[i+n/2] = u_fft[i] - twiddle*v_fft[i]
		twiddle = twiddle * twiddleStep
	}
	return p_fft
}
