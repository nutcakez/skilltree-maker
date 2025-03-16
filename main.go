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

	startNode := NewDefaultImgNode(2000, 2000, images[32])
	startNode.startNode = true
	startNode.radius = 300
	game.display.skillTree.AddNode(startNode)

	// ran1 nodes
	theShield := startNode.AddByDegree(22.5, images[1])
	game.display.skillTree.AddNode(theShield)
	hypnoBot := startNode.AddByDegree(22.5+45, images[2])
	game.display.skillTree.AddNode(hypnoBot)
	plasmaGrenade := startNode.AddByDegree(22.5+90, images[3])
	game.display.skillTree.AddNode(plasmaGrenade)
	iceBlaster := startNode.AddByDegree(22.5+135, images[4])
	game.display.skillTree.AddNode(iceBlaster)
	aiSoldier := startNode.AddByDegree(22.5+180, images[5])
	game.display.skillTree.AddNode(aiSoldier)
	poisonTrap := startNode.AddByDegree(22.5+225, images[6])
	game.display.skillTree.AddNode(poisonTrap)
	flameThrower := startNode.AddByDegree(22.5+270, images[7])
	game.display.skillTree.AddNode(flameThrower)
	medic := startNode.AddByDegree(22.5+315, images[8])
	game.display.skillTree.AddNode(medic)
	// ran1 nodes end

	// rank2 nodes
	theShieldHp1 := theShield.AddByDegree(11.25, images[7])
	theShieldHp2 := theShield.AddByDegree(11.25+22.5, images[7])
	game.display.skillTree.AddNode(theShieldHp1)
	game.display.skillTree.AddNode(theShieldHp2)

	hypnoBotCd1 := hypnoBot.AddByDegree(11.25+(22.5*2), images[15])
	hypnoBotCd2 := hypnoBot.AddByDegree(11.25+(22.5*3), images[15])
	game.display.skillTree.AddNode(hypnoBotCd1)
	game.display.skillTree.AddNode(hypnoBotCd2)

	plasmaGrenadeRange1 := plasmaGrenade.AddByDegree(11.25+(22.5*4), images[20])
	plasmaGrenadeDmg2 := plasmaGrenade.AddByDegree(11.25+(22.5*5), images[5])
	game.display.skillTree.AddNode(plasmaGrenadeRange1)
	game.display.skillTree.AddNode(plasmaGrenadeDmg2)

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
