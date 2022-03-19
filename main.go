package main

import (
	"fmt"
	"math"
)

func plug(t [][]int, line, coll int) (function string) {
	function = function + "("
	for a := 0; a < coll-1; a++ {
		if t[line][a] == 1 {
			function = function + "a_" + fmt.Sprint(a) + " & "
		} else {
			function = function + "!a_" + fmt.Sprint(a) + " & "
		}
	}
	if len(function) != 0 {
		out := []byte(function)
		function = string(out[:len(out)-3])
	}
	function = function + ") | "
	return function
}

func DNF(t [][]int) (function string) {
	a := math.Sqrt(float64(len(t))) + 1
	coll := int(a)

	for line := 0; line < len(t); line++ {
		if t[line][coll-1] == 1 {
			function = function + plug(t, line, coll)
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
