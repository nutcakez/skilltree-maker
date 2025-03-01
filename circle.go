package main

import (
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
		showBigRadius: true,
		x:             x,
		y:             y,
		radius:        radius,
		ownRadius:     10,
		strokeWidth:   5,
		circles:       make([]*circle, 0),
	}
}

func (c *circle) Update() {}

func (c *circle) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, c.x, c.y, c.ownRadius, red, true)
	for _, v := range c.circles {
		v.Draw(screen)
	}
}

func (c *circle) AddByRatio(a, b float32) {

}
