package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nutcakez/skilltree-maker/display"
	"github.com/nutcakez/skilltree-maker/skilltree"
	"github.com/nutcakez/skilltree-maker/util"
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
		panner:         display.NewPanning(),
		screen2OffsetX: 50,
		screen2OffsetY: 50,
		display:        display.NewDisplay(200, 200, 500, 500),
	}

	images := util.ReadAllImageFromFolder("96x96")
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
	rank2Circle := util.Circle{
		X:      2000,
		Y:      2000,
		Radius: 600,
	}

	startNode := skilltree.NewDefaultImgNode(2000, 2000, images[32])
	startNode.StartNode = true
	startNode.Active = true
	startNode.Radius = 300
	game.display.SkillTree.AddNode(startNode)

	colossusHp1 := startNode.AddByDegree(-22.5, hpImg)
	colossusCd1 := startNode.AddByDegree(22.5, cdImg)
	game.display.SkillTree.AddNode(colossusHp1)
	game.display.SkillTree.AddNode(colossusCd1)

	colossus := colossusHp1.AddByDegreeWithOtherCircle(0, rank2Circle, hpImg)
	colossus.AddMutualConnection(colossusCd1)
	game.display.SkillTree.AddNode(colossus)

	game.display.Node = startNode
	game.display.SetStartPosition(nil)

	if err := ebiten.RunGame(&game); err != nil {
	}
}

type game struct {
	objects                        []IDrawUpdate
	panner                         *display.Panning
	screen2                        *ebiten.Image
	screen2OffsetX, screen2OffsetY int
	display                        *display.Display
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
