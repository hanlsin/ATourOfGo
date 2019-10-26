package main

import (
	"fmt"
	"github.com/hanlsin/ATourOfGo/ATourOfGo/exercise"
	"golang.org/x/tour/wc"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	Lat, Long float64
}

func main() {
	// Pointer
	fmt.Println("\n Pointer")
	{
		i, j := 42, 2701
		p := &i
		fmt.Println(*p)
		*p = 21
		fmt.Println(i)

		p = &j
		*p = *p / 37
		fmt.Println(j)

		var p2 *int
		p2 = new(int)
		*p2 = 10
		fmt.Printf("Type: %T, Value: %v, *Value: %v\n", p2, p2, *p2)
	}

	// Structs
	fmt.Println("\n Structs")
	{
		fmt.Println(Vertex{1, 2})

		v := Vertex{1, 2}
		fmt.Println("v.X=", v.X)
		v.X = 4
		fmt.Println("v.X=", v.X)

		p := &v
		p.X = 1e3
		p.Y = 4
		fmt.Println(v)
		fmt.Println(p)
		fmt.Println("p.X=", p.X)
		fmt.Println("p.Y=", p.Y)
		fmt.Println(*p)
		fmt.Println("*p.X=", (*p).X)
		fmt.Println("*p.Y=", (*p).Y)

		v1 := Vertex{1, 2}
		v2 := Vertex{X: 1}
		v3 := Vertex{}
		p1 := &Vertex{3, 4}
		fmt.Println(v1, v2, v3, p1, *p1)
	}

	// Array
	fmt.Println("\n Array")
	{
		var a [2]string
		a[0] = "hello"
		a[1] = "world"
		fmt.Println(a[0], a[1])
		fmt.Println(a)

		primes := [6]int{2, 3, 5, 7, 11, 13}
		fmt.Println(primes)

		var p1 []int = primes[1:4]
		fmt.Println(p1)
	}

	// Slices
	fmt.Println("\n Slices")
	{
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
		fmt.Println("\nSlice Literals")
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
		fmt.Println("\nSlice Defaults")
		fmt.Println(st[0:len(st)])
		fmt.Println(st[:len(st)])
		fmt.Println(st[0:])
		fmt.Println(st[:])
		fmt.Println(st)

		// Slice length and capacity
		fmt.Println("\nSlice Capacity")
		s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		printSlice(&s)

		// Slice the slice to give it zero length.
		s = s[:0]
		printSlice(&s)

		// Extend its length
		s = s[:4]
		printSlice(&s)

		// Drop its firtst two values
		s = s[2:]
		printSlice(&s)

		// Nil slices
		fmt.Println("\nNil Slice")
		var nilS []int
		printSlice(&nilS)
		if nilS == nil {
			fmt.Println("Nil!")
		}
		fmt.Println(nilS)

		// Creating a slice with make
		fmt.Println("\nSlice with make")
		aa := make([]int, 5)
		fmt.Print("aa = ")
		printSlice(&aa)
		bb := make([]int, 0, 5)
		fmt.Print("bb = ")
		printSlice(&bb)
		bb = bb[:cap(bb)]
		fmt.Print("bb = ")
		printSlice(&bb)
		cc := bb[:2]
		fmt.Print("cc = ")
		printSlice(&cc)
		dd := cc[2:5]
		fmt.Print("dd = ")
		printSlice(&dd)
		ee := cc[2:]
		fmt.Print("ee = ")
		printSlice(&ee)
		ff := make([]int, 5, 5)
		fmt.Print("ff = ")
		printSlice(&ff)
		gg := []int {0, 0, 0, 0, 0}
		fmt.Print("gg = ")
		printSlice(&gg)

		// Slices can contain any type, including other slices.
		fmt.Println("\nSlices of slices")
		// Create a tic-tac-toe board.
		board := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}

		// players take turns
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
		fmt.Println("\nAppend to Slice")
		var ss []int
		printSlice(&ss)
		ss = append(ss, 0)
		printSlice(&ss)
		ss = append(ss, 1, 2, 3, 4)
		printSlice(&ss)
		ss = append(ss, ss...)
		printSlice(&ss)
		for i := 0; i < len(ss); i++ {
			fmt.Printf("[%d] %v ", i, ss[i])
		}
		fmt.Println()

		// Range
		fmt.Println("\nRange")
		var pow = []int{2, 4, 6, 8, 10}
		for i, v := range pow {
			fmt.Printf("2**%d = %v\n", i, v)
		}

		pow = make([]int, 10)
		printSlice(&pow)
		for i := range pow {
			pow[i] = 1 << uint(i) // pow[i] = pow[i-1]*2
		}
		printSlice(&pow)
		// You can skip the index or value by assigning to _.
		for _, v := range pow {
			fmt.Printf("%d\n", v)
		}
	}

	// Exercise: Slices
	fmt.Println("\n Exercise: Slices")
	{
		yRange := 10
		xRange := 10
		randSrc := rand.NewSource(time.Now().UnixNano())
		randGen := rand.New(randSrc)
		pic := make([][]uint8, yRange)
		for i := range pic {
			pic[i] = make([]uint8, xRange)
			for j := range pic[i] {
				pic[i][j] = uint8(randGen.Intn(100))
				fmt.Printf("[%d][%d] %v\n", i, j, pic[i][j])
			}
		}
		fmt.Printf("Type: %T, Value: %v\n", pic, pic)
		fmt.Println(*exercise.Pic(2, 2, &pic))
	}

	// Maps
	fmt.Println("\n Maps")
	{
		var m map[string]Vertex2
		if m == nil {
			fmt.Println("Nil!!")
		}
		fmt.Printf("Type: %T, Value: %v\n", m, m)

		m = make(map[string]Vertex2)
		if m != nil {
			fmt.Println("No Nil!!")
		}
		fmt.Printf("Type: %T, Value: %v\n", m, m)

		m["Bell labs"] = Vertex2{
			40.68433, -74.39967,
		}
		fmt.Printf("Type: %T, Value: %v\n", m, m)
		fmt.Println(m["Bell labs"])

		// Map literals
		var m2 = map[string]Vertex2{
			"Bell Labs": {40.68443, -74.39967},
			"Google":    {37.42202, -122.08408},
			// If the top-level type is just a type name,
			// you can omit it from the elements of the literal.
			"Test": Vertex2{123, 456},
		}
		fmt.Printf("Type: %T, Value: %v\n", m2, m2)

		v, ok := m2["Google"]
		fmt.Println("Value =", v, "Present?", ok)
		delete(m2, "Google")

		v, ok = m2["Google"]
		fmt.Println("Value =", v, "Present?", ok)

		m["YP"] = Vertex2{123, 456}
		for k, v := range m {
			fmt.Printf("m[%s] = %v\n", k, v)
		}
	}

	// Exercise: Maps
	fmt.Println("\n Exercise: Maps")
	{
		fmt.Println(exercise.WordCount("a a a aaa, bbb, dd, ca,ajlk  asd "))
		wc.Test(exercise.WordCount)
	}

	// Function Values
	fmt.Println("\n Function Values")
	{
		add := func(x, y float64) float64 {
			return x + y
		}
		fmt.Println(add(1, 1))
		fmt.Println(compute(add))

		hypot := func(x, y float64) float64 {
			return math.Sqrt(x*x + y*y)
		}
		fmt.Println(hypot(5, 12))
		fmt.Println(compute(hypot))

		fmt.Println(compute(math.Pow))
	}

	// Function Closures
	fmt.Println("\n Function Closures")
	{
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(pos(i), neg(-2*i))
		}
	}

	// Exercise: Fibonacci closure
	fmt.Println("\n Exercise: Fibonacci closure")
	{
		f := exercise.Fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}
	}
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func printSlice(s *[]int) {
	fmt.Printf("len=%d cap=%d %v\n", len(*s), cap(*s), *s)
}
