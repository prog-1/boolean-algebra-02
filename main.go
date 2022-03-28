package main

import (
	"fmt"
	"strconv"
)

func DNF(table [][]int) string {
	var result string
	for rows, columns := range table {
		if columns[len(columns)-1] == 1 {
			if len(result) != 0 {
				result += " & "
			}
			result += "("
			for i := 0; i < len(columns)-1; i++ {
				if table[rows][i] == 1 {
					result += "a_" + strconv.Itoa(i)
				} else {
					result += "!a_" + strconv.Itoa(i)
				}
				if i != len(columns)-2 {
					result += " | "
				}
			}
			result += ")"
		}
	}
	return result
}

func main() {
	table := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(DNF(table))
}
