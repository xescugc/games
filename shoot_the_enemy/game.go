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
	Player  *Player
	Enemy   *Enemy
	Bullets []*Bullet
}

func (g *Game) Update() error {
	if g.Enemy.Lives <= 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return terminatedErr
		}
	}
	g.Player.Update(g)
	g.Enemy.Update()

	removeBullets := make([]int, 0, 0)
	for i, b := range g.Bullets {
		if b.Y < 0 || b.Dead {
			removeBullets = append(removeBullets, i)
			continue
		}
		b.Update(g)
		b.CheckCollision(g.Enemy)
	}

	for i, idx := range removeBullets {
		g.Bullets = append(g.Bullets[:idx+i], g.Bullets[idx+i+1:]...)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenW, screenH = screen.Size()
	if g.Enemy.Lives <= 0 {
		ebitenutil.DebugPrintAt(screen, "You Won!", (screenW/2)-20, screenH/2)
		ebitenutil.DebugPrintAt(screen, "(ESC to exit)", (screenW/2)-35, (screenH/2)+20)
		return
	}
	g.Player.Draw(screen)
	g.Enemy.Draw(screen)

	for _, b := range g.Bullets {
		b.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Lives: %d", g.Enemy.Lives))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
