package entity

import (
	"fmt"
	"math/rand"
)

const (
	
)

func init() {
	fmt.Println("Entity package initialized")
}

type Entity struct {
	x, y float64
	turningSpeed float64
	movementSpeed float64
	color float64
}

func (e Entity) NewEntity(areaWidth, areaHeight int) Entity {
	colors := [...]float64 {0.0, 0.8, 1.2, 2.2, 3.0, 3.6, 4.4, 5.0, 5.6}

	e.x = float64(rand.Intn(areaWidth))
	e.y = float64(rand.Intn(areaHeight))
	e.turningSpeed = 0.01
	e.movementSpeed = 0.1
	e.color = colors[rand.Intn(len(colors))]
	return e
}