package structures

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Entity package initialized")
}

type Entity struct {
	ID int
	GroupID int
	Pos Rect
	Rotation float64
	Rotating float64
	TurningSpeed float64
	MovementSpeed float64
	Color float64

	FollowingID int

	Area Rect
}

func NewEntity(i, areaX, areaY, areaW, areaH int) Entity {
	// Represented by hue shift of original image
	colors := [...]float64 {0.0, 0.8, 1.2, 2.2, 3.0, 3.6, 4.4, 5.0, 5.6}
	e := Entity{
		ID: i,
		GroupID: rand.Intn(len(colors)),
		Pos: Rect{
			X: float64(rand.Intn(areaW)), 
			Y: float64(rand.Intn(areaH)), 
			W: 10, 
			H: 10},
		Rotation: 0.0,
		Rotating: 0.0,
		TurningSpeed: 0.5,
		MovementSpeed: 30,

		FollowingID: -1,

		Area: Rect {
			X: float64(areaX),
			Y: float64(areaY),
			W: float64(areaW),
			H: float64(areaH),
		},
	}
	e.Color = colors[e.GroupID]
	return e
}

func (e Entity) Xpos() float64{
	return e.Pos.X + float64(e.Area.X)
}

func (e Entity) Ypos() float64{
	return e.Pos.Y + float64(e.Area.Y)
}

func (e Entity) Xcenter() float64{
	return (e.Pos.X + float64(e.Area.X)) + (e.Pos.W / 2)
}

func (e Entity) Ycenter() float64{
	return (e.Pos.Y + float64(e.Area.Y)) + (e.Pos.H / 2)
}

func (e *Entity) Move() {
	switch (e.FollowingID){
	case -1:
		// Move randomly until another entity is in front of you
		e.Rotating = (rand.Float64() + 0.1) * e.TurningSpeed * float64(rand.Intn(3)-1)
		time.AfterFunc(time.Duration(rand.Intn(1500)+500)*time.Millisecond, e.Move)
	}
}

func (e *Entity) Update(delta float64) {
	e.Rotation += e.Rotating * delta
	e.Pos.X += math.Sin(e.Rotation * math.Pi) * e.MovementSpeed * delta
	if e.Pos.X > float64(e.Area.W){
		e.Pos.X = float64(e.Area.W)
	}
	if e.Pos.X < 0{
		e.Pos.X = 0
	}
	e.Pos.Y += math.Cos(e.Rotation * math.Pi) * e.MovementSpeed * delta
	if e.Pos.Y > float64(e.Area.H){
		e.Pos.Y = float64(e.Area.H)
	}
	if e.Pos.Y < 0{
		e.Pos.Y = 0
	}
}