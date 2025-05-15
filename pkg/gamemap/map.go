package gamemap

import (
	"fmt"
	"slices"
	"strings"
)

type Grid struct {
	grid         [][]Tile
	x_len, y_len int
}

func NewGrid(x, y int) *Grid {
	return &Grid{
		x_len: x,
		y_len: y,
		grid:  make([][]Tile, x),
	}
}

func GetTestGrid() *Grid {
	return &Grid{
		x_len: 8,
		y_len: 8,
		grid: [][]Tile{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 2, 2, 0, 0, 0},
			{0, 0, 0, 2, 2, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
}

func (g *Grid) String() string {
	dgrid := make([][]Tile, len(g.grid))
	for i := range g.grid { // deep copy the grid...
		dgrid[i] = make([]Tile, len(g.grid[i]))
		copy(dgrid[i], g.grid[i])
	}

	// place player on the grid
	// x_pos, y_pos := g.Player.GetPosition()
	// dgrid[y_pos][x_pos] = Player

	// draw the grid:
	var gridStrs []string
	for _, y := range dgrid {
		var gridLine []string
		for _, x := range y {
			gridLine = append(gridLine, fmt.Sprintf("%d", x))
		}
		gridStrs = append(gridStrs, "["+strings.Join(gridLine, " ")+"]")
	}
	slices.Reverse(gridStrs)

	return strings.Join(gridStrs, "\n")
}

func (g *Grid) InBounds(newX, newY int) bool {
	// First check for X in-bounds, followed by check Y in-bounds.
	return (newX >= 0 && newX < g.x_len) && (newY >= 0 && newY < g.y_len)
}

func (g *Grid) IsWalkable(newX, newY int) bool {
	tile := g.grid[newY][newX]
	switch tile {
	case TileWall:
		return false
	default:
		return true
	}
}

func (g *Grid) Width() int  { return g.x_len }
func (g *Grid) Height() int { return g.y_len }

func (g *Grid) GetTile(x, y int) Tile {
	return g.grid[y][x]
}
