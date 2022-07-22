package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func insert(size, elemSliceRange int) {
	var sorted []int
	for i := 0; i < size; i++ {
		sorted = append(sorted, rand.Intn(elemSliceRange) - rand.Intn(elemSliceRange))
	}
	tStart := time.Now()
	for i := 1; i < len(sorted); i++ {
		j := i
		for j > 0 {
			if sorted[j-1] > sorted[j] {
				sorted[j-1], sorted[j] = sorted[j], sorted[j-1]
			}
			j--
		}
	}
	tFinish := time.Now()
	//fmt.Println(sorted)
	fmt.Println("Insert duration: ", tFinish.Sub(tStart))
	wg.Done()
}

func insertV2(size, elemSliceRange int) {
	var unsorted []int
	for i := 0; i < size; i++ {
		unsorted = append(unsorted, rand.Intn(elemSliceRange) - rand.Intn(elemSliceRange))
	}
	tStart := time.Now()
	var sorted []int
	var index int
	for range unsorted {
		min := elemSliceRange
		for i, v := range unsorted {
			if v < min {
				index = i
				min = v
			}
		}
		sorted = append(sorted, min)
		unsorted = append(unsorted[:index], unsorted[index+1:]...)
	}
	tFinish := time.Now()
	//fmt.Println(sorted)
	fmt.Println("InsertV2 duration: ", tFinish.Sub(tStart))
	wg.Done()
}

func bubble(size, elemSliceRange int) {
	var sorted []int
	for i := 0; i < size; i++ {
		sorted = append(sorted, rand.Intn(elemSliceRange) - rand.Intn(elemSliceRange))
	}
	tStart := time.Now()
	for i := 0; i < len(sorted) ; i++ {
		for j := 0; j < len(sorted) - i - 1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	tFinish := time.Now()
	//fmt.Println(sorted)
	fmt.Println("Bubble duration: ", tFinish.Sub(tStart))
	wg.Done()
}

func shaker(size, elemSliceRange int) {
	var sorted []int
	for i := 0; i < size; i++ {
		sorted = append(sorted, rand.Intn(elemSliceRange) - rand.Intn(elemSliceRange))
	}
	tStart := time.Now()
	var begin, end = 0, len(sorted) - 1
	for begin < end {
		if sorted[begin] > sorted[begin+1] {
			sorted[begin], sorted[begin+1] = sorted[begin+1], sorted[begin]
		}
		if sorted[end] < sorted[end-1] {
			sorted[end], sorted[end-1] = sorted[end-1], sorted[end]
		}
		begin++
		end--
	}
	tFinish := time.Now()
	//fmt.Println(sorted)
	fmt.Println("Shaker duration: ", tFinish.Sub(tStart))
	wg.Done()
}

func selection(size, elemSliceRange int) {
	var sorted []int
	for i := 0; i < size; i++ {
		sorted = append(sorted, rand.Intn(elemSliceRange) - rand.Intn(elemSliceRange))
	}
	tStart := time.Now()
	for i := 0; i < len(sorted); i++ {
		min := i
		for j := i; j < len(sorted); j++ {
			if sorted[j] < sorted[i] {min = j}
		}
		sorted[i], sorted[min] = sorted[min], sorted[i]
	}
	tFinish := time.Now()
	//fmt.Println(sorted)
	fmt.Println("Select duration: ", tFinish.Sub(tStart))
	wg.Done()
}

func comb(size, elemSliceRange int) {
	var sorted []int
	for i := 0; i < size; i++ {
		sorted = append(sorted, rand.Intn(elemSliceRange) - rand.Intn(elemSliceRange))
	}
	tStart := time.Now()
	step := len(sorted) - 1
	for step > 0 {
		for j := 0; j + step <= len(sorted) - 1; j++ {
			if sorted[j] > sorted[j+step] {
				sorted[j], sorted[j+step] = sorted[j+step], sorted[j]
			}
		}
		step = int(float32(step)/1.247)
	}
	tFinish := time.Now()
	//fmt.Println(sorted)
	fmt.Println("Comb duration: ", tFinish.Sub(tStart))
	wg.Done()
}

func main() {
	rand.Seed(time.Now().Unix())
	var elemSliceRange = 100_000
	var size = 100_000
	wg.Add(6)
	go selection(size, elemSliceRange)
	go bubble(size, elemSliceRange)
	go shaker(size, elemSliceRange)
	go comb(size, elemSliceRange)
	go insert(size, elemSliceRange)
	go insertV2(size, elemSliceRange)
	wg.Wait()
}
