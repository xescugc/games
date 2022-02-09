package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/panda.png
var Panda_png []byte

//go:embed assets/snake.png
var Snake_png []byte

//go:embed assets/bullet.png
var Bullet_png []byte

var (
	screenW = 420
	screenH = 340
)

type Object struct {
	X, Y float64
	W, H float64

	Image image.Image
}

func main() {
	ebiten.SetWindowTitle("Shoot the enemy")
	ebiten.SetWindowResizable(true)
	screenW, screenH = ebiten.WindowSize()

	playerImg, _, err := image.Decode(bytes.NewReader(Panda_png))
	if err != nil {
		log.Fatal(err)
	}

	enemyImg, _, err := image.Decode(bytes.NewReader(Snake_png))
	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		Player: &Player{
			Object: Object{
				X: 0, Y: float64(screenH - playerImg.Bounds().Dy()),
				W: 10, H: 10,
				Image: ebiten.NewImageFromImage(playerImg),
			},
		},
		Enemy: &Enemy{
			Object: Object{
				X: 0, Y: 0,
				W: 10, H: 10,
				Image: ebiten.NewImageFromImage(enemyImg),
			},
			Speed: 7,
			Lives: 10,
		},
		Bullets: make([]*Bullet, 0, 0),
	}
	if err := ebiten.RunGame(g); err != nil {
		if err == terminatedErr {
			return
		}
		log.Fatal(err)
	}
}
