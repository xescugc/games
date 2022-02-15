package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenW = 420
	screenH = 340
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowTitle("Snake")
	ebiten.SetWindowResizable(true)
	screenW, screenH = ebiten.ScreenSizeInFullscreen()
	g := &Game{
		Snake: &Snake{
			Parts: []*Part{
				NewPart(float64(screenW)/2, float64(screenH)/2, ebiten.KeyW),
			},
		},
		Point: NewPoint(),
	}
	if err := ebiten.RunGame(g); err != nil {
		if err == terminatedErr {
			return
		}
		log.Fatal(err)
	}
}
