package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	/* naked return */
	// NOTE: Naked return statements should be used only in short functions,
	// as with the example shown here.
	// They can harm readability in longer functions.
	return
}

func main() {
	fmt.Println(split(17))
}
