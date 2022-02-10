package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Red = color.RGBA{R: 255, G: 0, B: 0, A: 255}

type Point struct {
	X, Y, W, H float64
	Hit        bool
}

func NewPoint() *Point {
	return &Point{
		X: float64(rand.Intn(screenW + 1)),
		Y: float64(rand.Intn(screenH + 1)),
		W: 10, H: 10,
	}
}

func (p *Point) Update() {
	if p.Hit {
		p.X = float64(rand.Intn(screenW + 1))
		p.Y = float64(rand.Intn(screenH + 1))
		p.Hit = false
	}
}

func (p *Point) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.X, p.Y, p.W, p.H, Red)
}
