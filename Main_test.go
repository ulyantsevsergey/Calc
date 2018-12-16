package main

import "testing"

//test for add
func testSum(t *testing.T){
	total:= sum(5,5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
