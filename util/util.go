package util

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Circle struct {
	X, Y, Radius float64
}

func PointInCircle(cx, cy, cr float32, x, y int32) bool {
	dx := float32(x) - cx
	dy := float32(y) - cy
	return dx*dx+dy*dy <= cr*cr
}

func PointInRect(x, y, rx, ry, width, height int) bool {
	return x >= rx && x <= rx+width && y >= ry && y <= ry+height
}

func ReadAllImageFromFolder(path string) []*ebiten.Image {
	images := make([]*ebiten.Image, 0)

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return images
	}

	for i, fileName := range files {
		fmt.Println(i, fileName.Name())
		img, _, err := ebitenutil.NewImageFromFile(path + "/" + fileName.Name())
		if err != nil {
			fmt.Println(err)
		}
		images = append(images, img)
	}

	return images
}
