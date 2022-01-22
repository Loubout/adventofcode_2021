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

func main() {
	file, err := os.Open(actualInput)
	check(err)

	scanner := bufio.NewScanner(file)
    var lines []int
    for scanner.Scan() {
        val, err := strconv.Atoi(scanner.Text())
        lines = append(lines, val)
        check(err)
    }


    var count, currrentVal, previousVal int
    count = 0
    for i := 1; i < len(lines); i++ {
        fmt.Println(currrentVal, previousVal)
        currrentVal = lines[i]
        previousVal = lines[i - 1]

        if currrentVal > previousVal {
            count++
        }
    }
    fmt.Println(count)
}
