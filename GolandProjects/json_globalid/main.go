package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type structure struct {
	Test []struct {
		Id int `json:"global_id"`
	}
}

func main() {
	var struc structure
	var res int
	inpFile, _ := os.Open("data.json")
	input, _ := io.ReadAll(inpFile)
	json.Unmarshal(input, &struc)
	for i := range struc.Test {
		res += struc.Test[i].Id
	}
	fmt.Println(res)
}
