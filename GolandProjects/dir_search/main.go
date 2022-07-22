package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)
func walkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text()[0] != 00 {
			fmt.Println(scanner.Text())
		}
	}
	file.Close()
	return nil
}

func main() {
	filepath.Walk(".\\task", walkFunc)
}