package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const exampleInput = "./example.input"
const actualInput = "./actual.input"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	Zero rune = '0'
	One       = '1'
)

func main() {
	file, err := os.Open(actualInput)
	check(err)

	scanner := bufio.NewScanner(file)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		bits := []rune(line)
		lines = append(lines, bits)
		check(err)
	}

	var bitsPerNumber int = len(lines[0])
	var gammaRateBinary, epsilonRateBinary []rune
	var zeroCount, oneCount int
	for bitPosition := 0; bitPosition < bitsPerNumber; bitPosition++ {
		zeroCount, oneCount = 0, 0
		for numberIndex := 0; numberIndex < len(lines); numberIndex++ {
			currentBit := lines[numberIndex][bitPosition]
			if currentBit == Zero {
				zeroCount++
			}
			if currentBit == One {
				oneCount++
			}
		}

		if zeroCount > oneCount {
			gammaRateBinary = append(gammaRateBinary, Zero)
			epsilonRateBinary = append(epsilonRateBinary, One)
		} else {
			gammaRateBinary = append(gammaRateBinary, One)
			epsilonRateBinary = append(epsilonRateBinary, Zero)
		}
	}

	gammaRate, err := strconv.ParseInt(string(gammaRateBinary), 2, 64)

	epsilonRate, err := strconv.ParseInt(string(epsilonRateBinary), 2, 64)

	fmt.Println(gammaRate, epsilonRate)
	fmt.Println(gammaRate * epsilonRate)
}
