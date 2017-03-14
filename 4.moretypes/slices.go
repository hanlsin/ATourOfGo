package main

import "fmt"
import "strings"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	/*
		For the array

			var a [10]int

		these slice expressions are equivalent:

			a[0:10]
			a[:10]
			a[0:]
			a[:]
	*/
	fmt.Println(st[0:len(st)])
	fmt.Println(st[:len(st)])
	fmt.Println(st[0:])
	fmt.Println(st[:])

	// Slice length and capacity
	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its firtst two values
	s = s[2:]
	printSlice(s)

	// Nil slices
	var nilS []int
	printSlice(nilS)
	if nilS == nil {
		fmt.Println("Nil!")
	}

	// Creating a slice with make
	aa := make([]int, 5)
	printSlice(aa)
	bb := make([]int, 0, 5)
	printSlice(bb)
	bb = bb[:cap(bb)]
	printSlice(bb)
	cc := bb[:2]
	printSlice(cc)
	dd := cc[2:5]
	printSlice(dd)
	ee := cc[2:]
	printSlice(ee)

	// Slices can contain any type, including other slices.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println(strings.Join(strings.Fields("Hello World"), "$"))

	// Appending to a slice
	var ss []int
	printSlice(ss)
	ss = append(ss, 0)
	printSlice(ss)
	ss = append(ss, 1, 2, 3, 4)
	printSlice(ss)

	// Range
	var pow = []int{2, 4, 6, 8, 10}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)
	printSlice(pow)
	for i := range pow {
		pow[i] = 1 << uint(i) // pow[i] = pow[i-1]*2
	}
	printSlice(pow)
	// You can skip the index or value by assigning to _.
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
