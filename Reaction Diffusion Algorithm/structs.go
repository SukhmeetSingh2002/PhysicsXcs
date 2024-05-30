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

type Cell struct {
	FeedRate float64
	KillRate float64
}

type Grid struct {
	Cells     [][]Cell
	StyleMap  [][]Cell
}

func NewGrid(width, height int, feed, kill float64) *Grid {

	width, height = height, width

	var center_x int = width / 2
	var center_y int = height / 2

	cells := make([][]Cell, height)
	styleMap := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
		styleMap[i] = make([]Cell, width)
		for j := range cells[i] {
			// changing the feed and kill rates
			feedRate := feed + float64(j-center_x)*(0.08-0.0)/float64(width-1)
			killRate := kill + float64(i-center_y)*(0.07-0.03)/float64(height-1)

			clip(&feedRate, 0, 1)
			clip(&killRate, 0, 1)

			styleMap[i][j] = Cell{FeedRate: feedRate, KillRate: killRate}
		}
	}
	return &Grid{Cells: cells, StyleMap: styleMap}
}
