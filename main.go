package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	red    = color.RGBA{255, 0, 0, 255}
	green  = color.RGBA{0, 255, 0, 1}
	orange = color.RGBA{255, 127, 80, 1}
)

func main() {
	ebiten.SetWindowSize(1280, 960)
	game := game{
		objects: make([]IDrawUpdate, 0),
	}

	radiusShows := false

	mainCircle := NewDefaultCircle(300, 300, 100)
	mainCircle.showBigRadius = radiusShows
	mainCircle.AddByDegree(45)
	mainCircle.AddByDegree(135)
	// mainCircle.AddByRatio(2, 4)
	// mainCircle.AddByRatio(1, 8)
	// mainCircle.AddByRatio(2, 8)
	// c1 := mainCircle.AddByRatio(3, 8)
	// c1.radius = 150
	// c1.showBigRadius = radiusShows
	// c2 := c1.AddByDegree(90)
	// c2.radius = 200
	// c2.AddByDegree(90)
	// c2.AddByDegree(130)
	// c2.AddByDegree(170)
	// c1.AddByDegree(180)
	// c1.AddByDegree(220)
	// c1.AddByDegree(240)
	// mainCircle.AddByRatio(4, 8)
	// mainCircle.AddByRatio(5, 8)
	// mainCircle.AddByRatio(6, 8)
	// mainCircle.AddByRatio(7, 8)
	// mainCircle.AddByRatio(8, 8)

	game.objects = append(game.objects, mainCircle)

	if err := ebiten.RunGame(&game); err != nil {
	}
}

type game struct {
	objects []IDrawUpdate
}

func (g *game) Update() error {

	for i := range g.objects {
		g.objects[i].Update()
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for i := range g.objects {
		g.objects[i].Draw(screen)
	}
}

func (g *game) Layout(a, b int) (int, int) { return 1280, 960 }

type IDrawUpdate interface {
	Draw(*ebiten.Image)
	Update()
}
