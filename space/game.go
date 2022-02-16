package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var terminatedErr = errors.New("terminated")

type Game struct {
	Ship    *Ship
	Aliens  *Aliens
	Bullets []*Bullet
}

type Object struct {
	X, Y, W, H float64
}

func (o Object) IsColliding(c Object) bool {
	selfLeft := o.X
	selfRight := o.X + o.W
	selfTop := o.Y
	selfBottom := o.Y + o.H

	enemyLeft := c.X
	enemyRight := c.X + c.W
	enemyTop := c.Y
	enemyBottom := c.Y + c.H

	return selfRight > enemyLeft && selfLeft < enemyRight && selfBottom > enemyTop && selfTop < enemyBottom
}

func (g *Game) Update() error {
	if len(g.Aliens.Aliens) == 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return terminatedErr
		}
	}
	g.Ship.Update()
	g.Aliens.Update()

	removeBullets := make([]int, 0, 0)
	for i, b := range g.Bullets {
		if b.Dead {
			removeBullets = append(removeBullets, i)
			continue
		}
		b.Update()
		for _, a := range g.Aliens.Aliens {
			if b.IsColliding(a.Object) {
				b.Dead = true
				a.Dead = true
			}
		}
	}

	for i, idx := range removeBullets {
		g.Bullets = append(g.Bullets[:idx+i], g.Bullets[idx+i+1:]...)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Bullets = append(g.Bullets, NewBullet(g.Ship))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if len(g.Aliens.Aliens) == 0 {
		ebitenutil.DebugPrintAt(screen, "You Won!", (screenW/2)-20, screenH/2)
		ebitenutil.DebugPrintAt(screen, "(ESC to exit)", (screenW/2)-35, (screenH/2)+20)
		return
	}
	g.Ship.Draw(screen)
	g.Aliens.Draw(screen)
	for _, b := range g.Bullets {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
