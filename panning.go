package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Panning struct {
	OffsetX, OffsetY int
	PrevX, PrevY     int
	Zoom             float64
	Pressed          bool
	Changed          bool
}

func NewPanning() *Panning {
	return &Panning{
		Zoom: 2,
	}
}

func (p *Panning) Update() {
	fmt.Println("panner", p.OffsetX, p.OffsetY)
	_, wheelY := ebiten.Wheel()
	p.Changed = false

	if wheelY > 0 {
		p.Zoom += 0.25
		p.Changed = true
	}
	if wheelY < 0 {
		p.Zoom -= 0.25
		if p.Zoom <= 0.2 {
			p.Zoom = 0.2
		}
		p.Changed = true
	}
	if p.Changed {
		fmt.Println(p.Zoom)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.OffsetX, p.OffsetY = 0, 0
	}
	if p.Pressed {
		x, y := ebiten.CursorPosition()
		p.OffsetX -= int(float64(p.PrevX-x) * p.Zoom)
		p.OffsetY -= int(float64(p.PrevY-y) * p.Zoom)

		p.PrevX = x
		p.PrevY = y
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		p.Pressed = true
		p.PrevX, p.PrevY = ebiten.CursorPosition()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		p.Pressed = false
	}

	// offset maxing:
	if p.OffsetX > 0 {
		p.OffsetX = 0
	}
	if p.OffsetY > 0 {
		p.OffsetY = 0
	}
}

func (p *Panning) Draw(screen *ebiten.Image) {}
