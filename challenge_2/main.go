package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoding: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input[:len(input)-2])
	fmt.Println("input : ", len(input))

	repeatLetter, repeatTimes, repeatLoc := countRepeatLetter(input)
	fmt.Println("letter : ", repeatLetter)
	fmt.Println("mostRepeat : ", repeatTimes)
	fmt.Println("loc : ", repeatLoc)
}

func countRepeatLetter(input string) (string, int, int) {
	currentRepeat := 1
	repeatLetter := "L"
	startRepeatLoc := 0
	mostRepeat := 1
	mostRepeatLetter := ""
	startMostRepeat := 0

	for i, v := range input {
		letter := string(v)
		if letter == repeatLetter {
			currentRepeat += 1
		} else {
			if currentRepeat > 1 && currentRepeat > mostRepeat {
				mostRepeat = currentRepeat
				mostRepeatLetter = repeatLetter
				startMostRepeat = startRepeatLoc
			}
			currentRepeat = 1
			repeatLetter = letter
			startRepeatLoc = i
		}
	}

	return mostRepeatLetter, mostRepeat, startMostRepeat
}
