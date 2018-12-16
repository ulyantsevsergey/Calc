package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
}

// add two numbers 1 + 2
func add(a, b int) (c int) {
	c = a + b + b*2
	return c
}
