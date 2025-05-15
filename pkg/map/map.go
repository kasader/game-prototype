package grid

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kasader/game-prototype/pkg/player"
	"github.com/kasader/game-prototype/pkg/tile"
)

type Grid struct {
	x_len  uint16
	y_len  uint16
	grid   [][]tile.Tile
	Player *player.Player
}

func NewGrid(x, y uint16) *Grid {
	return &Grid{
		x_len: x,
		y_len: y,
		grid:  make([][]tile.Tile, x),
	}
}

func GetTestGrid() *Grid {
	return &Grid{
		x_len: 8,
		y_len: 8,
		grid: [][]tile.Tile{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 2, 2, 0, 0, 0},
			{0, 0, 0, 2, 2, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
		Player: player.GetTestPlayer(),
	}
}

func (g *Grid) String() string {
	dgrid := make([][]tile.Tile, len(g.grid))
	for i := range g.grid { // deep copy the grid...
		dgrid[i] = make([]tile.Tile, len(g.grid[i]))
		copy(dgrid[i], g.grid[i])
	}

	// place player on the grid
	x_pos, y_pos := g.Player.GetPosition()
	dgrid[y_pos][x_pos] = tile.Player

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
