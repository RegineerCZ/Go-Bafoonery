package structures

import "fmt"

func init() {
	fmt.Println("Vector package initialized")
}

type Vector struct {
	X,Y float64
}

func NewVector(x,y float64) Vector {
	v := Vector{}

	v.X = x
	v.Y = y

	return v
}