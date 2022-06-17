package engine

import "github.com/hajimehoshi/ebiten/v2"

type EbitenController struct{}

func (EbitenController) Left() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft)
}

func (EbitenController) Right() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight)
}

func (EbitenController) Jump() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace)
}
