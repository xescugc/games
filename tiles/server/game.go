package main

import (
	"github.com/xescugc/games/tiles/player"
)

type Game struct {
	Screen  Screen
	Players map[string]*player.Player
}

type Screen struct {
	W, H int
}

func NewGame(sid string) *Game {
	g := &Game{
		Screen: Screen{
			H: 240,
			W: 240,
		},
		Players: make(map[string]*player.Player),
	}
	g.AddPlayer(sid)
	return g
}

func (g *Game) AddPlayer(sid string) {
	w := 16
	g.Players[sid] = &player.Player{
		ID: sid,
		Object: player.Object{
			X: float64(g.Screen.W/2) - float64(w/2),
			Y: float64(g.Screen.H - w),
			W: float64(w),
			H: float64(w),
		},
		Facing: "W",
	}
}

func (g *Game) RemovePlayer(sid string) {
	delete(g.Players, sid)
}

func (g *Game) MovePlayer(sid, dir string) {
	for _, p := range g.Players {
		if p.ID != sid {
			continue
		}
		p.Moving = true
		p.MovingCount += 1
		switch dir {
		case "S":
			p.Facing = "S"
			p.Y += 1
			if p.Y > (float64(g.Screen.H) - p.H) {
				p.Y = float64(g.Screen.H) - p.H
			}
		case "W":
			p.Facing = "W"
			p.Y -= 1
			if p.Y < 0 {
				p.Y = 0
			}
		case "A":
			p.Facing = "A"
			p.X -= 1
			if p.X < 0 {
				p.X = 0
			}
		case "D":
			p.Facing = "D"
			p.X += 1
			if p.X > (float64(g.Screen.W) - p.W) {
				p.X = float64(g.Screen.W) - p.W
			}
		default:
			p.Moving = false
			p.MovingCount = 0
		}

	}
}
