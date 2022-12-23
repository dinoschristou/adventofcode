package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	name      string
	flowValue int
	adjacent  []*Node
}

type Cave struct {
	nodes  map[string]*Node
	dist   [][]int
	keyMap map[string]int
}

func (n *Node) String() string {
	res := fmt.Sprintf("Node: %s, flow: %d, adj: ", n.name, n.flowValue)
	for _, a := range n.adjacent {
		res += fmt.Sprintf("%s, ", a.name)
	}
	return res
}

var cache = make(map[string]int)

func dist(n1, n2 *Node) int {
	visited := make(map[string]bool)
	queue := [][]*Node{}
	queue = append(queue, []*Node{n1})
	if n1 == n2 {
		return 0
	}

	if v, ok := cache[n1.name+n2.name]; ok {
		return v
	}
	for len(queue) > 0 {
		currPath := queue[0]
		queue = queue[1:]

		curr := currPath[len(currPath)-1]

		if _, ok := visited[curr.name]; !ok {
			visited[curr.name] = true

			for _, a := range curr.adjacent {
				newPath := make([]*Node, len(currPath))
				copy(newPath, currPath)
				newPath = append(newPath, a)
				queue = append(queue, newPath)
				if a == n2 {
					cache[n1.name+n2.name] = len(newPath) - 1
					return len(newPath) - 1
				}
			}
		}
	}
	cache[n1.name+n2.name] = -1
	return -1
}

func main() {
	f := "input.txt"
	nodes := parseFiles(f)

	ks := []string{}
	for k, _ := range nodes {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	keyMap := make(map[string]int)
	for i, k := range ks {
		keyMap[k] = i
	}

	valvesToOpen := []string{}
	for k, v := range nodes {
		if v.flowValue > 0 {
			valvesToOpen = append(valvesToOpen, k)
		}
	}
	timeLimit := 30
	// pop
	pathCache = make(map[string]int)
	maxPressure, result := getMaxPressure(timeLimit, valvesToOpen, nodes)

	fmt.Println(maxPressure)
	fmt.Println(result)

	fmt.Println(valvesToOpen)
	combinations := Combinations(valvesToOpen)

	pathCache = make(map[string]int)
	partitions := []Partition{}
	for i := 0; i < len(combinations); i++ {
		p := Partition{
			me:       combinations[i],
			elephant: getAlternate(valvesToOpen, combinations[i]),
		}
		fmt.Println(combinations[i])
		partitions = append(partitions, p)
	}
	maxPressure = 0
	fmt.Printf("Number of partitiosns: %d\n", len(partitions))
	for i, p := range partitions {
		if i > 0 && i%1000 == 0 {
			fmt.Printf("Done %d partitions\n", i)
		}
		meP, _ := getMaxPressure(26, p.me, nodes)
		elephantP, _ := getMaxPressure(26, p.elephant, nodes)
		if meP+elephantP > maxPressure {
			maxPressure = meP + elephantP
		}
	}
	fmt.Println(maxPressure)
}

func getAlternate(whole []string, part []string) []string {
	res := []string{}
	for _, w := range whole {
		found := false
		for _, p := range part {
			if w == p {
				found = true
				break
			}
		}
		if !found {
			res = append(res, w)
		}
	}
	return res
}

func Combinations(set []string) (subsets [][]string) {
	length := uint(len(set))
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

type Partition struct {
	me       []string
	elephant []string
}

var pathCache map[string]int

func getMaxPressure(timeLimit int, valvesToOpen []string, nodes map[string]*Node) (int, path) {
	start := "AA"
	maxPressure := 0
	result := path{}
	stack := []PathState{}
	stack = append(stack, PathState{
		path:            path{start},
		minutes_elapsed: 0,
		openValves:      openValves{},
	})

	sort.StringSlice(valvesToOpen).Sort()
	key := strings.Join(valvesToOpen, ",")
	if v, ok := pathCache[key]; ok {
		return v, path{}
	}
	for len(stack) > 0 {

		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currValve := curr.path[len(curr.path)-1]
		if curr.minutes_elapsed >= timeLimit || len(curr.path) == len(valvesToOpen)+1 {
			pressure := 0
			for k, v := range curr.openValves {
				openTime := max(timeLimit-v, 0)
				pressure += nodes[k].flowValue * openTime
			}
			if pressure > maxPressure {
				maxPressure = pressure
				result = curr.path
			}
		} else {
			for _, nextValve := range valvesToOpen {
				if _, ok := curr.openValves[nextValve]; !ok {

					dist := dist(nodes[currValve], nodes[nextValve])
					minutes_elapsed := curr.minutes_elapsed + dist + 1
					if dist == math.MaxInt32 {
						continue
					}
					newOpenValves := curr.openValves.copy()
					newOpenValves[nextValve] = minutes_elapsed
					newPath := make(path, len(curr.path))
					copy(newPath, curr.path)
					newPath = append(newPath, nextValve)

					stack = append(stack, PathState{
						path:            newPath,
						minutes_elapsed: minutes_elapsed,
						openValves:      newOpenValves,
					})
				}
			}
		}

	}

	pathCache[key] = maxPressure
	return maxPressure, result
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

type PathState struct {
	path            []string
	minutes_elapsed int
	openValves      openValves
}

type openValves map[string]int

type path []string

func (o openValves) copy() openValves {
	res := make(openValves)
	for k, v := range o {
		res[k] = v
	}
	return res
}

func printMatrix(distMatrix [][]int) {
	for _, row := range distMatrix {
		for _, v := range row {
			if v == math.MaxInt32 {
				fmt.Printf("∞  ")
			} else {
				fmt.Printf("%2d ", v)
			}
		}
		fmt.Println()
	}
}

func parseFiles(f string) map[string]*Node {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z,\s]+)`)

	nodes := make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if len(matches) > 0 {
			name := matches[1]

			if _, ok := nodes[name]; !ok {
				nodes[name] = &Node{
					name: name,
				}
			}
			current := nodes[name]
			flow, _ := strconv.Atoi(matches[2])
			current.flowValue = flow
			rawAdj := strings.Split(strings.ReplaceAll(matches[3], " ", ""), ",")
			for _, adj := range rawAdj {
				if adj == "" {
					continue
				}
				if _, ok := nodes[adj]; !ok {
					nodes[adj] = &Node{
						name: adj,
					}
				}
				current.adjacent = append(current.adjacent, nodes[adj])
			}

		}
	}
	return nodes
}
