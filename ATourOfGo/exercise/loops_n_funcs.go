package exercise

import (
	"fmt"
	"math"
)

/**
Exercise: Loops and Functions
As a simple way to play with functions and loops, implement the square root function using Newton's method.

In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:


Z(n+1) = Z(n) - (Z(n)^2 - x) / (2*Z(n))


To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

z := float64(1)
z := 1.0
*/

func SqrtT(x float64, t int) float64 {
	z := 1.0

	for ; t > 0; t-- {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func Sqrt(x float64) float64 {
	z, p := float64(1), 0.
	delta := float64(1)

	const DELTA_MIN = 1e-12
	fmt.Println("Delata minimum =", DELTA_MIN)
	for delta > DELTA_MIN {
		p = z
		z = z - (z*z-x)/(2*z)
		delta = math.Abs(z - p)
		fmt.Println(delta, ": Z= ", z)
	}
	fmt.Println("p=", p, "z=", z, "delta=", delta)
	return z
}
