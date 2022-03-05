package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Columns   int
	TileSize  int
	Tiles     []int
	Players   map[string]*Player
	CurrentID string
}

func (g *Game) Update() error {
	if p, ok := g.Players[g.CurrentID]; ok {
		p.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var xNum = screenW / g.TileSize
	for i, t := range g.Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%xNum)*g.TileSize), float64((i/xNum)*g.TileSize))

		sx := (t % g.Columns) * g.TileSize
		sy := (t / g.Columns) * g.TileSize
		screen.DrawImage(tilesetImg.(*ebiten.Image).SubImage(image.Rect(sx, sy, sx+g.TileSize, sy+g.TileSize)).(*ebiten.Image), op)
	}
	for _, p := range g.Players {
		p.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
