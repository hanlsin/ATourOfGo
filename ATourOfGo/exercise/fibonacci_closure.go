package exercise

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	idx := 0
	n1 := 1
	n2 := 0

	return func() int {
		idx++
		if idx == 1 {
			return 0
		} else if idx == 2 {
			return 1
		} else {
			n := n1 + n2
			n2 = n1
			n1 = n
			return n
		}
	}
}
