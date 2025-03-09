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
		screen2:        ebiten.NewImage(2000, 2000),
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
	c1 := otherMainCircle.AddImgByDegree(45, images[3])
	c2 := c1.AddImgByDegree(45, images[4])

	c2.AddImgByDegree(45, images[5])
	c2.AddImgByDegree(90, images[7])
	c2.showBigRadius = radiusShows
	c3 := c1.AddImgByDegree(135, images[5])
	c3.AddImgByDegree(45, images[6])
	c4 := c3.AddImgByDegree(90, images[8])

	c5 := c4.AddImgByDegree(50, images[11])
	c4.AddImgByDegree(100, images[12])
	c4.AddImgByDegree(150, images[13])

	c5.AddImgByDegree(60, images[14])
	c5.AddImgByDegree(90, images[15])
	c5.AddImgByDegree(120, images[16])

	game.display.node = otherMainCircle
	// game.objects = append(game.objects, otherMainCircle)

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
	// g.panner.Update()
	// for i := range g.objects {
	// 	g.objects[i].Update(g.panner.offsetX, g.panner.offsetY, g.screen2OffsetX, g.screen2OffsetY, g.panner.zoom)
	// }
	//
	// g.screen2.Clear()
	g.display.Update()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf(`FPS: %0.2f, TPS: %0.2f`, ebiten.ActualFPS(), ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
	// for i := range g.objects {
	// 	g.objects[i].Draw(g.screen2)
	// }
	//
	// // we want to show 500x500 picture
	//
	// op := ebiten.DrawImageOptions{}
	// // if we zoom in it means we want to see a bigger picture from screen2 but scaled to 500x500
	//
	// // figure out the scale here
	// op.GeoM.Scale(1/g.panner.zoom, 1/g.panner.zoom)
	// op.GeoM.Translate(float64(g.screen2OffsetX), float64(g.screen2OffsetY))
	//
	// // which part of the stuff we want?
	// rect := image.Rect(g.panner.offsetX, g.panner.offsetY, int(500*g.panner.zoom), int(500*g.panner.zoom))
	// img := ebiten.NewImageFromImage(g.screen2.SubImage(rect))
	//
	// vector.StrokeRect(screen, 0+float32(g.screen2OffsetX), 0+float32(g.screen2OffsetY), 500, 500, 2, color.RGBA{255, 0, 0, 255}, true)
	//
	// screen.DrawImage(img, &op)
	g.display.Draw(screen)
}

func (g *game) Layout(a, b int) (int, int) { return 1280, 960 }

type IDrawUpdate interface {
	Draw(*ebiten.Image)
	Update(int, int, int, int, float64)
}
