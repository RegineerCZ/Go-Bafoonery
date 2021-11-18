package main

import (
	"fmt"
	"image/color"
	_ "image/png"

	"github.com/RegineerCZ/Go-Bafoonery/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1280
	screenHeight = 960
	playAreaWidth = 640
	playAreaHeight = 480

	playAreaLft = (screenWidth /2) - (playAreaWidth /2)
	playAreaTop = 10
)

var (
	redDotImage *ebiten.Image
	entities = [10]structures.Entity {}
)

func init() {
	var err error
	redDotImage, _, err = ebitenutil.NewImageFromFile("img/test/red_dot.png")
	if err != nil {
		fmt.Println("New Image From File Error")
		fmt.Println(err)
		return
	}

	for i :=0; i< 10; i++{
		entities[i] = structures.NewEntity(playAreaWidth, playAreaHeight)
	}

	fmt.Println("Done")
}

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/*w, h := gophersImage.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.
	op.GeoM.Rotate(float64(g.count%360) * 2 * math.Pi / 360)

	// Move the image to the screen's center.
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	screen.DrawImage(gophersImage, op)*/

	
	ebitenutil.DrawLine(screen, playAreaLft, playAreaTop, playAreaLft+playAreaWidth, playAreaTop, color.White)
	ebitenutil.DrawLine(screen, playAreaLft, playAreaTop, playAreaLft, playAreaTop+playAreaHeight, color.White)
	ebitenutil.DrawLine(screen, playAreaLft+playAreaWidth, playAreaTop, playAreaLft+playAreaWidth, playAreaTop+playAreaHeight, color.White)
	ebitenutil.DrawLine(screen, playAreaLft, playAreaTop+playAreaHeight, playAreaLft+playAreaWidth, playAreaTop+playAreaHeight, color.White)

	for i :=0; i< 10; i++{
		drawEntity(screen, entities[i].Color, entities[i].X, entities[i].Y)
	}

	//ebitenutil.DebugPrintAt(screen, "Test Value: "+fmt.Sprint(theta), playAreaLft + 10, playAreaTop + playAreaHeight + 10)
}

func drawEntity(screen *ebiten.Image, colorShift, x, y float64) {
	//fmt.Printf("Drawing entity: [%g, %g] - col: %g\n", x, y, colorShift)
	img := ebiten.NewImageFromImage(redDotImage)
	op := &ebiten.DrawImageOptions{}
	//tps := ebiten.MaxTPS()
	//theta := 2.0 * math.Pi * float64((g.count/10)%tps) / float64(tps)
	//op.ColorM.RotateHue(theta)
	op.ColorM.RotateHue(colorShift)
	op.GeoM.Translate(playAreaLft+ x, playAreaTop+ y)
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Bafoonery")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}