package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EbitenInput struct{}

func (e *EbitenInput) GetDirection() (dx, dy int) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		dx = -1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		dx = +1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		dy = +1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		dy = -1
	}
	return
}
