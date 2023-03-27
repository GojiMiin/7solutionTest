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

	repeatLetter, repeatTimes := countRepeatLetter(input)
	fmt.Println("letter : ", repeatLetter)
	fmt.Println("mostRepeat : ", repeatTimes)

}

func countRepeatLetter(input string) (string, int) {
	currentRepeat := 1
	repeatLetter := "L"
	mostRepeat := 1
	mostRepeatLetter := ""

	for _, v := range input {
		letter := string(v)
		if letter == repeatLetter {
			currentRepeat += 1
		} else {
			if currentRepeat > 1 && currentRepeat > mostRepeat {
				mostRepeat = currentRepeat
				mostRepeatLetter = repeatLetter
			}
			currentRepeat = 1
			repeatLetter = letter
		}
	}
	return mostRepeatLetter, mostRepeat
}
