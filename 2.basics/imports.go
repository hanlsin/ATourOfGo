package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have (float)%g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have (int)%d problems.\n", int(math.Sqrt(7)))
}
