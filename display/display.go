package display

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nutcakez/skilltree-maker/skilltree"
)

type Display struct {
	posX, posY float64
	w, h       int
	visible    bool
	hoverText  string
	textFace   text.GoTextFace
	Node       *skilltree.Node
	SkillTree  *skilltree.SkillTree
	offscreen  *ebiten.Image
	panning    *Panning
}

func NewDisplay(posX, posY float64, width, height int) *Display {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}

	textFace := text.GoTextFace{
		Source: s,
		Size:   20,
	}
	display := &Display{
		posX:      posX,
		posY:      posY,
		w:         width,
		h:         height,
		visible:   false,
		textFace:  textFace,
		Node:      nil,
		offscreen: ebiten.NewImage(1000, 1000),
		panning:   NewPanning(),
		SkillTree: &skilltree.SkillTree{
			Nodes: make([]*skilltree.Node, 0),
		},
	}

	return display
}

func (d *Display) SetStartPosition(node *skilltree.Node) {
	var posX, posY float32
	if node == nil {
		posX = d.SkillTree.Nodes[0].X
		posY = d.SkillTree.Nodes[0].Y
	} else {
		for i := range d.SkillTree.Nodes {
			if d.SkillTree.Nodes[i] == node {
				posX = d.SkillTree.Nodes[i].X
				posY = d.SkillTree.Nodes[i].Y
			}
		}
	}

	midX := posX
	midY := posY
	d.panning.OffsetX = -int(midX - float32(d.w*int(d.panning.Zoom))/float32(2))
	d.panning.OffsetY = -int(midY - float32(d.h*int(d.panning.Zoom))/float32(2))
	fmt.Println("offsets", d.panning.OffsetX, d.panning.OffsetY)
}

func (d *Display) Update() {
	d.panning.Update()
	_, d.hoverText = d.SkillTree.Update(d.panning.OffsetX, d.panning.OffsetY, int(d.posX), int(d.posY), d.panning.Zoom)
}

func (d *Display) Draw(screen *ebiten.Image) {
	d.offscreen.Fill(color.RGBA{0, 0, 255, 200})
	d.SkillTree.Draw(d.offscreen)
	op := ebiten.DrawImageOptions{}
	// if we zoom in it means we want to see a bigger picture from screen2 but scaled to 500x500

	// figure out the scale here
	op.GeoM.Scale(1/d.panning.Zoom, 1/d.panning.Zoom)
	op.GeoM.Translate(float64(d.posX), float64(d.posY))

	// which part of the stuff we want
	rect := image.Rect(d.panning.OffsetX, d.panning.OffsetY, int(float64(d.w)*d.panning.Zoom), int(float64(d.h)*d.panning.Zoom))

	vector.StrokeRect(screen, 0+float32(d.posX), 0+float32(d.posY), float32(d.w), float32(d.h), 2, color.RGBA{255, 0, 0, 255}, true)

	screen.DrawImage(d.offscreen.SubImage(rect).(*ebiten.Image), &op)

	if d.hoverText != "" {
		d.drawHoverText(screen)
	}
}

func (d *Display) drawHoverText(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	w, h := text.Measure(d.hoverText, &d.textFace, 0)
	offsetX := 20
	offsetY := 20

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x)+float64(offsetX), float64(y)+float64(offsetY))
	op.LineSpacing = 20 * 1.2

	numOfLines := strings.Count(d.hoverText, "\n") + 1
	vector.DrawFilledRect(screen, float32(x+offsetX), float32(y+offsetY), float32(w), float32(h*float64(numOfLines)), color.RGBA{59, 59, 59, 255}, true)

	text.Draw(screen, d.hoverText, &d.textFace, op)
}
