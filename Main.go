package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(1, 2))
}


// add two numbers
func Sum(a, b int) (c int) {
	c = a + b
	return
}
