package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

func main() {
	// struct type
	fmt.Println("\n Struct Type")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// non-struct type
	fmt.Println("\n Non-struct Type")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer Receivers
	fmt.Println("\n Pointer Receivers")
	{
		v := Vertex{3, 4}
		v.Scale(10)
		fmt.Println(v.Abs())
	}

	// Pointers and functions
	fmt.Println("\n Pointers and Functions")
	{
		v := Vertex{3, 4}
		Scale(&v, 10)
		//Scale(v, 10) // compile error!
		fmt.Println(Abs(v))
	}
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}
