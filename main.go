package main

import (
	"fmt"
	"strconv"
)

func DNF(table [][]int) (res string) {
	for _, cols := range table {
		if cols[len(cols)-1] == 1 {
			if len(res) != 0 {
				res += " | "
			}
			res += "("
			for i, v := range cols {
				if i != len(cols)-1 {
					if i != 0 {
						res += " & "
					}
					if v == 0 {
						res += "!a_" + strconv.Itoa(i)
					} else {
						res += "a_" + strconv.Itoa(i)
					}
				}
			}
			res += ")"
		}
	}

	return res
}

func main() {
	table := [][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	fmt.Println("Program that finds a Boolean function in a disjunctive normal form (DNF) for a given truth table:")
	fmt.Println("")
	fmt.Println(DNF(table))

}
