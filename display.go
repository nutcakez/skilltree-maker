package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Display struct {
	posX, posY float64
	w, h       int
	visible    bool
	node       *Node
	offscreen  *ebiten.Image
	panning    *Panning
}

func NewDisplay(posX, posY float64, width, height int) *Display {
	return &Display{
		posX:      posX,
		posY:      posY,
		w:         width,
		h:         height,
		visible:   false,
		node:      nil,
		offscreen: ebiten.NewImage(1000, 1000),
		panning:   NewPanning(),
	}
}

func (d *Display) Update() {
	d.panning.Update()
	// g.objects[i].Update(g.panner.offsetX, g.panner.offsetY, g.screen2OffsetX, g.screen2OffsetY, g.panner.zoom)
	d.node.Update(d.panning.offsetX, d.panning.offsetY, int(d.posX), int(d.posY), d.panning.zoom)
	d.offscreen.Fill(color.RGBA{0, 0, 255, 200})
}

func (d *Display) Draw(screen *ebiten.Image) {
	d.node.Draw(d.offscreen)
	op := ebiten.DrawImageOptions{}
	// if we zoom in it means we want to see a bigger picture from screen2 but scaled to 500x500

	// figure out the scale here
	// op.GeoM.Scale(1/d.panning.zoom, 1/d.panning.zoom)
	op.GeoM.Translate(float64(d.posX), float64(d.posY))

	// which part of the stuff we want?
	rect := image.Rect(d.panning.offsetX, d.panning.offsetY, int(float64(d.w)*d.panning.zoom), int(float64(d.h)*d.panning.zoom))
	img := ebiten.NewImageFromImage(d.offscreen.SubImage(rect))

	vector.StrokeRect(screen, 0+float32(d.posX), 0+float32(d.posY), float32(d.w), float32(d.h), 2, color.RGBA{255, 0, 0, 255}, true)

	screen.DrawImage(img, &op)
}
