package main

import "testing"

func TestDNF(t *testing.T) {
	for _, tc := range []struct {
		table [][]int
		want  string
	}{
		{[][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
		}, "(!a_0 & !a_1) | (a_0 & !a_1) | (a_0 & a_1)"},
		{[][]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 0},
		}, "(!a_0 & a_1) | (a_0 & !a_1)"},
	} {
		got := DNF(tc.table)
		if got != tc.want {
			t.Errorf("DNF(%v) = %v, want %v", tc.table, got, tc.want)
		}
	}
}
