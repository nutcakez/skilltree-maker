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

	// radiusShows := true

	// attack 5
	// hp 7
	// cd 15
	// range 20

	// dmgImg := images[5]
	hpImg := images[7]
	cdImg := images[15]
	// rangeImg := images[20]
	//
	rank2Circle := Circle{
		x:      2000,
		y:      2000,
		radius: 600,
	}

	startNode := NewDefaultImgNode(2000, 2000, images[32])
	startNode.startNode = true
	startNode.active = true
	startNode.radius = 300
	game.display.skillTree.AddNode(startNode)

	colossusHp1 := startNode.AddByDegree(-22.5, hpImg)
	colossusCd1 := startNode.AddByDegree(22.5, cdImg)
	game.display.skillTree.AddNode(colossusHp1)
	game.display.skillTree.AddNode(colossusCd1)

	colossus := colossusHp1.AddByDegreeWithOtherCircle(0, rank2Circle, hpImg)
	colossus.AddMutualConnection(colossusCd1)
	game.display.skillTree.AddNode(colossus)

	game.display.node = startNode
	game.display.SetStartPosition(nil)

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
