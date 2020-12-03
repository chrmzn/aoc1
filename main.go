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
		lineInt, err := strconv.ParseInt(line, 10, 0)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Line: %d\n", i)
		for i, firstVal := range previousData {
			for j, secondVal := range previousData {
				if i == j {
					continue
				}
				if firstVal+secondVal+lineInt == 2020 {
					prod := firstVal * secondVal * lineInt
					fmt.Printf("Equal to 2020: %d, %d, %d; Product: %d\n", firstVal, secondVal, lineInt, prod)
				}
			}
		}
		previousData = append(previousData, lineInt)

	}
}
