package main

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Object
}

func (p *Player) Update(g *Game) {
	if inpututil.KeyPressDuration(ebiten.KeyA) != 0 {
		p.X -= 5
		if p.X < 0 {
			p.X = 0
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyD) != 0 {
		p.X += 5
		if p.X+float64(p.Image.Bounds().Dx()) > float64(screenW) {
			p.X = float64(screenW) - float64(p.Image.Bounds().Dx())
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		bulletImg, _, err := image.Decode(bytes.NewReader(Bullet_png))
		if err != nil {
			log.Fatal(err)
		}

		g.Bullets = append(g.Bullets, &Bullet{
			Object: Object{
				X: p.X + float64(p.Image.Bounds().Dx()/2), Y: p.Y,
				Image: ebiten.NewImageFromImage(bulletImg),
			},
			Speed: -15,
		})
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Image.(*ebiten.Image), op)
}
