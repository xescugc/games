package main

import (
	"errors"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var terminatedErr = errors.New("terminated")

type Game struct {
	Snake *Snake
	Point *Point
}

func (g *Game) Update() error {
	if g.Snake.Dead {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return terminatedErr
		}
	}
	g.Snake.Update()
	g.Point.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenW, screenH = screen.Size()
	if g.Snake.Dead {
		ebitenutil.DebugPrintAt(screen, "You Lose!", (screenW/2)-20, screenH/2)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Points: %d", len(g.Snake.Parts)-1), (screenW/2)-20, (screenH/2)+20)
		ebitenutil.DebugPrintAt(screen, "(ESC to exit)", (screenW/2)-35, (screenH/2)+40)
		return
	}
	// The -1 is because the Head is the first element
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Points: %d", len(g.Snake.Parts)-1))
	g.Snake.Draw(screen)
	g.Point.Draw(screen)
	g.Snake.CheckPointCollision(g.Point)
	g.Snake.CheckSelfCollision()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
