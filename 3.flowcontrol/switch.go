package main

import (
	"fmt"
	"runtime"
	"time"
)

func printToday(today time.Weekday) time.Weekday {
	fmt.Println("Today is", today, "!!!!!!!!!!!!!!")
	return time.Saturday
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
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
	switch time.Saturday {
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
