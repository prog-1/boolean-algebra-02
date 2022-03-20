package main

import "testing"

func TestDNF(t *testing.T) {
	for _, tc := range []struct {
		input [][]int
		want  string
	}{
		{[][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
		}, "(!a_0 & !a_1) | (a_0 & !a_1) | (a_0 & a_1)"},

		{[][]int{
			{0, 0, 0, 0},
			{0, 0, 1, 1},
			{0, 1, 0, 0},
			{0, 1, 1, 1},
			{1, 0, 0, 1},
			{1, 0, 1, 0},
			{1, 1, 0, 1},
			{1, 1, 1, 0},
		}, "(!a_0 & !a_1 & a_2) | (!a_0 & a_1 & a_2) | (a_0 & !a_1 & !a_2) | (a_0 & a_1 & !a_2)"},
	} {
		got := DNF(tc.input)
		if got != tc.want {
			t.Errorf("DNF(%v), got = %v, want = %v", tc.input, got, tc.want)
		}
	}
}
