package main

import (
	"fmt"
)

func ValidateSolution(m [][]int) bool {
	sqOne := make(map[int]int)
	sqTwo := make(map[int]int)
	sqThree := make(map[int]int)
	for i := 0; i < len(m); i++ {
		checkRow := make(map[int]int)
		checkCol := make(map[int]int)
		if i == 0 || i == 3 || i == 6 {
			for key := range sqOne {
				delete(sqOne, key)
			}
			for key := range sqTwo {
				delete(sqOne, key)
			}
			for key := range sqThree {
				delete(sqOne, key)
			}
		}
		for j := 0; j < len(m); j++ {
			if m[i][j] == 0 || m[j][i] == 0 {
				return false
			}
			if _, exists := checkRow[m[i][j]]; exists {
				return false
			}
			if _, exists := checkCol[m[j][i]]; exists {
				return false
			}
			if j <= 2 {
				if _, exists := sqOne[m[i][j]]; exists {
					return false
				}
				sqOne[m[i][j]] = j
			}
			if 3 <= j && j <= 5 {
				if _, exists := sqTwo[m[i][j]]; exists {
					return false
				}
				sqTwo[m[i][j]] = j
			}
			if j >= 6 {
				if _, exists := sqThree[m[i][j]]; exists {
					return false
				}
				sqThree[m[i][j]] = j
			}
			checkRow[m[i][j]], checkCol[m[j][i]] = j, j
		}
	}
	return true
}

func main() {
	var test = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	fmt.Println(ValidateSolution(test))
}
