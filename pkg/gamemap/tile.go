package gamemap

type Tile byte

const (
	TileEmpty Tile = iota
	TilePlayer
	TileWall
)
