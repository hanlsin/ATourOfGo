package main

import (
	"fmt"
	"github.com/hanlsin/ATourOfGo/ATourOfGo/exercise"
	"golang.org/x/tour/reader"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Readers!!")

	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Printf("b[:n] = %U\n", b[:n])

		if err == io.EOF {
			break
		}
	}

	// Exercise: Readers
	fmt.Println("\n Exercise: Readers")
	{
		stream := exercise.NewMyReader("abcsdrlkjasdkhopiquewtrakdjshfkaj HAekljasldkfhAlices")
		b := make([]byte, 8)
		for {
			n, err := stream.Read(b)
			fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
			fmt.Printf("b[:n] = %q\n", b[:n])

			if err == io.EOF {
				break
			}
		}

		reader.Validate(stream)
		reader.Validate(&exercise.MyReader{})
	}

	// Exercise: rot13Reader
	fmt.Println("\n Exercise: rot13Reader")
	{
		s := strings.NewReader("Lbh penpxrq gur pbqr!")
		r := exercise.NewReader(s)
		io.Copy(os.Stdout, r)
	}
}
