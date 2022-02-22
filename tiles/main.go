package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/TilesetFloor.png
var Tileset_png []byte

//go:embed assets/Walk.png
var Walk_png []byte

var (
	tilesetImg image.Image
	playerImg  image.Image
)

const (
	screenW = 240
	screenH = 240
)

func init() {
	tsi, _, err := image.Decode(bytes.NewReader(Tileset_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesetImg = ebiten.NewImageFromImage(tsi)

	wi, _, err := image.Decode(bytes.NewReader(Walk_png))
	if err != nil {
		log.Fatal(err)
	}
	playerImg = ebiten.NewImageFromImage(wi)
}

func main() {
	ebiten.SetWindowTitle("Tiles")
	ebiten.SetWindowSize(screenW*2, screenH*2)
	g := &Game{
		Columns:  22,
		TileSize: 16,
		Tiles: []int{
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 154, 204, 203, 156, 264, 264, 264, 264, 264, 264,
			155, 155, 155, 155, 155, 204, 177, 177, 203, 155, 155, 155, 155, 155, 155,
			199, 199, 199, 199, 199, 182, 177, 177, 181, 199, 199, 199, 199, 199, 199,
			264, 264, 264, 264, 264, 198, 182, 181, 200, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
		},
		Player: NewPlayer(),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
