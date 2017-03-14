package main

import "fmt"

func main() {
	// defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")

	// Stacking defers
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Hello")
}
