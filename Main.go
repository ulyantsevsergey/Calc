package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(add(1, 2))
}

//test for add
func testSum(t *testing.T){
	total:= Sum(5,5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}


// add two numbers
func Sum(a, b int) (c int) {
	c = a + b
	return
}
