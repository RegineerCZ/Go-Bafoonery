package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"time"

	"github.com/RegineerCZ/Go-Bafoonery/drawing"
	"github.com/RegineerCZ/Go-Bafoonery/structures"
	"github.com/hajimehoshi/ebiten/v2"
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
	entities = [100]structures.Entity {}
	newTime, oldTime int64 = 0, 0
	b = structures.NewButton("+", structures.NewRect(100, playAreaHeight+50, 30, 30), func() {fmt.Println("Press !")})
)

func init() {
	for i :=0; i< 100; i++{
		entities[i] = structures.NewEntity(i, playAreaLft, playAreaTop, playAreaWidth-10, playAreaHeight-10)
		entities[i].Move()
	}
}

type Game struct {
	count int
}

func (g *Game) Update() error {
	newTime    = time.Now().UnixNano()
    deltaTime := float64(((newTime - oldTime) / 1000000)) * 0.001
    oldTime    = newTime

	for i :=0; i< 100; i++{
		entities[i].Update(deltaTime)
	}

	b.Update()

	g.count += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawing.DrawBox(screen, playAreaLft-5, playAreaTop-5, playAreaWidth+10, playAreaHeight+10, color.White)
	drawing.DrawButton(screen, b)
	
	for i :=0; i< 100; i++{
		drawing.DrawEntity(screen, entities[i])
	}

	//ebitenutil.DebugPrintAt(screen, "Test Value: "+fmt.Sprint(theta), playAreaLft + 10, playAreaTop + playAreaHeight + 10)
}



func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	oldTime = time.Now().UnixNano()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Bafoonery")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}