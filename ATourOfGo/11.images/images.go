package main

import (
	"fmt"
	"github.com/hanlsin/ATourOfGo/ATourOfGo/exercise"
	"golang.org/x/tour/pic"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	// Exercise: Images
	fmt.Println("\n Exercise: Images")
	{
		m := exercise.MyImage{}
		pic.ShowImage(m)
	}
}
