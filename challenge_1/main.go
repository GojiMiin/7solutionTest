package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	data := readFile()
	fmt.Println("data len : ", len(data))

	// process data
	dataLoc := 0
	summary := data[0][0]
	fmt.Println("start summary : ", summary)
	data = data[1:]
	for i := range data {
		// fmt.Println(i, v)
		fmt.Println(data[i])
		if data[i][dataLoc] > data[i][dataLoc+1] {
			summary += data[i][dataLoc]
		} else {
			summary += data[i][dataLoc+1]
			dataLoc = dataLoc + 1
		}
		fmt.Println(data[i][dataLoc])
	}
	fmt.Println("result : ", summary)
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
