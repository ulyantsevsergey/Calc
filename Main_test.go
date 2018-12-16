package main

import "testing"

//test for sum

func testSum(t *testing.T) {
	cases := []struct {
		a, b, sumExpected int
	}{
		{1, 1, 2},
		{-1, -1, 0},
		{0, 0, 0},
		{2, -1, 1},
		{-1, 2, 1},
		{-1, 2, 11}, //тест должен упасть
	}
	for _, i := range cases {
		sum := Sum(i.a, i.b)
		if sum != i.sumExpected {
			t.Errorf("Сумма %d и %d должна быть %d  посчитали %d", i.a, i.b, i.sumExpected, sum)
		}
	}
}
