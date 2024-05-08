package main

func clip[T int | float64 | float32](x *T, min T, max T) {
	if *x < min {
		*x = min
	}
	if *x > max {
		*x = max
	}
}
func laplacian(a bool, x int, y int) float64 {
	var sum float64 = 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			var temp float64
			if a {
				temp = pixels_t1[x+i][y+j].conc_a
			} else {
				temp = pixels_t1[x+i][y+j].conc_b
			}
			sum += laplacian_kernel[i+1][j+1] * temp
		}
	}
	return sum
}

func swap(pixels_t1 *[width][height]Pixel, pixels_t2 *[width][height]Pixel) {
	var pixels_temp [width][height]Pixel = *pixels_t1
	*pixels_t1 = *pixels_t2
	*pixels_t2 = pixels_temp
}
