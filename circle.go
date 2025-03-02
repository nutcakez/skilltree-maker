package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type circle struct {
	showBigRadius                        bool
	x, y, ownRadius, radius, strokeWidth float32
	circles                              []*circle
}

func NewDefaultCircle(x, y, radius float32) *circle {
	return &circle{
		showBigRadius: false,
		x:             x,
		y:             y,
		radius:        radius,
		ownRadius:     10,
		strokeWidth:   2,
		circles:       make([]*circle, 0),
	}
}

func (c *circle) Update() {}

func (c *circle) Draw(screen *ebiten.Image) {
	for _, v := range c.circles {
		vector.StrokeLine(screen, c.x, c.y, v.x, v.y, 2, orange, true)
	}

	vector.DrawFilledCircle(screen, c.x, c.y, c.ownRadius, red, true)
	if c.showBigRadius {
		vector.StrokeCircle(screen, c.x, c.y, c.radius, c.strokeWidth, green, true)

	}
	for _, v := range c.circles {
		v.Draw(screen)
	}
}

func (c *circle) AddByDegree(degree float64) *circle {
	newX, newY := GetPointOnCircle(float64(c.x), float64(c.y), float64(c.radius), degree)

	newCircle := NewDefaultCircle(float32(newX), float32(newY), 50)

	c.circles = append(c.circles, newCircle)

	return newCircle
}

func (c *circle) AddByRatio(a, b float32) *circle {
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
