package main

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	Object
	Speed float64
	Dead  bool
}

func NewBullet(s *Ship) *Bullet {
	w := float64(bulletimg.Bounds().Dx())
	h := float64(bulletimg.Bounds().Dy())
	return &Bullet{
		Object: Object{
			X: s.X + s.W/2 - w/2,
			Y: s.Y + h,
			W: w,
			H: h,
		},
		Speed: -15,
	}
}

func (b *Bullet) Update() {
	b.Y += b.Speed
	if b.Y < 0 {
		b.Dead = true
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(bulletimg.(*ebiten.Image), op)
}
