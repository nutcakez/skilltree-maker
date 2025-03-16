package main

import (
	"fmt"
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
	skillTree  *SkillTree
	offscreen  *ebiten.Image
	panning    *Panning
}

func NewDisplay(posX, posY float64, width, height int) *Display {
	display := &Display{
		posX:      posX,
		posY:      posY,
		w:         width,
		h:         height,
		visible:   false,
		node:      nil,
		offscreen: ebiten.NewImage(4000, 4000),
		panning:   NewPanning(),
		skillTree: &SkillTree{
			nodes: make([]*Node, 0),
		},
	}

	return display
}

func (d *Display) SetStartPosition(node *Node) {
	var posX, posY float32
	if node == nil {
		posX = d.skillTree.nodes[0].x
		posY = d.skillTree.nodes[0].y
		// w = d.skillTree.nodes[0].img.Bounds().Dx()
		// h = d.skillTree.nodes[0].img.Bounds().Dy()
	} else {
		for i := range d.skillTree.nodes {
			if d.skillTree.nodes[i] == node {
				posX = d.skillTree.nodes[i].x
				posY = d.skillTree.nodes[i].y
				// w = d.skillTree.nodes[0].img.Bounds().Dx()
				// h = d.skillTree.nodes[0].img.Bounds().Dy()
			}
		}
	}

	midX := posX
	midY := posY
	d.panning.offsetX = -int(midX - float32(d.w*int(d.panning.zoom))/float32(2))
	d.panning.offsetY = -int(midY - float32(d.h*int(d.panning.zoom))/float32(2))
	fmt.Println("offsets", d.panning.offsetX, d.panning.offsetY)
}

func (d *Display) Update() {
	d.panning.Update()
	d.skillTree.Update(d.panning.offsetX, d.panning.offsetY, int(d.posX), int(d.posY), d.panning.zoom)
}

func (d *Display) Draw(screen *ebiten.Image) {
	d.offscreen.Fill(color.RGBA{0, 0, 255, 200})
	// d.node.Draw(d.offscreen)
	d.skillTree.Draw(d.offscreen)
	op := ebiten.DrawImageOptions{}
	// if we zoom in it means we want to see a bigger picture from screen2 but scaled to 500x500

	// figure out the scale here
	op.GeoM.Scale(1/d.panning.zoom, 1/d.panning.zoom)
	op.GeoM.Translate(float64(d.posX), float64(d.posY))

	// which part of the stuff we want?
	rect := image.Rect(d.panning.offsetX, d.panning.offsetY, int(float64(d.w)*d.panning.zoom), int(float64(d.h)*d.panning.zoom))

	vector.StrokeRect(screen, 0+float32(d.posX), 0+float32(d.posY), float32(d.w), float32(d.h), 2, color.RGBA{255, 0, 0, 255}, true)

	screen.DrawImage(d.offscreen.SubImage(rect).(*ebiten.Image), &op)
}
