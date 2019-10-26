package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var c, python, java bool
var i, j int = 1, 2
var x, y string = "abc", "def"

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	pd *int
)

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	fmt.Println("\n Variable with Init")
	var a, b, c = true, false, "no!"
	fmt.Println(i, j, a, b, c)

	fmt.Println(x, y)

	// Short variable declarations
	fmt.Println("\n Short Variable Declaration")
	k := 3
	x, y := true, "false"
	fmt.Println(k, x, y)

	// basic types
	fmt.Println("\n Basic Types")
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", pd, pd)
	pd = new(int)
	*pd = 13
	fmt.Printf("Type: %T Value: %v Int: %v\n", pd, pd, *pd)

	// zero variables
	fmt.Println("\n Zero Variables")
	{
		var i int
		var f float64
		var b bool
		var s string
		fmt.Printf("%v %v %v %q\n", i, f, b, s)
	}

	// type conversion
	fmt.Println("\n Type Conversion")
	{
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y))
		fmt.Println(x, y, f)
		var z uint = uint(f)
		fmt.Println(x, y, f, z)
		fmt.Printf("Type: %T, Value: %v\n", x, x)
		fmt.Printf("Type: %T, Value: %v\n", y, y)
		fmt.Printf("Type: %T, Value: %v\n", f, f)
		fmt.Printf("Type: %T, Value: %v\n", z, z)
	}

	// type inference
	fmt.Println("\n Type Inference")
	{
		v := 42           // int
		f := 3.142        // float64
		g := 0.867 + 0.5i // complex128
		fmt.Printf("v is of type %T\n", v)
		fmt.Printf("f is of type %T\n", f)
		fmt.Printf("g is of type %T\n", g)
		fmt.Println(v, f, g)
	}

	// Constants
	fmt.Println("\n Constants")
	{
		const World = "월드"
		fmt.Println("Hello", World)
		fmt.Println("Happy", Pi, "Day")

		const Truth = true
		fmt.Println("Go rules?", Truth)
	}

	// Numeric Constants
	fmt.Println("\n Numeric Constants")
	{
		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
	}
}

func init() {
	fmt.Println("init")

	x = "XXX"
	y = "YYY"

	fmt.Println(x, y)
}