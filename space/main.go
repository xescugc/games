package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/ship.png
var Ship_png []byte

//go:embed assets/alien.png
var Alien_png []byte

//go:embed assets/bullet.png
var Bullet_png []byte

var (
	screenW, screenH int

	shipimg, alienimg, bulletimg image.Image
)

func init() {
	si, _, err := image.Decode(bytes.NewReader(Ship_png))
	if err != nil {
		log.Fatal(err)
	}
	shipimg = ebiten.NewImageFromImage(si)

	ai, _, err := image.Decode(bytes.NewReader(Alien_png))
	if err != nil {
		log.Fatal(err)
	}
	alienimg = ebiten.NewImageFromImage(ai)

	bi, _, err := image.Decode(bytes.NewReader(Bullet_png))
	if err != nil {
		log.Fatal(err)
	}
	bulletimg = ebiten.NewImageFromImage(bi)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowTitle("Space")
	screenW, screenH = ebiten.WindowSize()
	g := &Game{
		Ship:   NewShip(),
		Aliens: NewAliens(),
	}
	if err := ebiten.RunGame(g); err != nil {
		if err == terminatedErr {
			return
		}
		log.Fatal(err)
	}
}
