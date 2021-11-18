package structures

import "fmt"

func init() {
	fmt.Println("Rect package initialized")
}

type Rect struct {
	X,Y,W,H float64
}

func NewRect(x,y,w,h float64) Rect {
	v := Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}

	return v
}