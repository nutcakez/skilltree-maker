package main

import "github.com/hajimehoshi/ebiten/v2"

type Display struct {
	posX, posY float64
	w, h       int
	visible    bool
}

func (d *Display) Update() {}

func (d *Display) Draw(screen *ebiten.Image) {}
