package main

import (
	"fmt"
	"github.com/hanlsin/ATourOfGo/ATourOfGo/exercise"
	"math"
	"runtime"
	"time"
)

func main() {
	fmt.Println("\n For loop")
	{
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println(sum)

		// Go's "while"
		for ; sum < 1000; {
			sum += sum
		}
		fmt.Println(sum)

		// Go's "while"
		for sum < 1000 {
			sum += sum
		}
		fmt.Println(sum)

		// an infinite loop
		/*
			for {
			}
		*/
	}

	// if
	fmt.Println("\n If")
	{
		fmt.Println(sqrt(2), sqrt(-4))
		fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	}

	// exercise: Loops and Functions
	fmt.Println("\n Exercise: Loops and Functions")
	{
		fmt.Println("Want:", math.Sqrt(2))
		fmt.Println("10 T:", exercise.SqrtT(2, 10))
		fmt.Println("11 T:", exercise.SqrtT(2, 11))
		fmt.Println("Mine:", exercise.Sqrt(2))
	}

	// switch
	fmt.Println("\n Switch")
	{
		goos := runtime.GOOS
		fmt.Println("Go OS = ", goos)

		fmt.Print("Go runs on ")

		switch os := goos; os {
		case "linux":
			fmt.Println("Linux.")
			fallthrough
		case "darwin":
			fmt.Println("OS X.")
		default:
			fmt.Printf("GOOS = %s.\n", runtime.GOOS)
		}

		// Switch evaluation order
		fmt.Println("When is Saturday?")
		today := time.Now().Weekday()
		switch time.Friday {
		case today:
			fmt.Println("Today!")
		case today + 1:
			fmt.Println("Tomorrow.")
		case today + 2:
			fmt.Println("In two days.")
		case printToday(today):
			fallthrough
		default:
			fmt.Println("Too far away.")
		}

		// Switch with no condition
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning.")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
		}
	}

	// defer
	fmt.Println("\n Defer")
	deferMain()
}

func deferMain() {
	// defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")

	// Stacking defers
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Hello")
}

func printToday(today time.Weekday) time.Weekday {
	fmt.Println("Today is", today, "!!!!!!!!!!!!!!")
	return time.Saturday
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("%g < %g\n", v, lim)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
