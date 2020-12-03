package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File contents: %s", data)

}
