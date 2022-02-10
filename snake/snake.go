package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var Green = color.RGBA{R: 0, G: 255, B: 0, A: 255}

// This has to match the W/H of the Parts as then
// each Draw we advance 1 Part and so the logic of
// turning would work, if not then we would need to
// keep track when the change of direction was done
// so the other parts know when to turn
var SnakeSpeed float64 = 5
var directions = map[ebiten.Key]func(*Part){
	ebiten.KeyW: func(p *Part) {
		p.Y -= SnakeSpeed
	},
	ebiten.KeyS: func(p *Part) {
		p.Y += SnakeSpeed
	},
	ebiten.KeyA: func(p *Part) {
		p.X -= SnakeSpeed
	},
	ebiten.KeyD: func(p *Part) {
		p.X += SnakeSpeed
	},
}

var opositeDirection = map[ebiten.Key]ebiten.Key{
	ebiten.KeyW: ebiten.KeyS,
	ebiten.KeyS: ebiten.KeyW,
	ebiten.KeyA: ebiten.KeyD,
	ebiten.KeyD: ebiten.KeyA,
}
var arrowsToKeys = map[ebiten.Key]ebiten.Key{
	ebiten.KeyArrowDown:  ebiten.KeyS,
	ebiten.KeyArrowUp:    ebiten.KeyW,
	ebiten.KeyArrowRight: ebiten.KeyD,
	ebiten.KeyArrowLeft:  ebiten.KeyA,
}

type Snake struct {
	Parts []*Part
	Dead  bool
}

type Part struct {
	X, Y, W, H float64
	Direction  ebiten.Key // W,A,S,D
	New        bool
}

func NewPart(x, y float64, d ebiten.Key) *Part {
	return &Part{
		X: x, Y: y,
		W: 5, H: 5,
		Direction: d,
	}
}

func (s *Snake) Update() {
	direction := s.Head().Direction

	for k, v := range arrowsToKeys {
		if inpututil.IsKeyJustPressed(k) || inpututil.IsKeyJustPressed(v) {
			direction = v
			break
		}
	}

	h := s.Head()
	// We cannot go into ourselves
	if opo := opositeDirection[h.Direction]; opo == direction {
		direction = h.Direction
	}

	// We need to seave the X,Y in which the instruction has to be applied
	parentDirection := direction
	for _, p := range s.Parts {
		// The new Parts will not move
		// and will move on the next
		// call as this will make it seams
		// like the tail did grow
		if p.New {
			p.New = false
			continue
		}
		dirFn, ok := directions[parentDirection]
		// If we do not handle we do not move
		if !ok {
			continue
		}
		dirFn(p)
		aux := p.Direction
		p.Direction = parentDirection
		parentDirection = aux

		if p.X < 0 {
			p.X = float64(screenW)
		} else if p.X > float64(screenW) {
			p.X = 0
		} else if p.Y < 0 {
			p.Y = float64(screenH)
		} else if p.Y > float64(screenH) {
			p.Y = 0
		}
	}

}

func (s *Snake) Head() *Part { return s.Parts[0] }
func (s *Snake) Tail() *Part { return s.Parts[len(s.Parts)-1] }

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, p := range s.Parts {
		ebitenutil.DrawRect(screen, p.X, p.Y, p.W, p.H, Green)
	}
}

func (s *Snake) CheckPointCollision(p *Point) {
	h := s.Head()
	selfLeft := h.X
	selfRight := h.X + h.W
	selfTop := h.Y
	selfBottom := h.Y + h.H

	enemyLeft := p.X
	enemyRight := p.X + p.W
	enemyTop := p.Y
	enemyBottom := p.Y + p.H

	if selfRight > enemyLeft && selfLeft < enemyRight && selfBottom > enemyTop && selfTop < enemyBottom {
		p.Hit = true
		s.AddPoint()
	}
}

func (s *Snake) AddPoint() {
	t := s.Tail()
	s.Parts = append(s.Parts, NewPart(t.X, t.Y, t.Direction))
	s.Tail().New = true
}

func (s *Snake) CheckSelfCollision() {
	for i, p := range s.Parts {
		// We skip the head and the new elements
		// as they are placed below the last one
		if i == 0 || p.New {
			continue
		}
		h := s.Head()
		selfLeft := h.X
		selfRight := h.X + h.W
		selfTop := h.Y
		selfBottom := h.Y + h.H

		enemyLeft := p.X
		enemyRight := p.X + p.W
		enemyTop := p.Y
		enemyBottom := p.Y + p.H

		if selfRight > enemyLeft && selfLeft < enemyRight && selfBottom > enemyTop && selfTop < enemyBottom {
			s.Dead = true
			break
		}
	}
}
