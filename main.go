package main

import (
	"fmt"
)

func DNF(table [][]int) string {
	var s string
	for _, column := range table {
		if column[2] == 1 {
			if len(s) > 0 {
				s = s + " | "
			}
			s = s + "("
			if column[0] == 0 {
				s = s + "!a_0"
			} else {
				s = s + "a_0"
			}
			s = s + " & "
			if column[1] == 0 {
				s = s + "!a_1"
			} else {
				s = s + "a_1"
			}
			s = s + ")"
		}
	}
	return s
}

func CNF(table [][]int) string {
	var s string
	var i1, i2, i3, i4 [4]int // (!a_0 | !a_1) & (!a_0 | a_1) & (a_0 | !a_1) & (a_0 | a_1)
	for _, column := range table {
		if column[2] == 1 {
			a, b := column[0], column[1]
			if a == 0 && b == 0 {
				i1[0], i1[1], i1[2] = 1, 1, 1
			}
			if a == 0 && b == 1 {
				i2[0], i2[1], i2[3] = 1, 1, 1
			}
			if a == 1 && b == 0 {
				i3[0], i3[2], i3[3] = 1, 1, 1
			}
			if a == 1 && b == 1 {
				i4[1], i4[2], i4[3] = 1, 1, 1
			}
		} else {
			a, b := column[0], column[1]
			if a == 0 && b == 0 {
				i1[0], i1[1], i1[2], i1[3] = 1, 1, 1, 1
			}
			if a == 0 && b == 1 {
				i2[0], i2[1], i2[2], i2[3] = 1, 1, 1, 1
			}
			if a == 1 && b == 0 {
				i3[0], i3[1], i3[2], i3[3] = 1, 1, 1, 1
			}
			if a == 1 && b == 1 {
				i4[0], i4[1], i4[2], i4[3] = 1, 1, 1, 1
			}
		}
	}
	StringExists := false
	if i1[0] == 1 && i2[0] == 1 && i3[0] == 1 && i4[0] == 1 {
		s = s + "(!a_0 | !a_1)"
		StringExists = true
	}
	if i1[1] == 1 && i2[1] == 1 && i3[1] == 1 && i4[1] == 1 {
		if StringExists {
			s = s + " & "
		}
		s = s + "(!a_0 | a_1)"
		StringExists = true
	}
	if i1[2] == 1 && i2[2] == 1 && i3[2] == 1 && i4[2] == 1 {
		if StringExists {
			s = s + " & "
		}
		s = s + "(a_0 | !a_1)"
		StringExists = true
	}
	if i1[3] == 1 && i2[3] == 1 && i3[3] == 1 && i4[3] == 1 {
		if StringExists {
			s = s + " & "
		}
		s = s + "(a_0 | a_1)"
	}
	return s
}

func main() {
	table := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(DNF(table))
	fmt.Println(CNF(table))
}
