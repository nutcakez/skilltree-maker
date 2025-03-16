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

	dmgImg := images[5]
	hpImg := images[7]
	cdImg := images[15]
	rangeImg := images[20]

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
	theShieldHp1 := theShield.AddByDegreeWithOtherCircle(11.25, rank2Circle, images[7])
	theShieldHp2 := theShield.AddByDegreeWithOtherCircle(11.25+22.5, rank2Circle, images[7])
	game.display.skillTree.AddNode(theShieldHp1)
	game.display.skillTree.AddNode(theShieldHp2)

	hypnoBotCd1 := hypnoBot.AddByDegreeWithOtherCircle(11.25+(22.5*2), rank2Circle, cdImg)
	theShieldHp2.AddMutualParentConnection(hypnoBotCd1)
	theShieldHp2.AddMutualChildConnection(hypnoBotCd1)
	hypnoBotCd2 := hypnoBot.AddByDegreeWithOtherCircle(11.25+(22.5*3), rank2Circle, cdImg)
	game.display.skillTree.AddNode(hypnoBotCd1)
	game.display.skillTree.AddNode(hypnoBotCd2)

	plasmaGrenadeRange1 := plasmaGrenade.AddByDegreeWithOtherCircle(11.25+(22.5*4), rank2Circle, images[20])
	plasmaGrenadeDmg2 := plasmaGrenade.AddByDegreeWithOtherCircle(11.25+(22.5*5), rank2Circle, images[5])
	plasmaGrenadeRange1.AddMutualChildConnection(hypnoBotCd2)
	plasmaGrenadeRange1.AddMutualParentConnection(hypnoBotCd2)
	game.display.skillTree.AddNode(plasmaGrenadeRange1)
	game.display.skillTree.AddNode(plasmaGrenadeDmg2)

	iceBlasterDmg1 := iceBlaster.AddByDegreeWithOtherCircle(11.25+(22.5*6), rank2Circle, dmgImg)
	iceBlasterRange1 := iceBlaster.AddByDegreeWithOtherCircle(11.25+(22.5*7), rank2Circle, rangeImg)
	game.display.skillTree.AddNode(iceBlasterDmg1)
	game.display.skillTree.AddNode(iceBlasterRange1)
	iceBlasterDmg1.AddMutualConnection(plasmaGrenadeDmg2)

	aiSoldierDmg1 := aiSoldier.AddByDegreeWithOtherCircle(11.25+(22.5*8), rank2Circle, dmgImg)
	aiSoldierHp1 := aiSoldier.AddByDegreeWithOtherCircle(11.25+(22.5*9), rank2Circle, hpImg)
	game.display.skillTree.AddNode(aiSoldierHp1)
	game.display.skillTree.AddNode(aiSoldierDmg1)
	aiSoldierDmg1.AddMutualConnection(iceBlasterRange1)

	poisonTrapDot1 := poisonTrap.AddByDegreeWithOtherCircle(11.25+(22.5*10), rank2Circle, dmgImg)
	poisonTrapDmg1 := poisonTrap.AddByDegreeWithOtherCircle(11.25+(22.5*11), rank2Circle, dmgImg)
	game.display.skillTree.AddNode(poisonTrapDmg1)
	game.display.skillTree.AddNode(poisonTrapDot1)
	poisonTrapDot1.AddMutualConnection(aiSoldierHp1)

	flameThrowerDmg1 := flameThrower.AddByDegreeWithOtherCircle(11.25+(22.5*12), rank2Circle, dmgImg)
	flameThrowerRange1 := flameThrower.AddByDegreeWithOtherCircle(11.25+(22.5*13), rank2Circle, rangeImg)
	game.display.skillTree.AddNode(flameThrowerDmg1)
	game.display.skillTree.AddNode(flameThrowerRange1)
	flameThrowerDmg1.AddMutualConnection(poisonTrapDmg1)

	// medicCd1 := medic.Add

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
