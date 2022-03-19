package main

import "testing"

func TestDNF(t *testing.T) {
	for _, tc := range []struct {
		name  string
		table [][]int
		want  string
	}{
		{"",
			[][]int{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
				{1, 1, 1},
			}, "(!a_0 & !a_1) | (a_0 & !a_1) | (a_0 & a_1)"},
		{"Xor",
			[][]int{
				{0, 0, 0},
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 0},
			}, "(!a_0 & a_1) | (a_0 & !a_1)"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := DNF(tc.table); got != tc.want {
				t.Errorf("DNF(%v) = %v, want %v", tc.table, got, tc.want)
			}
		})
	}
}

func BenchmarkDNF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DNF([][]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 0}})
	}
}
