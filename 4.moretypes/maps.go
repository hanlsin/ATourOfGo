package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell labs"])

	// Map literals
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{40.68443, -74.39967},
		"Google":    Vertex{37.42202, -122.08408},
		// If the top-level type is just a type name,
		// you can omit it from the elements of the literal.
		"Test": {123, 456},
	}
	fmt.Println(m2)

	v, ok := m2["Google"]
	fmt.Println("Value =", v, "Present?", ok)
	delete(m2, "Google")
	v, ok = m2["Google"]
	fmt.Println("Value =", v, "Present?", ok)

	m["YP"] = Vertex{123, 456}
	for k, v := range m {
		fmt.Printf("m[%s] = %v\n", k, v)
	}
}
