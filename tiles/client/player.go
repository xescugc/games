package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/xescugc/games/tiles/message"
	"github.com/xescugc/games/tiles/player"
)

var (
	facingToTile = map[string]int{
		ebiten.KeyS.String(): 0,
		ebiten.KeyW.String(): 1,
		ebiten.KeyA.String(): 2,
		ebiten.KeyD.String(): 3,
	}
)

type Player struct {
	player.Player
}

func (p *Player) Update() {
	if inpututil.KeyPressDuration(ebiten.KeyW) != 0 {
		err := wsc.WriteJSON(message.NewMoveMessage(p.ID, ebiten.KeyW.String()))
		if err != nil {
			log.Fatal(err)
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyS) != 0 {
		err := wsc.WriteJSON(message.NewMoveMessage(p.ID, ebiten.KeyS.String()))
		if err != nil {
			log.Fatal(err)
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyD) != 0 {
		err := wsc.WriteJSON(message.NewMoveMessage(p.ID, ebiten.KeyD.String()))
		if err != nil {
			log.Fatal(err)
		}
	} else if inpututil.KeyPressDuration(ebiten.KeyA) != 0 {
		err := wsc.WriteJSON(message.NewMoveMessage(p.ID, ebiten.KeyA.String()))
		if err != nil {
			log.Fatal(err)
		}
		// We only want to say that it has stoppend when it's moving
	} else if p.Moving {
		err := wsc.WriteJSON(message.NewMoveMessage(p.ID, ""))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	sx := facingToTile[p.Facing] * int(p.W)
	i := (p.MovingCount / 5) % 4
	sy := i * int(p.H)
	screen.DrawImage(playerImg.(*ebiten.Image).SubImage(image.Rect(sx, sy, sx+int(p.W), sy+int(p.H))).(*ebiten.Image), op)
}
