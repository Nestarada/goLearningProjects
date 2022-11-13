package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var width int
	fmt.Scanln(&width)
	var n, k int
	fmt.Scanln(&n, &k)
	var photos []string
	var heights []int
	var buff string
	for i := 0; i < n; i++ {
		fmt.Scanln(&buff)
		photos = strings.Split(buff, "x")
		intWidth, _ := strconv.Atoi(photos[0])
		intHeight, _ := strconv.Atoi(photos[1])
		newHeight := float64(width*intHeight) / float64(intWidth)
		if newHeight-float64(int(newHeight)) > 0 {
			newHeight = float64(int(newHeight + 1))
		}
		heights = append(heights, int(newHeight))
	}
	sort.Ints(heights)
	var min, max int
	for i, j := 0, len(heights)-1; i < k; i, j = i+1, j-1 {
		min += heights[i]
		max += heights[j]
	}
	fmt.Println(min)
	fmt.Println(max)
}
