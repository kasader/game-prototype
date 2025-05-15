package grid

type Grid struct {
	x    uint16
	y    uint16
	grid [][]string
}

func NewGrid(x, y uint16) *Grid {
	return &Grid{
		x:    x,
		y:    y,
		grid: make([][]string, x),
	}
}
