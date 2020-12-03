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
	fmt.Printf("File contents: %s", contents)

	data := string(contents)
	rowData := strings.Split(data, "\n")

	// var previousData []int

	for _, line := range rowData {
		fmt.Println(line)
		i, err := strconv.ParseInt(line, 10, 0)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Line: %d\n", i)

	}
}
