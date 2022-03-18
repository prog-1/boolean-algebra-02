package main

import "fmt"

func s(t [][]int, line int) string {
	if t[line][0] == 1 && t[line][1] == 1 {
		return "(a_0 & a_1) | "
	}
	if t[line][0] == 1 && t[line][1] == 0 {
		return "(a_0 & !a_1) | "
	}
	if t[line][0] == 0 && t[line][1] == 1 {
		return "(!a_0 & a_1) | "
	}
	return "(!a_0 & !a_1) | "
}

func DNF(t [][]int) (function string) {
	for line := 0; line < 4; line++ {
		if t[line][2] == 1 {
			function = function + s(t, line)
		}
	}
	if len(function) != 0 {
		out := []byte(function)
		function = string(out[:len(out)-3])
	}
	return function
}

func main() {
	t := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(DNF(t))

}
