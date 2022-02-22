package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	facingToTile = map[ebiten.Key]int{
		ebiten.KeyS: 0,
		ebiten.KeyW: 1,
		ebiten.KeyA: 2,
		ebiten.KeyD: 3,
	}
)

type Player struct {
	Object
	Facing      ebiten.Key
	Moving      bool
	MovingCount int
}

func NewPlayer() *Player {
	w := 16
	return &Player{
		Object: Object{
			X: float64(screenW/2) - float64(w/2),
			Y: float64(screenH - w),
			W: float64(w),
			H: float64(w),
		},
		Facing: ebiten.KeyW,
	}
}

func (p *Player) Update() {
	p.Moving = true
	p.MovingCount += 1
	if inpututil.KeyPressDuration(ebiten.KeyW) != 0 {
		p.Facing = ebiten.KeyW
		p.Y -= 1
		if p.Y < 0 {
			p.Y = 0
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyS) != 0 {
		p.Facing = ebiten.KeyS
		p.Y += 1
		if p.Y > (screenH - p.H) {
			p.Y = screenH - p.H
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyD) != 0 {
		p.Facing = ebiten.KeyD
		p.X += 1
		if p.X > (screenW - p.W) {
			p.X = screenW - p.W
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyA) != 0 {
		p.Facing = ebiten.KeyA
		p.X -= 1
		if p.X < 0 {
			p.X = 0
		}
	} else {
		p.Moving = false
		p.MovingCount = 0
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	sx := facingToTile[p.Facing] * int(p.W)
	i := (p.MovingCount / 5) % 4
	sy := i * int(p.H)
	screen.DrawImage(playerImg.(*ebiten.Image).SubImage(image.Rect(sx, sy, sx+int(p.W), sy+int(p.H))).(*ebiten.Image), op)
}
