package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(1, 2))
}

// add two numbers
func Sum(a, b int) int {
	return a + b
}

// sub two numbers
func Sub(a, b int) int {
	return a - b
}
