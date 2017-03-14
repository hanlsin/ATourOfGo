package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	var a, b, c = true, false, "no!"
	fmt.Println(i, j, a, b, c)

	// Short variable declarations
	k := 3
	x, y := true, "false"
	fmt.Println(k, x, y)
}
