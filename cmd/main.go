// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"slices"
	"strings"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/kasader/game-prototype/pkg/entities/player"
	"github.com/kasader/game-prototype/pkg/gamemap"
	grid "github.com/kasader/game-prototype/pkg/gamemap"
	"github.com/kasader/game-prototype/pkg/input"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	rkeyboard "github.com/hajimehoshi/ebiten/v2/examples/resources/images/keyboard"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var fontFace = text.NewGoXFace(bitmapfont.Face)

var keyboardImage *ebiten.Image

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	img, _, err := image.Decode(bytes.NewReader(rkeyboard.Keyboard_png))
	if err != nil {
		log.Fatal(err)
	}

	keyboardImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	state       *GameState
	keys        []ebiten.Key
	debugstring string
}

type GameState struct {
	Grid   *grid.Grid
	Player *player.Player
	Input  input.Input
	// TODO: Enemies, Items, NPCs, etc.
}

func (w *GameState) Update(input input.Input) {
	dx, dy := input.GetDirection()
	w.Player.TryMove(dx, dy, w.Grid)
}

func (g *Game) Update() error {
	if g.state == nil {
		g.state = &GameState{
			Grid:   grid.GetTestGrid(),
			Player: player.GetTestPlayer(),
			Input:  &input.EbitenInput{},
		}
	}
	g.state.Update(g.state.Input)

	/////////////////////////////
	// REMOVE ME LATER (DEBUG) //
	x_pos, y_pos := g.state.Player.GetPosition()
	g.debugstring = fmt.Sprintf("(X:%d, Y:%d)", x_pos, y_pos)
	/////////////////////////////
	return nil
}

var mplusFaceSource *text.GoTextFaceSource

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.debugstring)

	screenW, screenH := screen.Bounds().Dx(), screen.Bounds().Dy()
	lines := RenderGrid(g.state.Grid, g.state.Player)

	// Dynamic font size: fit grid to height
	numRows := len(lines)
	if numRows == 0 {
		return
	}

	// Use a monospaced assumption: fixed height per line
	fontSize := float64(screenH) / float64(numRows)

	// Approx character width-to-height ratio for most monospace fonts ~0.6
	charRatio := 0.6

	// Use widest line to determine width scaling
	maxCols := 0
	for _, line := range lines {
		if len(line) > maxCols {
			maxCols = len(line)
		}
	}

	// Adjust font size if it would overflow horizontally
	horizFitSize := float64(screenW) / (charRatio * float64(maxCols))
	if horizFitSize < fontSize {
		fontSize = horizFitSize
	}

	// Build font face
	fontFace := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   fontSize,
	}

	// Calculate baseline line height
	lineHeight := fontSize

	// Center grid block vertically
	startY := (float64(screenH) - lineHeight*float64(numRows)) / 2

	for i, line := range lines {
		// Width in pixels (estimated from char count and font size)
		linePixelWidth := charRatio * fontSize * float64(len(line))
		x := (float64(screenW) - linePixelWidth) / 2
		y := startY + float64(i)*lineHeight

		op := &text.DrawOptions{}
		op.GeoM.Translate(x, y)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, line, fontFace, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Keyboard (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func RenderGrid(grid *gamemap.Grid, player *player.Player) []string {
	var out []string
	for y := 0; y < grid.Height(); y++ {
		var row []string
		for x := 0; x < grid.Width(); x++ {
			if player.X == x && player.Y == y {
				row = append(row, "P")
			} else {
				row = append(row, grid.GetTile(x, y).String())
			}
		}
		out = append(out, "["+strings.Join(row, " ")+"]")
	}
	slices.Reverse(out)
	return out
}
