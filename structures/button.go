package structures

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	fmt.Println("Button package initialized")
}

type Button struct {
	Rect Rect
	Text string
	Func func()
}

func NewButton(text string, r Rect, f func()) Button {
	b := Button{
		Text: text,
		Rect: r,
		Func: f,
	}

	return b
}

func (b *Button) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		v := Vector{X: float64(x), Y: float64(y)}

		if v.X < b.Rect.X { return }
		if v.X > (b.Rect.X + b.Rect.W) { return }
		if v.Y < b.Rect.Y { return }
		if v.Y > (b.Rect.Y + b.Rect.H) { return }

		b.Func()
	}
}