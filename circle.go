package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Node struct {
	showBigRadius                        bool
	x, y, ownRadius, radius, strokeWidth float32
	circles                              []*Node
	img                                  *ebiten.Image
}

func NewDefaultCircle(x, y, radius float32) *Node {
	return &Node{
		showBigRadius: false,
		x:             x,
		y:             y,
		radius:        radius,
		ownRadius:     10,
		strokeWidth:   2,
		circles:       make([]*Node, 0),
		img:           nil,
	}
}

func (c *Node) Update() {}

func (c *Node) Draw(screen *ebiten.Image) {
	for _, v := range c.circles {
		vector.StrokeLine(screen, c.x, c.y, v.x, v.y, 2, orange, true)
	}

	if c.img == nil {
		vector.DrawFilledCircle(screen, c.x, c.y, c.ownRadius, red, true)
	} else {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.x), float64(c.y))
		screen.DrawImage(c.img, &op)
	}
	if c.showBigRadius {
		vector.StrokeCircle(screen, c.x, c.y, c.radius, c.strokeWidth, green, true)

	}
	for _, v := range c.circles {
		v.Draw(screen)
	}
}

func (c *Node) AddByDegree(degree float64) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(float64(c.x), float64(c.y), float64(c.radius), degree)

	newCircle := NewDefaultCircle(float32(newX), float32(newY), 50)

	c.circles = append(c.circles, newCircle)

	return newCircle
}

func (c *Node) AddByRatio(a, b float32) *Node {
	ratio := a / b
	d := 360 * ratio

	return c.AddByDegree(float64(d))
}

func GetPointOnCircle(cx, cy, radius, degree float64) (float64, float64) {
	// Convert degrees to radians
	radians := degree * (math.Pi / 180)

	// Calculate new coordinates
	newX := cx + radius*math.Cos(radians)
	newY := cy + radius*math.Sin(radians)

	fmt.Println(newX, newY)
	return newX, newY
}
