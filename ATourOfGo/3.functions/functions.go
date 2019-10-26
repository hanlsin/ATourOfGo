package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func addPtr(x, y *int) int {
	return *x + *y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	/* Go's return values may be named. If so, they are treated as variables defined at the top of the function. */
	x = sum * 4 / 9
	y = sum - x

	/* naked return */
	// NOTE: Naked return statements should be used only in short functions,
	// as with the example shown here.
	// They can harm readability in longer functions.
	return
}

func main() {
	fmt.Println(add(42, 13))

	x, y := 22, 11
	fmt.Println(addPtr(&x, &y))

	// multiple results
	fmt.Println("\n Multiple Results")
	a, b := "hello", "world"
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)

	// named results
	fmt.Println("\n Named Results")
	fmt.Println(split(17))

	// imports
	fmt.Println("\n Imports")
	fmt.Printf("Now you have (float)%g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have (int)%d problems.\n", int(math.Sqrt(7)))

	// Packages
	fmt.Println("\n Packages")
	{
		fmt.Println("My favorite number is ", rand.Intn(10))

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		fmt.Println("New random number is ", r1.Intn(10))

		rand.Seed(time.Now().UnixNano())
		fmt.Println("Second New random number is ", rand.Intn(10))
	}
}
