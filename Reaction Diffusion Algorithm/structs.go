package main

type Pixel struct {
	conc_a float64
	conc_b float64
}
type FavoriteValues struct {
	feed float64
	k    float64
	D_a  float64
	D_b  float64
}
type Orientation struct { // models shape of container
	center_x int
	center_y int
	dx       float64 // inreases diffusion rate by dx per pixel going away from center
	dy       float64
}
