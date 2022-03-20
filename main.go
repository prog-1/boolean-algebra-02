package main

import (
	"fmt"
	"strings"
)

func DNF(table [][]int) (res string) {
	for rows, cols := range table {
		if cols[len(cols)-1] == 1 {
			if len(res) != 0 {
				res += " | "
			}
			res += "("
			for i := 0; i < len(cols)-1; i++ {
				if table[rows][i] == 1 {
					res += "a_" + fmt.Sprint(i) + " & "
				} else {
					res += "!a_" + fmt.Sprint(i) + " & "
				}
			}
			res = strings.TrimRight(res, "& ")
			res += ")"
		}
	}
	return res
}

func main() {
	t := [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 1},
		{0, 1, 0, 0},
		{0, 1, 1, 1},
		{1, 0, 0, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}
	fmt.Println(DNF(t))

}
