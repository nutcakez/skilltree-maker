package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	red    = color.RGBA{255, 0, 0, 255}
	green  = color.RGBA{0, 255, 0, 1}
	orange = color.RGBA{255, 127, 80, 1}
)

func main() {
	ebiten.SetWindowSize(1280, 960)

	f, err := os.Create("cpuprofile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	game := game{
		objects:        make([]IDrawUpdate, 0),
		panner:         NewPanning(),
		screen2OffsetX: 50,
		screen2OffsetY: 50,
		display:        NewDisplay(200, 200, 500, 500),
	}

	images := ReadAllImageFromFolder("96x96")
	fmt.Println(images)

	radiusShows := true

	otherMainCircle := NewDefaultImgNode(300, 700, images[2])
	otherMainCircle.showBigRadius = radiusShows
	otherMainCircle.startNode = true
	otherMainCircle.onActivate = func() {
		fmt.Println("yes, this is not the default")
	}
	c1 := otherMainCircle.AddByDegree(45, images[3])
	c2 := c1.AddByDegree(45, images[4])

	c2.AddByDegree(45, images[5])
	c2.AddByDegree(90, images[7])
	c2.showBigRadius = radiusShows
	c3 := c1.AddByDegree(135, images[5])
	c3.AddByDegree(45, images[6])
	c4 := c3.AddByDegree(90, images[8])

	c5 := c4.AddByDegree(50, images[11])
	c4.AddByDegree(100, images[12])
	c4.AddByDegree(150, images[13])

	c5.AddByDegree(60, images[14])
	c5.AddByDegree(90, images[15])
	c5.AddByDegree(120, images[16])

	game.display.node = otherMainCircle
	game.display.skillTree.AddNode(otherMainCircle)

	if err := ebiten.RunGame(&game); err != nil {
	}
}

type game struct {
	objects                        []IDrawUpdate
	panner                         *Panning
	screen2                        *ebiten.Image
	screen2OffsetX, screen2OffsetY int
	display                        *Display
}

func (g *game) Update() error {
	g.display.Update()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf(`FPS: %0.2f, TPS: %0.2f`, ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
	g.display.Draw(screen)
}

func (g *game) Layout(a, b int) (int, int) { return 1280, 960 }

type IDrawUpdate interface {
	Draw(*ebiten.Image)
	Update(int, int, int, int, float64)
}
