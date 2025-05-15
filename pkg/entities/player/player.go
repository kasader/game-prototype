package player

import "github.com/kasader/game-prototype/pkg/grid"

type Player struct {
	x_pos int16
	y_pos int16

	x_min int16
	x_max int16
	y_min int16
	y_max int16
}

func GetTestPlayer() *Player {
	return &Player{
		x_min: 0,
		x_max: 8,
		y_min: 0,
		y_max: 8,
	}
}

func (p *Player) UpdatePosition(x, y int16, g *grid.Grid) {
	newX := p.x_pos + x
	if newX >= p.x_min && newX < p.x_max {
		p.x_pos = newX
	}

	newY := p.y_pos + y
	if newY >= p.y_min && newY < p.y_max {
		p.y_pos = newY
	}
}

func (p *Player) GetPosition() (x, y int16) {
	return p.x_pos, p.y_pos
}
