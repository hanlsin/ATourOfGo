package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("This code is error when you build: ", math.pi)
	fmt.Println("This code is error when you build: ", math.Pi)
	fmt.Println("Because an exported name should start with a capital letter.")
	fmt.Println("In this case 'pi' in 'math' package, ")
}
