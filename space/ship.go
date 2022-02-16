package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Ship struct {
	Object
}

func NewShip() *Ship {
	return &Ship{
		Object: Object{
			X: float64(screenW/2) - float64(shipimg.Bounds().Dx()/2),
			Y: float64(screenH - shipimg.Bounds().Dy()),
			W: float64(shipimg.Bounds().Dx()),
			H: float64(shipimg.Bounds().Dy()),
		},
	}
}

func (s *Ship) Update() {
	if inpututil.KeyPressDuration(ebiten.KeyA) != 0 {
		s.X -= 5
		if s.X < 0 {
			s.X = 0
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyD) != 0 {
		s.X += 5
		if s.X+s.W > float64(screenW) {
			s.X = float64(screenW) - s.W
		}
	}
}

func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(shipimg.(*ebiten.Image), op)
}
