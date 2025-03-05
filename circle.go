package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Node struct {
	showBigRadius, active                bool
	x, y, ownRadius, radius, strokeWidth float32
	childs                               []*Node
	img                                  *ebiten.Image
}

func NewDefaultCircle(x, y, radius float32) *Node {
	return &Node{
		showBigRadius: false,
		active:        false,
		x:             x,
		y:             y,
		radius:        radius,
		ownRadius:     10,
		strokeWidth:   2,
		childs:        make([]*Node, 0),
		img:           nil,
	}
}

func NewDefaultImgNode(x, y float32, img *ebiten.Image) *Node {
	return &Node{
		showBigRadius: false,
		active:        false,
		x:             x,
		y:             y,
		radius:        300,
		ownRadius:     10,
		strokeWidth:   2,
		childs:        make([]*Node, 0),
		img:           img,
	}
}

func (c *Node) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		if !c.checkClick() {
			for i := range c.childs {
				c.childs[i].Update()
			}
		}
	}
}

func (c *Node) checkClick() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		if c.img != nil {
			w := c.img.Bounds().Dx()
			h := c.img.Bounds().Dy()
			if PointInRect(x, y, int(c.x-float32(w/2)), int(c.y-float32(h/2)), w, h) {
				c.active = !c.active
				return true
			}
		} else {
			if PointInCircle(c.x, c.y, c.ownRadius, int32(x), int32(y)) {
				c.active = !c.active
				return true
			}
		}
	}

	return false

}

func (c *Node) Draw(screen *ebiten.Image) {
	for _, v := range c.childs {
		vector.StrokeLine(screen, c.x, c.y, v.x, v.y, 2, orange, true)
	}

	var color color.RGBA

	if c.active {
		color = green
	} else {
		color = red
	}

	if c.img == nil {
		vector.DrawFilledCircle(screen, c.x, c.y, c.ownRadius, color, true)
	} else {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.x-float32(c.img.Bounds().Dx()/2)), float64(c.y-float32(c.img.Bounds().Dy()/2)))
		screen.DrawImage(c.img, &op)
		strokeWith := 5
		vector.StrokeRect(screen,
			float32(c.x-float32(c.img.Bounds().Dx()/2)),
			float32(c.y-float32(c.img.Bounds().Dy()/2)),
			float32(c.img.Bounds().Dx()),
			float32(c.img.Bounds().Dy()),
			float32(strokeWith),
			color,
			true,
		)
	}
	if c.showBigRadius {
		vector.StrokeCircle(screen, c.x, c.y, c.radius, c.strokeWidth, green, true)

	}
	for _, v := range c.childs {
		v.Draw(screen)
	}
}

func (c *Node) AddByDegree(degree float64) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(float64(c.x), float64(c.y), float64(c.radius), degree)

	newCircle := NewDefaultCircle(float32(newX), float32(newY), 50)

	c.childs = append(c.childs, newCircle)

	return newCircle
}

func (c *Node) AddImgByDegree(degree float64, img *ebiten.Image) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(float64(c.x), float64(c.y), float64(c.radius), degree)

	newNode := NewDefaultImgNode(float32(newX), float32(newY), img)

	c.childs = append(c.childs, newNode)

	return newNode
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
