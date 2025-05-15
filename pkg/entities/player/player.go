package player

import "github.com/kasader/game-prototype/pkg/gamemap"

type Player struct {
	X int
	Y int

	// FIXME: This should be read-in from the grid bounds, etc.
	x_min int
	x_max int
	y_min int
	y_max int
}

func GetTestPlayer() *Player {
	return &Player{
		x_min: 0,
		x_max: 8,
		y_min: 0,
		y_max: 8,
	}
}

func (p *Player) GetPosition() (x, y int) {
	return p.X, p.Y
}

func (p *Player) TryMove(dx, dy int, g *gamemap.Grid) {
	newX := p.X + dx
	newY := p.Y + dy

	if g.InBounds(newX, newY) && g.IsWalkable(newX, newY) {
		p.X = newX
		p.Y = newY
	}
}
