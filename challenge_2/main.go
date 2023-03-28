package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input = ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input[:len(input)-2])

	decodeArr, initErr := initDecodeArr(string(input[0]))
	numberToAdd := 0
	result := ""

	if initErr != nil {
		log.Fatal(initErr)
		os.Exit(1)
	}

	for i, v := range input[1:] {
		decodeArr = compareEncode(string(v), i, decodeArr, &numberToAdd)
		if decodeArr == nil {
			log.Fatal("wrong letter in input")
			os.Exit(1)
		}
	}

	for i := 0; i < len(decodeArr); i++ {
		decodeArr[i] = decodeArr[i] + numberToAdd
		result += strconv.Itoa(decodeArr[i])
	}

	fmt.Println("output = ", result)

}

func initDecodeArr(firstLetter string) ([]int, error) {
	switch firstLetter {
	case "L":
		return []int{1, 0}, nil
	case "R":
		return []int{0, 1}, nil
	case "=":
		return []int{0, 0}, nil
	default:
		return []int{}, fmt.Errorf("wrong letter set")
	}
}

func compareEncode(curLetter string, curLoc int, decoded []int, numberToAdd *int) []int {
	curDecoded := decoded[curLoc+1]
	switch curLetter {
	case "L":
		nextDecoded := curDecoded - 1
		if nextDecoded < 0 && nextDecoded < *numberToAdd {
			*numberToAdd = -nextDecoded
		}
		newDecoded := append(decoded, nextDecoded)
		return newDecoded
	case "R":
		nextDecoded := curDecoded + 1
		newDecoded := append(decoded, nextDecoded)
		return newDecoded
	case "=":
		newDecoded := append(decoded, curDecoded)
		return newDecoded
	default:
		return nil
	}
}
