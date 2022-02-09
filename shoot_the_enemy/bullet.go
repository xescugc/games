package main

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	Object

	Speed int
	Dead  bool
}

func (b *Bullet) Update(g *Game) {
	b.Y += float64(b.Speed)
	if b.Y < 0 {
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.Image.(*ebiten.Image), op)
}

func (b *Bullet) CheckCollision(e *Enemy) {
	selfLeft := b.X
	selfRight := b.X + float64(b.Image.Bounds().Dx())
	selfTop := b.Y
	selfBottom := b.Y + float64(b.Image.Bounds().Dy())

	enemyLeft := e.X
	enemyRight := e.X + float64(e.Image.Bounds().Dx())
	enemyTop := e.Y
	enemyBottom := e.Y + float64(e.Image.Bounds().Dy())

	if selfRight > enemyLeft && selfLeft < enemyRight && selfBottom > enemyTop && selfTop < enemyBottom {
		b.Dead = true
		e.TakeHit()
	}
}
