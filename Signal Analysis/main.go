package main

import (
	"fmt"
	"math"
)

type signal func(float64) float64

func f_t(t float64) float64 {
	return 1 * math.Sin(2*math.Pi*6*t) + 2*math.Sin(2*math.Pi*10*t) + 3*math.Sin(2*math.Pi*15*t)
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
	const N_sample int = 4096 * 2
	const T float64 = 1 // 2*pi if it's 4*t,2*t etc. else take 1

	var samples []complex128 = make([]complex128, N_sample)
	var t float64 = T / float64(N_sample)
	for i := 0; i < N_sample; i++ {
		samples[i] = complex(f_t(float64(i)*t), 0)
	}
	var results []Result = analyzeSignal(samples, T, N_sample)
	for _, val := range results {
		fmt.Printf("Frequency: %v, Amplitude: %v, Phase: %v\n", val.Frequency, val.Amplitude, val.Phase)
	}
	verification, errors := verifyDecomposition(results, f_t, N_sample, T, 1e-4)
	analyzeError(errors)
	if verification {
		fmt.Println("Decomposition is correct.")
	} else {
		fmt.Println("Decomposition is erroneous.")
	}

	// infile_path := flag.String("infile", "Asine.wav", "Input wav file path")
	// flag.Parse()
	// file, _ := os.Open(*infile_path)
	// reader := wav.NewReader(file)
	// defer file.Close()

	// // read .wav file
	// fmt.Println("Reading wav file...")
	// var wavSamples []complex128
	// for {
	// 	samples, err := reader.ReadSamples()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	for _, sample := range samples {
	// 		wavSamples = append(wavSamples, complex(float64(reader.IntValue(sample, 0)), 0))
	// 	}
	// }

	// T, err := reader.Duration()
	// if err != nil {
	// 	panic(err)
	// }
	// var fs float64 = float64(len(wavSamples)) / T.Seconds()
	// var N_sample int = len(wavSamples)
	// var d float64 = 1 / fs

	// // var original_sample_len int = N_sample

	// // 0-padding
	// for N_sample&(N_sample-1) != 0 {
	// 	wavSamples = append(wavSamples, complex(0, 0))
	// 	N_sample++
	// }
	// fft_data := FFT(wavSamples, N_sample)
	// var freqs []float64 = make([]float64, N_sample)
	// freqs = fftfreq(N_sample, d)

	// fmt.Println("length of wavSamples: ", len(wavSamples))
	// fmt.Printf("length of wavSamples:%v fs:%v \n", len(wavSamples), fs)

	// var results []Result =  analyzeSignal(wavSamples, 2.048, 16384)
	// for _, val := range results {
	// 	fmt.Printf("Frequency: %v, Amplitude: %v, Phase: %v\n", val.Frequency, val.Amplitude, val.Phase)
	// }
	// verification, errors := verifyDecomposition(results, f_t, 16384, T.Seconds(), 1e-4)
	// analyzeError(errors)
	// if verification {
	// 	fmt.Println("Decomposition is correct.")
	// } else {
	// 	fmt.Println("Decomposition is erroneous.")
	// }

}
