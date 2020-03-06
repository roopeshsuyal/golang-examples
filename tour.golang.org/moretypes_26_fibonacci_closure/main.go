package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current, prev := 0, 0
	return func() int {
		old_prev := prev
		if prev == 0 {
			if current == 1 {
				prev = 1
			} else {
				current = 1
			}
			return prev
		}
		prev = current
		current = current + old_prev
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
