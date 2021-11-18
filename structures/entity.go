package structures

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Entity package initialized")
}

type Entity struct {
	X, Y float64
	TurningSpeed float64
	MovementSpeed float64
	Color float64
}

func NewEntity(areaWidth, areaHeight int) Entity {
	// Represented by hue shift of original image
	colors := [...]float64 {0.0, 0.8, 1.2, 2.2, 3.0, 3.6, 4.4, 5.0, 5.6}
	e := Entity{}

	e.X = float64(rand.Intn(areaWidth))
	e.Y = float64(rand.Intn(areaHeight))
	e.TurningSpeed = 0.01
	e.MovementSpeed = 0.1
	e.Color = colors[rand.Intn(len(colors))]
	return e
}