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

type Grid struct {
	blockedPoints map[Point]bool
	minX, maxX    int
}

func (p Point) String() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func (p Point) between(other Point) []Point {
	result := []Point{}
	if p.x == other.x {
		if p.y < other.y {
			for i := p.y; i <= other.y; i++ {
				result = append(result, Point{p.x, i})
			}
			return result
		} else {
			for i := other.y; i <= p.y; i++ {
				result = append(result, Point{p.x, i})
			}
			return result
		}
	} else {
		if p.x < other.x {
			for i := p.x; i <= other.x; i++ {
				result = append(result, Point{i, p.y})
			}
			return result
		} else {
			for i := other.x; i <= p.x; i++ {
				result = append(result, Point{i, p.y})
			}
			return result
		}
	}
}

func (g Grid) String() string {
	s := ""
	maxY, minY, maxX, minX := g.extremities()

	for i := minY; i <= maxY; i++ {
		s += strconv.Itoa(i) + " "
		for j := minX; j <= maxX; j++ {
			if g.blockedPoints[Point{j, i}] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (g Grid) extremities() (int, int, int, int) {
	maxY, minY, maxX, minX := 0, 0, 0, 10000000

	for k, _ := range g.blockedPoints {
		if k.y > maxY {
			maxY = k.y
		}
		if k.y <= minY {
			minY = k.y
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.x < minX {
			minX = k.x
		}
	}
	return maxY, minY, maxX, minX
}

func (g Grid) abyss(p Point) bool {
	return p.x > g.maxX || p.x < g.minX
}

func (g Grid) nextMove(p Point) (Point, error) {
	newPoint := Point{p.x, p.y + 1}
	if g.blockedPoints[newPoint] {
		newPoint = Point{p.x - 1, p.y + 1}
		if g.blockedPoints[newPoint] {
			newPoint = Point{p.x + 1, p.y + 1}
			if g.blockedPoints[newPoint] {
				return Point{}, fmt.Errorf("No more moves")
			}
		}
	}
	return newPoint, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	g := Grid{blockedPoints: make(map[Point]bool)}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pointStrs := strings.Split(line, " -> ")

		points := []Point{}
		for _, pointStr := range pointStrs {
			coordinates := strings.Split(pointStr, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			points = append(points, Point{x, y})
		}

		for i := 0; i < len(points)-1; i++ {
			start := points[i]
			end := points[i+1]
			blockedPoints := start.between(end)
			for _, blockedPoint := range blockedPoints {
				g.blockedPoints[blockedPoint] = true
			}
		}
	}

	_, _, g.maxX, g.minX = g.extremities()
	fmt.Println(g)

	grains := 0
	for {
		err := g.run()
		if err != nil {
			break
		}
		grains++
	}
	fmt.Println(grains)
}

func (g Grid) run() error {
	p := Point{500, 0}
	for {
		newP, err := g.nextMove(p)
		if err != nil {
			g.blockedPoints[p] = true
			break
		}
		if g.abyss(newP) {
			return fmt.Errorf("Abyss")
		}
		p = newP
	}
	return nil
}
