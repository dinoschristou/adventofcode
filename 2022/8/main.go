package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parse(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			height, _ := strconv.Atoi(string(c))
			row = append(row, height)
		}
		grid = append(grid, row)
	}
	return grid
}

func first(grid [][]int) {
	// Map of visible trees.
	visible := make([][]bool, len(grid))
	for i := range visible {
		visible[i] = make([]bool, len(grid[i]))
		visible[i][0] = true
		visible[i][len(grid[i])-1] = true
	}
	for j := range visible[0] {
		visible[0][j] = true
		visible[len(visible)-1][j] = true
	}

	// Rows
	for i := 0; i < len(grid); i++ {
		tallestSoFar := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > tallestSoFar {
				visible[i][j] = true
				tallestSoFar = grid[i][j]
			}
		}
		tallestSoFar = 0
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] > tallestSoFar {
				visible[i][j] = true
				tallestSoFar = grid[i][j]
			}
		}
	}

	// Columns
	for j := 0; j < len(grid[0]); j++ {
		tallestSoFar := 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > tallestSoFar {
				visible[i][j] = true
				tallestSoFar = grid[i][j]
			}
		}
		tallestSoFar = 0
		for i := len(grid) - 1; i >= 0; i-- {
			if grid[i][j] > tallestSoFar {
				visible[i][j] = true
				tallestSoFar = grid[i][j]
			}
		}
	}

	count := 0
	for i := 0; i < len(visible); i++ {
		for j := 0; j < len(visible[i]); j++ {
			if visible[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func second(grid [][]int) {
	scenic := make([][]int, len(grid))
	for i := range scenic {
		scenic[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			scenic[i][j] = scenicValue(grid, i, j)
		}
	}

	max := 0
	// maxI := 0
	// maxJ := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if scenic[i][j] > max {
				max = scenic[i][j]
				// maxI = i
				// maxJ = j
			}
		}
	}
	fmt.Println(max)
	// fmt.Printf("i: %d, j: %d", maxI, maxJ)
}

func scenicValue(grid [][]int, i int, j int) int {
	row := grid[i]

	tallestSoFar := grid[i][j]

	lookRight := 0
	for x := j + 1; x < len(row); x++ {
		lookRight++
		if grid[i][x] >= tallestSoFar {
			break
		}
	}

	lookLeft := 0
	for x := j - 1; x >= 0; x-- {
		lookLeft++
		if grid[i][x] >= tallestSoFar {
			break
		}
	}

	lookDown := 0
	for x := i + 1; x < len(grid); x++ {
		lookDown++
		if grid[x][j] >= tallestSoFar {
			break
		}
	}

	lookUp := 0
	for x := i - 1; x >= 0; x-- {
		lookUp++
		if grid[x][j] >= tallestSoFar {
			break
		}
	}

	// fmt.Printf("Scenic value for (%d, %d): %d\n", i, j, lookLeft*lookRight*lookUp*lookDown)
	// fmt.Printf("lookLeft: %d, lookRight: %d, lookUp: %d, lookDown: %d\n", lookLeft, lookRight, lookUp, lookDown)
	return lookLeft * lookRight * lookUp * lookDown
}

func main() {

	grid := parse("input.txt")

	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}
	switch os.Args[1] {
	case "1":
		first(grid)
	case "2":
		second(grid)
	default:
		first(grid)
	}

}
