package gamemap

type Grid struct {
	grid [][]*Tile
}

func NewGrid(x, y int) *Grid {
	return &Grid{
		grid: make([][]*Tile, x),
	}
}

func GetTestGrid() *Grid {
	return &Grid{
		grid: [][]*Tile{
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileWall, TileEmpty, TileWall, TileWall, TileWall, TileEmpty, TileEmpty},
			{TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileWall, TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileWall, TileWall, TileWall, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileWall, TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileWall, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
			{TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty},
		},
	}
}

func (g *Grid) InBounds(newX, newY int) bool {
	// First check for X in-bounds, followed by check Y in-bounds.
	return (newX >= 0 && newX < g.Width()) && (newY >= 0 && newY < g.Height())
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

func (g *Grid) Width() int  { return len(g.grid[0]) }
func (g *Grid) Height() int { return len(g.grid) }

func (g *Grid) GetTile(x, y int) *Tile {
	return g.grid[y][x]
}
