package main

import "testing"

func TestProposition(t *testing.T) {
	for _, tc := range []struct {
		testNumber int
		table      [][]int
		want       string
	}{
		{1, [][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
		}, "(!a_0 | !a_1) & (a_0 | !a_1) & (a_0 | a_1)"},
		{2, [][]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 0},
		}, "(!a_0 | a_1) & (a_0 | !a_1)"},
		{3, [][]int{
			{0, 0, 0, 1},
			{0, 0, 1, 0},
			{0, 1, 0, 1},
			{0, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 1, 0},
			{1, 1, 0, 0},
			{1, 1, 1, 1},
		}, "(!a_0 | !a_1 | !a_2) & (!a_0 | a_1 | !a_2) & (!a_0 | a_1 | a_2) & (a_0 | a_1 | a_2)"},
	} {
		got := DNF(tc.table)
		if got != tc.want {
			t.Errorf("ERR: DNF(#%v): got = %v, want = %v", tc.testNumber, got, tc.want)
		}
	}
}
