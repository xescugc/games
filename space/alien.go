package main

import "github.com/hajimehoshi/ebiten/v2"

type Alien struct {
	Object
	Dead bool
}

func NewAlien(x, y float64) *Alien {
	return &Alien{
		Object: Object{
			X: x,
			Y: y,
			W: float64(alienimg.Bounds().Dx()),
			H: float64(alienimg.Bounds().Dy()),
		},
	}
}

func (a *Alien) Update() {
}

func (a *Alien) Draw(screen *ebiten.Image) {
	if a.Dead {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(a.X, a.Y)
	screen.DrawImage(alienimg.(*ebiten.Image), op)
}
