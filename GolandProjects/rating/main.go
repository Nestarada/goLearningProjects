package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type student struct {
	Rating []int
}

type group struct {
	Students []student
}

type avg struct {
	Average float32
}

func main() {
	var res avg
	var count int
	var n1 group
	read, _ := os.Open("js.txt")
	text, _ := io.ReadAll(read)
	json.Unmarshal(text, &n1)
	for ; count < len(n1.Students); count++ {
		res.Average += float32(len(n1.Students[count].Rating))
	}
	res.Average /= float32(count)
	end, _ := json.MarshalIndent(res, "", "    ")
	fmt.Print(string(end))
}
