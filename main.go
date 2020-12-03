package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("File contents: %s", contents)

	data := string(contents)
	rowData := strings.Split(data, "\n")

	var previousData []int64

	for _, line := range rowData {
		i, err := strconv.ParseInt(line, 10, 0)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Line: %d\n", i)
		for _, val := range previousData {
			if val+i == 2020 {
				prod := val * i
				fmt.Printf("Equal to 2020: %d, %d; Product: %d\n", val, i, prod)
			}
		}
		previousData = append(previousData, i)

	}
}
