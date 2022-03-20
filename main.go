package main

import (
	"fmt"
)

func intToString(i int) (s string) {
	for ; i > 0; i /= 10 {
		s = fmt.Sprint(string(byte(i%10+'0')), s)
	}
	if s == "" {
		return "0"
	}
	return s

}

func DNF(table [][]int) (result string) {
	for _, v := range table {
		if v[len(v)-1] == 1 {
			if result != "" {
				result += " | "
			}
			result += "("
			for i, v2 := range v {
				if i != len(v)-1 {
					if i != 0 {
						result += " & "
					}
					if v2 == 0 {
						result += "!a_" + intToString(i)
					} else {
						result += "a_" + intToString(i)
					}
				}
			}
			result += ")"
		}
	}
	return
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
