package skilltree

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nutcakez/skilltree-maker/util"
)

var (
	red    = color.RGBA{255, 0, 0, 255}
	green  = color.RGBA{0, 255, 0, 1}
	orange = color.RGBA{255, 127, 80, 1}
)

type Node struct {
	showBigRadius, Active, StartNode bool
	X, Y, Radius, StrokeWidth        float32
	offsetX, offsetY                 int
	childs                           []*Node
	parents                          []*Node
	img                              *ebiten.Image
	OnActivate                       func()
	Requirement                      func() bool
	Tags                             map[string]string
}

func NewDefaultImgNode(x, y float32, img *ebiten.Image) *Node {
	return &Node{
		showBigRadius: false,
		Active:        false, X: x,
		Y:           y,
		Radius:      300,
		StrokeWidth: 2,
		childs:      make([]*Node, 0),
		img:         img,
		OnActivate:  func() { fmt.Println("actived node at", x, y) },
		Requirement: func() bool { return true },
		Tags:        nil,
	}
}

func (c *Node) Update(offsetX, offsetY, windowOffsetX, windowOffsetY int, zoom float64) {
	c.offsetX = offsetX
	c.offsetY = offsetY
	c.checkClick(zoom, windowOffsetX, windowOffsetY)
	// if !c.checkClick(zoom, windowOffsetX, windowOffsetY) {
	// 	for i := range c.childs {
	// 		c.childs[i].Update(offsetX, offsetY, windowOffsetX, windowOffsetY, zoom)
	// 	}
	// }
}

func (c *Node) checkClick(zoom float64, windowOffsetX, windowOffsetY int) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()
		w := c.img.Bounds().Dx()
		h := c.img.Bounds().Dy()
		if util.PointInRect(x, y,
			int(float64(c.X-float32(w/2)+float32(c.offsetX))/zoom+float64(windowOffsetX)),
			int(float64(c.Y-float32(h/2)+float32(c.offsetY))/zoom+float64(windowOffsetY)),
			int(float64(w)/zoom),
			int(float64(h)/zoom)) && c.CanBeActivated() {
			c.Active = !c.Active
			c.OnActivate()
			return true
		}
	}

	return false
}

func (c *Node) Draw(screen *ebiten.Image) {
	var color color.RGBA

	if c.Active {
		color = green
	} else {
		color = red
	}

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)), float64(c.Y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)))
	screen.DrawImage(c.img, &op)
	strokeWith := 5
	vector.StrokeRect(screen,
		float32(c.X-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)),
		float32(c.Y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)),
		float32(c.img.Bounds().Dx()),
		float32(c.img.Bounds().Dy()),
		float32(strokeWith),
		color,
		true,
	)

	if c.showBigRadius {
		vector.StrokeCircle(screen, c.X+float32(c.offsetX), c.Y+float32(c.offsetY), c.Radius, c.StrokeWidth, green, true)
	}
	// for _, v := range c.childs {
	// 	v.Draw(screen)
	// }
}

func (n *Node) DrawLines(screen *ebiten.Image) {
	for _, v := range n.childs {
		// with antialias false its a bit better when zoomed out, should be set to false only when its zoomed out?
		vector.StrokeLine(screen, n.X+float32(n.offsetX), n.Y+float32(n.offsetY), v.X+float32(n.offsetX), v.Y+float32(n.offsetY), 2, orange, false)
	}
}

func (n *Node) AddByDegreeWithOtherCircle(degree float64, circle util.Circle, img *ebiten.Image) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(circle.X, circle.Y, circle.Radius, degree)

	newNode := NewDefaultImgNode(float32(newX), float32(newY), img)
	newNode.parents = append(newNode.parents, n)

	n.childs = append(n.childs, newNode)

	return newNode
}

func (c *Node) AddByDegree(degree float64, img *ebiten.Image) *Node {
	degree -= 90
	newX, newY := GetPointOnCircle(float64(c.X), float64(c.Y), float64(c.Radius), degree)

	newNode := NewDefaultImgNode(float32(newX), float32(newY), img)
	newNode.parents = append(newNode.parents, c)

	c.childs = append(c.childs, newNode)

	return newNode
}

func (c *Node) AddByRatio(a, b float32, img *ebiten.Image) *Node {
	ratio := a / b
	d := 360 * ratio

	return c.AddByDegree(float64(d), img)
}

func (n *Node) AddTag(key, value string) {
	if n.Tags == nil {
		n.Tags = map[string]string{
			key: value,
		}
	} else {
		n.Tags[key] = value
	}
}

// probably this is just lazy work from my side
func (n *Node) AddMutualConnection(node *Node) {
	n.AddMutualParentConnection(node)
	n.AddMutualChildConnection(node)
}

func (n *Node) AddMutualParentConnection(node *Node) {
	n.parents = append(n.parents, node)
	node.parents = append(node.parents, n)
}

func (n *Node) AddMutualChildConnection(node *Node) {
	n.childs = append(n.childs, node)
	node.childs = append(node.childs, n)
}

func (c *Node) CanBeActivated() bool {
	if c.Active {
		return false
	}

	var hasActiveParent bool
	for _, parent := range c.parents {
		if parent.Active {
			hasActiveParent = true
			break
		}
	}

	return c.Requirement()
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
