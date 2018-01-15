package examples

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodspointer() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

        v = Vertex{3, 4}
        Scale2(&v, 10)
        fmt.Println(Abs2(v))
}




