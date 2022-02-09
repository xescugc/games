package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	Object

	Speed int
	Lives int
}

func (e *Enemy) Update() {
	e.X += float64(e.Speed)
	if e.X < 0 {
		e.X = 0
		e.Speed = -e.Speed
	} else if e.X+float64(e.Image.Bounds().Dx()) > float64(screenW) {
		e.X = float64(screenW) - float64(e.Image.Bounds().Dx())
		e.Speed = -e.Speed
	}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(e.Image.(*ebiten.Image), op)
}

func (e *Enemy) TakeHit() {
	e.Lives -= 1
	if e.Speed > 0 {
		e.Speed += 5
	} else {
		e.Speed -= 5
	}
}
