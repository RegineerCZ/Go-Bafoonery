package drawing

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	fmt.Println("Basic Shapes package initialized")
}

func DrawBox(screen *ebiten.Image, x,y,w,h float64, c color.Color) {
	ebitenutil.DrawLine(screen, x, y, x+w, y, c)
	ebitenutil.DrawLine(screen, x, y, x, y+h, c)
	ebitenutil.DrawLine(screen, x+w, y, x+w, y+h, c)
	ebitenutil.DrawLine(screen, x, y+h, x+w, y+h, c)
}
