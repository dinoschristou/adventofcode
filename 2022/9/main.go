package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func adjacent(hx, hy, tx, ty int) bool {
	return (hx == tx && hy == ty) || (hx == tx && hy == ty+1) || (hx == tx && hy == ty-1) || (hx == tx+1 && hy == ty) || (hx == tx-1 && hy == ty) || (hx == tx+1 && hy == ty+1) || (hx == tx-1 && hy == ty-1) || (hx == tx+1 && hy == ty-1) || (hx == tx-1 && hy == ty+1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printVisited(visited map[Point]bool) {
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}

	fmt.Printf("Visited %v\n", count)
}

func main() {
	file, err := os.Open("input.txt")
	verbose := false

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}
	switch os.Args[1] {
	case "1":
		run(lines, 1, verbose)
	case "2":
		run(lines, 9, verbose)
	default:
		run(lines, 1, verbose)
	}
}

type Knot struct {
	id      int
	head    *Knot
	tail    *Knot
	x, y    int
	visited map[Point]bool
}

func run(lines []string, tailKnotCount int, verbose bool) {
	headKnot, tailKnot := buildKnots(tailKnotCount)
	printKnots(headKnot)
	for _, line := range lines {

		components := strings.Split(line, " ")
		dir := components[0]
		distance, _ := strconv.Atoi(components[1])

		for i := 0; i < distance; i++ {
			switch dir {
			case "R":
				headKnot.x++
			case "L":
				headKnot.x--
			case "U":
				headKnot.y++
			case "D":
				headKnot.y--
			}
			headKnot.visited[Point{x: headKnot.x, y: headKnot.y}] = true
			moveTail(headKnot.tail)
		}
		if verbose {
			printVisited(headKnot.visited)
		}
	}
	printVisited(tailKnot.visited)
}

func printKnots(knot *Knot) {
	if knot == nil {
		return
	}
	fmt.Printf("Knot %v at (%v,%v)\n", knot.id, knot.x, knot.y)
	printKnots(knot.tail)
}

func moveTail(t *Knot) {
	if t == nil {
		return
	}
	if !adjacent(t.head.x, t.head.y, t.x, t.y) {
		if t.head.x > t.x {
			t.x++
		} else if t.head.x < t.x {
			t.x--
		}

		if t.head.y > t.y {
			t.y++
		}
		if t.head.y < t.y {
			t.y--
		}
		moveTail(t.tail)
	}
	t.visited[Point{x: t.x, y: t.y}] = true
	return
}

func buildKnots(knotCount int) (*Knot, *Knot) {
	headKnot := &Knot{
		id:      0,
		head:    nil,
		tail:    nil,
		x:       0,
		y:       0,
		visited: make(map[Point]bool),
	}
	prevKnot := headKnot
	for i := 1; i <= knotCount; i++ {
		prevKnot.tail = &Knot{
			id:      i,
			head:    prevKnot,
			tail:    nil,
			x:       0,
			y:       0,
			visited: make(map[Point]bool),
		}
		prevKnot = prevKnot.tail
	}

	return headKnot, prevKnot
}
