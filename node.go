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
	showBigRadius, active, startNode     bool
	x, y, ownRadius, radius, strokeWidth float32
	offsetX, offsetY                     int
	childs                               []*Node
	parents                              []*Node
	img                                  *ebiten.Image
	onClickEffect                        func()
	requirement                          func() bool
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
		parents:       make([]*Node, 0),
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

func (c *Node) Update(offsetX, offsetY, windowOffsetX, windowOffsetY int, zoom float64) {
	c.offsetX = offsetX
	c.offsetY = offsetY
	if !c.checkClick(zoom, windowOffsetX, windowOffsetY) {
		for i := range c.childs {
			c.childs[i].Update(offsetX, offsetY, windowOffsetX, windowOffsetY, zoom)
		}
	}
}

func (c *Node) checkClick(zoom float64, windowOffsetX, windowOffsetY int) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		if c.img != nil {
			w := c.img.Bounds().Dx()
			h := c.img.Bounds().Dy()
			if PointInRect(x, y,
				int(float64(c.x-float32(w/2)+float32(c.offsetX))/zoom+float64(windowOffsetX)),
				int(float64(c.y-float32(h/2)+float32(c.offsetY))/zoom+float64(windowOffsetY)),
				int(float64(w)/zoom),
				int(float64(h)/zoom)) && c.CanBeActivated() {
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
		// with antialias false its a bit better when zoomed out, should be set to false only when its zoomed out?
		vector.StrokeLine(screen, c.x+float32(c.offsetX), c.y+float32(c.offsetY), v.x+float32(c.offsetX), v.y+float32(c.offsetY), 2, orange, false)
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
		op.GeoM.Translate(float64(c.x-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)), float64(c.y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)))
		screen.DrawImage(c.img, &op)
		strokeWith := 5
		vector.StrokeRect(screen,
			float32(c.x-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)),
			float32(c.y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)),
			float32(c.img.Bounds().Dx()),
			float32(c.img.Bounds().Dy()),
			float32(strokeWith),
			color,
			true,
		)
	}
	if c.showBigRadius {
		vector.StrokeCircle(screen, c.x+float32(c.offsetX), c.y+float32(c.offsetY), c.radius, c.strokeWidth, green, true)
	}
	for _, v := range c.childs {
		v.Draw(screen)
	}
}

func (c *Node) AddByDegree(degree float64) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(float64(c.x), float64(c.y), float64(c.radius), degree)

	newCircle := NewDefaultCircle(float32(newX), float32(newY), 50)
	newCircle.parents = append(newCircle.parents, c)

	c.childs = append(c.childs, newCircle)

	return newCircle
}

func (c *Node) AddImgByDegree(degree float64, img *ebiten.Image) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(float64(c.x), float64(c.y), float64(c.radius), degree)

	newNode := NewDefaultImgNode(float32(newX), float32(newY), img)
	newNode.parents = append(newNode.parents, c)

	c.childs = append(c.childs, newNode)

	return newNode
}

func (c *Node) AddByRatio(a, b float32) *Node {
	ratio := a / b
	d := 360 * ratio

	return c.AddByDegree(float64(d))
}

func (c *Node) CanBeActivated() bool {
	if c.startNode {
		return true
	}
	for _, parent := range c.parents {
		if parent.active {
			return true
		}
	}
	return false
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
