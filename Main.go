package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
}


// add two numbers
func sum(a, b int) (c int) {
	c = a + b
	return
}
