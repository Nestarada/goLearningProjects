package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("task.data.txt")
	reader := bufio.NewReader(file)
	var count int
	for {
		count++
		buff, _ := reader.ReadString(';')
		if buff == "0;" {
			fmt.Print(count)
			break
		}
	}
}
