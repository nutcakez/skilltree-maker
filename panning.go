package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Panning struct {
	offsetX, offsetY int
	prevX, prevY     int
	zoom             float64
	pressed          bool
	changed          bool
}

func NewPanning() *Panning {
	return &Panning{
		zoom: 2,
	}
}

func (p *Panning) Update() {
	fmt.Println("panner", p.offsetX, p.offsetY)
	_, wheelY := ebiten.Wheel()
	p.changed = false

	if wheelY > 0 {
		p.zoom += 0.25
		p.changed = true
	}
	if wheelY < 0 {
		p.zoom -= 0.25
		if p.zoom <= 0.2 {
			p.zoom = 0.2
		}
		p.changed = true
	}
	if p.changed {
		fmt.Println(p.zoom)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.offsetX, p.offsetY = 0, 0
	}
	if p.pressed {
		x, y := ebiten.CursorPosition()
		p.offsetX -= int(float64(p.prevX-x) * p.zoom)
		p.offsetY -= int(float64(p.prevY-y) * p.zoom)

		p.prevX = x
		p.prevY = y
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		p.pressed = true
		p.prevX, p.prevY = ebiten.CursorPosition()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		p.pressed = false
	}

	// offset maxing:
	if p.offsetX > 0 {
		p.offsetX = 0
	}
	if p.offsetY > 0 {
		p.offsetY = 0
	}
}

func (p *Panning) Draw(screen *ebiten.Image) {}
