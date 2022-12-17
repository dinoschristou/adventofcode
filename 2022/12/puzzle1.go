package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cell struct {
	Row, Col int
	Val      int
	Distance int
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

type Heightmap [][]*Cell

func main() {
	matrix := loadFile(os.Args[1])

	hmap := make(Heightmap, len(matrix))
	startCell := &Cell{}
	endCell := &Cell{}
	for i, row := range matrix {
		hmap[i] = make([]*Cell, len(row))
		for j, _ := range row {
			c := Cell{Row: i, Col: j, Distance: 0}
			if matrix[i][j] == 'S' {
				startCell = &c
				startCell.Val = 0
			} else if matrix[i][j] == 'E' {
				endCell = &c
				endCell.Val = 25
			} else {
				c.Val = int(matrix[i][j] - 'a')
			}
			hmap[i][j] = &c
		}
	}

	queue := []*Cell{}
	queue = append(queue, startCell)

	visitedCells := map[*Cell]bool{}
	visitedCells[startCell] = true
	for len(queue) > 0 {
		current := queue[0]

		if current == endCell {
			fmt.Println("Found the end")
			break
		}

		adjacentCells := hmap.GetMoves(current)

		for _, cell := range adjacentCells {
			if _, ok := visitedCells[cell]; !ok {
				queue = append(queue, cell)
				visitedCells[cell] = true
				cell.Distance = current.Distance + 1
			}
		}
		queue = queue[1:]
	}

	fmt.Printf("Visited %d cells\n", len(visitedCells))
	fmt.Printf("Distance to end: %d\n", endCell.Distance)

}

func (hmap Heightmap) GetMoves(c *Cell) []*Cell {
	moves := []*Cell{}
	if c.Row-1 >= 0 && hmap[c.Row-1][c.Col].Val <= c.Val+1 {
		moves = append(moves, hmap[c.Row-1][c.Col])
	}
	if c.Col+1 < len(hmap[0]) && hmap[c.Row][c.Col+1].Val <= c.Val+1 {
		moves = append(moves, hmap[c.Row][c.Col+1])
	}
	if c.Col-1 >= 0 && hmap[c.Row][c.Col-1].Val <= c.Val+1 {
		moves = append(moves, hmap[c.Row][c.Col-1])
	}
	if c.Row+1 < len(hmap) && hmap[c.Row+1][c.Col].Val <= c.Val+1 {
		moves = append(moves, hmap[c.Row+1][c.Col])
	}
	return moves
}
func loadFile(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	matrix := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}
	return matrix
}

func print(matrix [][]rune, cell Cell) {
	for i, row := range matrix {
		for j, _ := range row {
			if i == cell.Row && j == cell.Col {
				fmt.Printf(Red + "X" + Reset)
			} else {
				fmt.Printf(string(matrix[i][j]))
			}
		}
		fmt.Println()
	}
}
