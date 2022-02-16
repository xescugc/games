package main

import "github.com/hajimehoshi/ebiten/v2"

type Direction string

const (
	NumOfXAliens = 6
	NumOfYAliens = 3
)

type Aliens struct {
	Aliens []*Alien
	Speed  float64
}

func NewAliens() *Aliens {
	as := &Aliens{
		Aliens: make([]*Alien, 0, 18),
		Speed:  1,
	}

	// We want to have Aliens on the screen separated equally
	// so the 13 is the right number of odds so the even numbers
	// will be considered spaces
	sep := float64(screenW) / (NumOfXAliens*2 + 1)
	// Rows
	for i := 0; i < NumOfYAliens*3; i++ {
		if i%2 == 0 {
			continue
		}
		// Columns
		for y := 0; y < NumOfXAliens*2; y++ {
			if y%2 == 0 {
				continue
			}

			as.Aliens = append(as.Aliens, NewAlien(float64(y)*sep, float64(i*alienimg.Bounds().Dy())))
		}
	}
	return as
}

func (as *Aliens) IsRight() bool { return as.Speed > 0 }
func (as *Aliens) IsLeft() bool  { return as.Speed < 0 }

func (as *Aliens) Update() {
	changeDir := false
	removeAliens := make([]int, 0, 0)
	for i, a := range as.Aliens {
		if a.Dead {
			removeAliens = append(removeAliens, i)
			continue
		}
		a.X += as.Speed
		if (a.X+a.W) >= float64(screenW) && as.IsRight() {
			changeDir = true
		} else if a.X <= 0 && as.IsLeft() {
			changeDir = true
		}
		a.Update()
	}

	for i, idx := range removeAliens {
		as.Aliens = append(as.Aliens[:idx+i], as.Aliens[idx+i+1:]...)
	}

	if changeDir {
		as.Speed *= -1

		for _, a := range as.Aliens {
			a.Y += a.H / 2
			a.Update()
		}
	}
}

func (as *Aliens) Draw(screen *ebiten.Image) {
	for _, a := range as.Aliens {
		a.Draw(screen)
	}
}
