package main

import (
	"reflect"
	"testing"
)

func TestDNF(t *testing.T) {
	for _, tc := range []struct {
		name  string
		table [][]int
		want  string
	}{
		{"1 (example)", [][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
		}, "(!a_0 & !a_1) | (a_0 & !a_1) | (a_0 & a_1)"},

		{"2", [][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 0},
		}, "(!a_0 & !a_1)"},

		{"3", [][]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 0},
			{1, 1, 0},
		}, "(!a_0 & a_1)"},

		{"4", [][]int{
			{0, 0, 1},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, "(!a_0 & !a_1) | (!a_0 & a_1) | (a_0 & !a_1) | (a_0 & a_1)"},
	} {
		if got := DNF(tc.table); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("DNF(%v) = %v, want = %v", tc.name, got, tc.want)
		}
	}
}

func TestCNF(t *testing.T) {
	for _, tc := range []struct {
		name  string
		table [][]int
		want  string
	}{
		{"1", [][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
		}, "(a_0 | !a_1)"},

		{"2", [][]int{
			{0, 0, 1},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}, ""},

		{"3", [][]int{
			{0, 0, 1},
			{0, 1, 1},
			{1, 0, 0},
			{1, 1, 0},
		}, "(!a_0 | !a_1) & (!a_0 | a_1)"},

		{"4", [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 1},
		}, "(a_0 | !a_1) & (a_0 | a_1)"},

		{"5", [][]int{
			{0, 0, 1},
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 0},
		}, "(!a_0 | !a_1) & (!a_0 | a_1) & (a_0 | !a_1)"},

		{"6", [][]int{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 0},
			{1, 1, 0},
		}, "(!a_0 | !a_1) & (!a_0 | a_1) & (a_0 | a_1)"},

		{"7", [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 0, 1},
			{1, 1, 0},
		}, "(!a_0 | !a_1) & (a_0 | !a_1) & (a_0 | a_1)"},

		{"8", [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 1},
		}, "(!a_0 | a_1) & (a_0 | !a_1) & (a_0 | a_1)"},
	} {
		if got := CNF(tc.table); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("CNF(%v) = %v, want = %v", tc.name, got, tc.want)
		}
	}
}
