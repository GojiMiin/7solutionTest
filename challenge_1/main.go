package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	data := readFile()
	fmt.Println("data : ", data[0])
	fmt.Println("data : ", data[1])
	fmt.Println("data : ", data[2])
	fmt.Println("data len : ", len(data))
}

func readFile() [][]int {
	var arr [][]int
	jsonFile, openErr := ioutil.ReadFile("files/hard.json")
	if openErr != nil {
		fmt.Println(openErr)
	}

	unmarshalErr := json.Unmarshal(jsonFile, &arr)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
	}

	return arr
}
