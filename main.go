package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	red = color.RGBA{255, 0, 0, 1}
)

func main() {
	ebiten.SetWindowSize(1280, 960)
	game := game{}

	if err := ebiten.RunGame(&game); err != nil {
	}
}

type game struct{}

func (g *game) Update() error {

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	vector.StrokeCircle(screen, 10, 10, 10, 10, red, true)
}

func (g *game) Layout(a, b int) (int, int) { return 1280, 960 }
