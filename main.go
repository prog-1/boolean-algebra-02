package main

import "strconv"

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
						result += "!a_" + strconv.Itoa(i)
					} else {
						result += "a_" + strconv.Itoa(i)
					}
				}
			}
			result += ")"
		}
	}
	return
}
