package main

import (
	"testing"
)

//test for addition
func TestSum(t *testing.T) {
	cases := []struct {
		a, b, sumExpected int
	}{
		{1, 1, 2},
		{-1, -1, -2},
		{0, 0, 0},
		{2, -1, 1},
		{-1, 2, 1},
		{-1, 2, 1},
	}
	for _, i := range cases {
		sum := Sum(i.a, i.b)
		if sum != i.sumExpected {
			t.Errorf("Сумма %d и %d должна быть %d  посчитали %d", i.a, i.b, i.sumExpected, sum)
		}
	}
}

//test for subtraction
func TestSub(t *testing.T) {
	cases := []struct {
		a, b, subExpected int
	}{
		{1, 1, 0},
		{-1, -1, 0},
		{0, 0, 0},
		{2, -1, 3},
		{-1, 2, -3},
		{2, 1, 1},
	}
	for _, i := range cases {
		sub := Sub(i.a, i.b)
		if sub != i.subExpected {
			t.Errorf("Разница %d и %d должна быть %d  посчитали %d", i.a, i.b, i.subExpected, sub)
		}
	}
}
