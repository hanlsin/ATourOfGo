package main

import "fmt"
import "time"
import "math/rand"

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10))

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("New random number is ", r1.Intn(10))

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Second New random number is ", rand.Intn(10))
}
