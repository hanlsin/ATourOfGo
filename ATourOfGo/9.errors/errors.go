package main

import (
	"fmt"
	"github.com/hanlsin/ATourOfGo/ATourOfGo/exercise"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		When: time.Now(),
		What: "it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Exercise: Errors
	fmt.Println("\n Exercise: Errors")
	{
		fmt.Println(exercise.SqrtWithE(4))
		fmt.Println(exercise.SqrtWithE(-4))
	}
}
