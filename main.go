package main

import (
	"fmt"
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

	images := ReadAllImageFromFolder("96x96")
	fmt.Println(images)

	radiusShows := true

	otherMainCircle := NewDefaultImgNode(300, 700, images[2])
	otherMainCircle.showBigRadius = radiusShows
	c1 := otherMainCircle.AddImgByDegree(45, images[3])
	c2 := c1.AddImgByDegree(45, images[4])

	c2.AddImgByDegree(45, images[5])
	c2.AddImgByDegree(90, images[7])
	c2.showBigRadius = radiusShows
	c3 := c1.AddImgByDegree(135, images[5])
	c3.AddImgByDegree(45, images[6])
	c3.AddImgByDegree(90, images[8])
	// mainCircle := NewDefaultCircle(300, 300, 100)
	// mainCircle.showBigRadius = radiusShows
	// mainCircle.
	// 	AddByDegree(45).
	// 	AddByDegree(135)
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

	// game.objects = append(game.objects, mainCircle)
	game.objects = append(game.objects, otherMainCircle)

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
