package main

import (
	"fmt"
	"github.com/hanlsin/ATourOfGo/ATourOfGo/exercise"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	//a = v // error

	// Interfaces are implemented implicitly
	fmt.Println("\n Interfaces are implemented implicitly")
	{
		var i I = T{"abc"}
		fmt.Println(i.M())
	}

	// Interface values
	fmt.Println("\n Interface Values")
	{
		var i I
		var f F = 0.12
		i = &f
		fmt.Println(i.M())
		describe(i)
		i = T{"test"}
		fmt.Println(i.M())
		describe(i)
	}

	// nil underlying values
	fmt.Println("\n Nil Underlying Values")
	{
		var i I = &T2{"test2"}
		fmt.Println(i.M())
		describe(i)
		var t2 *T2
		i = t2
		fmt.Println(i.M())
		describe(i)
	}

	// nil interface values
	fmt.Println("\n Nil Interface Values")
	{
		var i I
		//fmt.Println(i.M()) // panic: SIGSEGV
		describe(i)
	}

	// the empty interface
	fmt.Println("\n The Empty Interface")
	{
		var i interface{}
		describe2(i)

		i = 42
		describe2(i)

		i = "hello"
		describe2(i)
	}

	// type assertions
	fmt.Println("\n Type Assertions")
	{
		var i interface{} = "hello"

		s := i.(string)
		fmt.Println(s)

		s, ok := i.(string)
		fmt.Println(s, ok)

		f, ok := i.(float64)
		fmt.Println(f, ok)

		//f = i.(float64)	// panic
		fmt.Println(f)
	}

	// type switches
	fmt.Println("\n Type Switches")
	{
		printType(20)
		printType("hello")
		printType(1.0)
		printType(true)
		printType(uint(10))
	}

	// Stringer
	fmt.Println("\n Stringers")
	{
		alice := Person{"alice", 19}
		bob := Person{"bob", 14}

		fmt.Println(alice)
		fmt.Println(bob.String())
	}

	// Exercise: Stringers
	fmt.Println("\n Exercise: Stringers")
	{
		ipaddr := &exercise.IPAddr{1, 2, 3, 4}
		fmt.Println(ipaddr)

		hosts := map[string]exercise.IPAddr {
			"loopback": {127, 0, 0, 1},
			"googleDNS": {8, 8, 8, 8},
		}
		for name, ip := range hosts {
			fmt.Printf("%v: %v\n", name, ip)
		}
	}
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func printType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float 64")
	default:
		fmt.Printf("Type: %T\n", v)
	}
}

func describe2(i interface{}) {
	fmt.Printf("(%T, %v)\n", i, i)
}

type T2 struct {
	S string
}

func (t2 *T2) M() string {
	if t2 == nil {
		return "<nil>"
	}
	return t2.S
}

func describe(i I) {
	fmt.Printf("(%T, %v)\n", i, i)
}

type F float64

func (f *F) M() string {
	return fmt.Sprintf("%f", *f)
}

type I interface {
	M() string
}

type T struct {
	S string
}

func (t T) M() string {
	return t.S
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
