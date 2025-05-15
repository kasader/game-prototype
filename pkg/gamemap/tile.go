package gamemap

type SomeAbstractionAboveATile struct {
	height int
}

type Tile struct {
	isWalkable bool
	symbol     string
}

func (t *Tile) String() string {
	return t.symbol
}

func (t *Tile) IsWalkable() bool {
	return t.isWalkable
}

var TileEmpty = &Tile{
	isWalkable: true,
	symbol:     "0",
}

var TileWall = &Tile{
	isWalkable: true,
	symbol:     "1",
}
