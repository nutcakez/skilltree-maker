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
	gray   = color.RGBA{125, 125, 125, 200}
)

type Node struct {
	showBigRadius, Active, Locked, StartNode bool
	X, Y, Radius, StrokeWidth                float32
	HoverText                                string
	offsetX, offsetY                         int
	childs                                   []*Node
	parents                                  []*Node
	img                                      *ebiten.Image
	OnActivate                               func()
	Requirement                              func() bool
	CustomCanBeActivated                     func() bool // Overwrites the whole CanBeActivated flow, giving full controll to the user. Can be used to create unusual things, like not requiring nodes to have an active parents to be activated.
	Tags                                     map[string]string
	RuleEngine                               *NodeRuleEngine
}

func NewDefaultImgNode(x, y float32, img *ebiten.Image) *Node {
	return &Node{
		showBigRadius: false,
		Active:        false,
		X:             x,
		Y:             y,
		Radius:        300,
		HoverText: `default hoverText to test multiline
potato 2`,
		StrokeWidth: 2,
		childs:      make([]*Node, 0),
		img:         img,
		OnActivate:  func() { fmt.Println("actived node at", x, y) },
		Requirement: func() bool { return true },
		Tags:        nil,
	}
}

func (c *Node) Update(offsetX, offsetY, windowOffsetX, windowOffsetY int, zoom float64) (clicked, hovered bool) {
	c.offsetX = offsetX
	c.offsetY = offsetY
	clicked, hovered = c.checkClick(zoom, windowOffsetX, windowOffsetY)
	return
}

func (c *Node) checkClick(zoom float64, windowOffsetX, windowOffsetY int) (clicked, hovered bool) {
	x, y := ebiten.CursorPosition()
	w := c.img.Bounds().Dx()
	h := c.img.Bounds().Dy()
	hovered = util.PointInRect(x, y,
		int(float64(c.X-float32(w/2)+float32(c.offsetX))/zoom+float64(windowOffsetX)),
		int(float64(c.Y-float32(h/2)+float32(c.offsetY))/zoom+float64(windowOffsetY)),
		int(float64(w)/zoom),
		int(float64(h)/zoom))
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && hovered {
		if c.CanBeActivated() {
			c.Active = !c.Active
			c.OnActivate()
			clicked = true
		}
	}

	return clicked, hovered
}

func (c *Node) Draw(screen *ebiten.Image) {
	var color color.RGBA

	if c.Locked {
		color = gray
	} else {
		if c.Active {
			color = green
		} else {
			color = red
		}
	}

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)), float64(c.Y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)))
	screen.DrawImage(c.img, &op)
	strokeWith := 5

	if color == gray {
		vector.DrawFilledRect(screen,
			float32(c.X-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)),
			float32(c.Y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)),
			float32(c.img.Bounds().Dx()),
			float32(c.img.Bounds().Dy()),
			color,
			true,
		)
	} else {
		vector.StrokeRect(screen,
			float32(c.X-float32(c.img.Bounds().Dx()/2)+float32(c.offsetX)),
			float32(c.Y-float32(c.img.Bounds().Dy()/2)+float32(c.offsetY)),
			float32(c.img.Bounds().Dx()),
			float32(c.img.Bounds().Dy()),
			float32(strokeWith),
			color,
			true,
		)
	}

	if c.showBigRadius {
		vector.StrokeCircle(screen, c.X+float32(c.offsetX), c.Y+float32(c.offsetY), c.Radius, c.StrokeWidth, green, true)
	}
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
	if c.CustomCanBeActivated != nil {
		return c.CustomCanBeActivated()
	}

	if c.Active || c.Locked {
		return false
	}

	var hasActiveParent bool
	for _, parent := range c.parents {
		if parent.Active {
			hasActiveParent = true
			break
		}
	}
	if !hasActiveParent {
		return false
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
