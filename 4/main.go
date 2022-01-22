package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const exampleInput = "./example.input"
const actualInput = "./actual.input"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Grid = [][]BingoNumber

type BingoNumber struct {
	value int 
	isMarked bool
}

type Bingo struct {
	numbers []int
	grids   []Grid
}

func parseDrawnNumbers(s string) []int {
	var result []int

	for _, numStr := range strings.Split(s, ",") {
		value, err := strconv.Atoi(numStr)
		check(err)
		result = append(result, value)
	}
	return result
}

func parseGridLine(s string) []int {
	var result []int
	for _, numStr := range strings.Fields(s) {
		number, err := strconv.Atoi(numStr)
		check(err)
		result = append(result, number)
	}
	return result
}

func parseGrid(gridLines []string) Grid {
	var grid [][]int = make([][]int, 5)
	println(len(gridLines))
	for i := 0; i < 5; i++ {
		grid[i] = parseGridLine(gridLines[i])
	}

	var result Grid = make([][]BingoNumber, 5)
	for i := 0; i < 5; i++ {
		resultLine := make([]BingoNumber, 5)
		for j := 0; j < 5; j++ {
			
			resultLine[j] = BingoNumber{grid[i][j], false}
		}
		result[i] = resultLine
	} 
	return result
}

func markGridForGivenNumber(grid Grid, number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j].value == number {
				grid[i][j].isMarked = true
			}
		}
	} 
}



func checkAllRows(grid Grid) bool {
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		for columnIndex := 0; columnIndex < 5; columnIndex++ {
			if grid[rowIndex][columnIndex].isMarked == false {
				break // false found in row => next row
			}
			return true // all row is marked
		}
	} 
	return false // no fully marked row found
}

func checkColumns(grid Grid) bool {
	for columnIndex := 0; columnIndex < 5; columnIndex++ {
		for rowIndex := 0; rowIndex < 5; rowIndex++ {
			if grid[columnIndex][rowIndex].isMarked == false {
				break // false found in column
			}
			return true // all column is marked
		}
	} 
	return false // no fully marked column found
}

func sumAllUnmarkedNumbers(grid Grid) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j].isMarked == false {
				sum += grid[i][j].value
			}
		}
	} 
	return sum
}

func main() {
	file, err := os.Open(exampleInput)
	check(err)

	scanner := bufio.NewScanner(file)
	var game Bingo

	scanner.Scan() // reading first line
	game.numbers = parseDrawnNumbers(scanner.Text())

	var gridLineBuffer []string

	scanner.Scan() // skipping first blank line after numbers

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			gridLineBuffer = append(gridLineBuffer, line)
		} else {

			grid := parseGrid(gridLineBuffer)
			game.grids = append(game.grids, grid)
			gridLineBuffer = make([]string, 0)
		}
	}
	grid := parseGrid(gridLineBuffer)
	game.grids = append(game.grids, grid) // last grid whatever

	fmt.Println(game.numbers)


	for i := 0; i < len(game.numbers); i++ {
		
		drawnNumber := game.numbers[i]
		fmt.Println("$$$$$$$$$$$$$$$$ NEW NUMBER $$$$$$$$$$$$$$$$", drawnNumber)
		for gridIndex := 0; gridIndex < len(game.grids); gridIndex++ {
			currentGrid := game.grids[gridIndex]
			markGridForGivenNumber(currentGrid, drawnNumber)
			fmt.Println("=================== GRID ========================", gridIndex)
			fmt.Println(currentGrid)
			rowCheck := checkAllRows(currentGrid)
			columnCheck := checkColumns(currentGrid)
			println(rowCheck, columnCheck)
			// if (rowCheck || columnCheck) {
			// 	println("winning grid ", gridIndex) 
			// 	unmarkedSum := sumAllUnmarkedNumbers(currentGrid)
			// 	println("unmarked numbers sum ", unmarkedSum)
			// 	println("winning number ", drawnNumber)
			// 	println(unmarkedSum * drawnNumber)
			// }
		}
	}
}
