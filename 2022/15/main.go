package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Point struct {
	x, y int
}

type Sensor struct {
	s, b Point
}

type Grid struct {
	sensors []Sensor
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Sensor at x=(\d+), y=(\d+): closest beacon is at x=([-]?\d+), y=(\d+)`)
	g := Grid{
		sensors: make([]Sensor, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			sx, _ := strconv.Atoi(matches[1])
			sy, _ := strconv.Atoi(matches[2])
			bx, _ := strconv.Atoi(matches[3])
			by, _ := strconv.Atoi(matches[4])
			g.sensors = append(g.sensors, Sensor{Point{sx, sy}, Point{bx, by}})
		}
	}
	c := g.NoBeacons(2000000)
	fmt.Printf("No beacons: %d\n", c)

	beacon, _ := g.findBeacon(4000000)
	fmt.Printf("Beacon: %v\n", beacon)

	bigx := big.NewInt(int64(beacon.x))
	bigy := big.NewInt(int64(beacon.y))
	res := bigx.Mul(bigx, big.NewInt(4000000))
	res = res.Add(res, bigy)
	fmt.Printf("Result: %v\n", res)
}

func (g Grid) IsBeacon(p Point) bool {
	for _, s := range g.sensors {
		if s.b == p {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Interval struct {
	min, max int
}

func (g Grid) nearbySensors(row int) []Sensor {
	nearby := []Sensor{}
	for _, s := range g.sensors {

		if row <= s.s.y+s.ManhattanDistance(s.b) && row >= s.s.y-s.ManhattanDistance(s.b) {
			nearby = append(nearby, s)
		}
	}
	return nearby
}

func (g Grid) findBeacon(limit int) (Point, error) {

	for row := 0; row <= limit; row++ {
		intervals := g.intervals(row)
		for i := 1; i < len(intervals); i++ {
			if intervals[i].max < 0 {
				continue
			}
			if intervals[i].min > limit {
				continue
			}

			if intervals[i].min == intervals[i-1].max+2 {
				p := Point{intervals[i-1].max + 1, row}
				if !g.IsBeacon(p) {
					return p, nil
				}
			}
		}
	}
	return Point{}, fmt.Errorf("no beacon found")
}

func (g Grid) intervals(row int) []Interval {
	ns := g.nearbySensors(row)
	intervals := []Interval{}
	for _, s := range ns {
		md := s.ManhattanDistance(s.b)
		vd := abs(row - s.s.y)
		xadj := md - vd
		intervals = append(intervals, Interval{s.s.x - xadj, s.s.x + xadj})
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].min < intervals[j].min
	})
	return compressIntervals(intervals)
}

func compressIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	res := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i].min <= res[len(res)-1].max+1 {
			res[len(res)-1].max = max(res[len(res)-1].max, intervals[i].max)
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func (s Sensor) WithinRange(p Point) bool {
	return s.ManhattanDistance(p) <= s.ManhattanDistance(s.b)
}
func (s Sensor) String() string {
	return fmt.Sprintf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", s.s.x, s.s.y, s.b.x, s.b.y)
}

func (s Sensor) ManhattanDistance(p Point) int {
	return abs(s.s.x-p.x) + abs(s.s.y-p.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g Grid) NoBeacons(y int) int {
	intervals := g.intervals(y)
	count := 0
	minx, _, maxx, _ := g.extremities()

	for i := minx; i <= maxx; i++ {
		p := Point{i, y}
		if g.IsBeacon(p) {
			continue
		}
		for _, interval := range intervals {
			if interval.min <= i && i <= interval.max {
				count += 1
				break
			}
		}

	}
	return count
}

func (g Grid) extremities() (int, int, int, int) {
	minX, minY, maxX, maxY := 2147483647, 2147483647, 0, 0

	for _, s := range g.sensors {
		if s.s.x-s.ManhattanDistance(s.b) < minX {
			minX = s.s.x - s.ManhattanDistance(s.b)
		}
		if s.s.y-s.ManhattanDistance(s.b) < minY {
			minY = s.s.y - s.ManhattanDistance(s.b)
		}
		if s.s.x+s.ManhattanDistance(s.b) > maxX {
			maxX = s.s.x + s.ManhattanDistance(s.b)
		}
		if s.s.y+s.ManhattanDistance(s.b) > maxY {
			maxY = s.s.y + s.ManhattanDistance(s.b)
		}
	}
	return minX, minY, maxX, maxY
}
