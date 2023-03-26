package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, readErr := readFile()
	if readErr != nil {
		log.Fatal(readErr)
	}

	// process data
	dataLoc := 0
	summary := data[0][0]
	fmt.Println("start process")
	data = data[1:]
	for i := range data {
		if data[i][dataLoc] < data[i][dataLoc+1] {
			dataLoc += 1
		}
		summary += data[i][dataLoc]
	}
	fmt.Println("result : ", summary)
}

func readFile() ([][]int, error) {
	var arr [][]int
	jsonFile, openErr := ioutil.ReadFile("files/hard.json")
	if openErr != nil {
		return nil, openErr
	}

	unmarshalErr := json.Unmarshal(jsonFile, &arr)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return arr, nil
}
