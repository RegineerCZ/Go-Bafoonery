package drawing

import (
	"fmt"
	"image/color"
	"math"

	"github.com/RegineerCZ/Go-Bafoonery/structures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	fmt.Println("Game Object package initialized")
}
var (
	redDotImage *ebiten.Image
)

func init() {
	var err error
	redDotImage, _, err = ebitenutil.NewImageFromFile("./img/test/red_dot.png")
	if err != nil {
		fmt.Println("New Image From File Error")
		fmt.Println(err)
		return
	}
}

func DrawEntity(screen *ebiten.Image, e structures.Entity) {
	ebitenutil.DrawLine(screen, e.Xcenter(), e.Ycenter(), e.Xcenter() + math.Sin(e.Rotation*math.Pi) * 10, e.Ycenter() + math.Cos(e.Rotation*math.Pi) * 10, color.White)

	img := ebiten.NewImageFromImage(redDotImage)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.RotateHue(e.Color)
	op.GeoM.Translate(e.Xpos(), e.Ypos())
	screen.DrawImage(img, op)
}

func DrawButton(screen *ebiten.Image, b structures.Button) {
	DrawBox(screen, b.Rect.X, b.Rect.Y, b.Rect.W, b.Rect.H, color.White)
	ebitenutil.DebugPrintAt(screen, b.Text, int(b.Rect.X + (b.Rect.W/2))-(6 * len(b.Text) / 2), int(b.Rect.Y + (b.Rect.H/2)) - (16 / 2))
}