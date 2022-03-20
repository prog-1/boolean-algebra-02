package main

import (
	"fmt"
	"strconv"
)

func DNF(table [][]int) (res string) {
	for _, v := range table {
		if v[len(v)-1] == 1 {
			res += "("
			for i, j := range v {
				if i != len(v)-1 {
					if i != 0 {
						res += " | "
					}
					if j == 0 {
						res += "!a_" + strconv.Itoa(i)
					} else {
						res += "a_" + strconv.Itoa(i)
					}
				}
			}
			res += ")"
			if res != "" {
				res += " & "
			}
		}
	}
	return
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
