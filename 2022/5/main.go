package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}
	switch os.Args[1] {
	case "1":
		moveCrates(false)
	case "2":
		moveCrates(true)
	default:
		moveCrates(false)
	}

}

func moveCrates(poshCrane bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	serialisedState := []string{}
	var state [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			serialisedState = append(serialisedState, line)
		} else {
			state = buildModel(serialisedState)
			break
		}
	}

	fmt.Println(state)

	for scanner.Scan() {
		line := scanner.Text()
		move(line, state, poshCrane)
	}

	fmt.Println(strings.Join(topBoxes(state), ""))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func topBoxes(state [][]string) []string {
	fmt.Println(len(state))
	topboxes := []string{}
	for i := 0; i < len(state); i++ {
		if len(state[i]) > 0 {
			topboxes = append(topboxes, state[i][len(state[i])-1])
		}
	}
	return topboxes
}

func move(move string, state [][]string, poshCrane bool) [][]string {
	// Decode the move: e.g. "move 2 from 4 to 9"
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := r.FindStringSubmatch(move)
	numberOfCrates, _ := strconv.Atoi(matches[1])
	fromPile, _ := strconv.Atoi(matches[2])
	toPile, _ := strconv.Atoi(matches[3])

	// zero base the piles
	fromPile--
	toPile--

	if poshCrane {
		cratesToMove := state[fromPile][len(state[fromPile])-numberOfCrates:]
		state[fromPile] = state[fromPile][:len(state[fromPile])-numberOfCrates]
		state[toPile] = append(state[toPile], cratesToMove...)
		fmt.Println("use the posh crane")
	} else {
		for i := 0; i < numberOfCrates; i++ {
			state[toPile] = append(state[toPile], state[fromPile][len(state[fromPile])-1])
			state[fromPile] = state[fromPile][:len(state[fromPile])-1]
		}
	}
	return state
}

func buildModel(serialisedState []string) [][]string {
	pilesLine := strings.Fields(serialisedState[len(serialisedState)-1])
	numberOfPiles, _ := strconv.Atoi(pilesLine[len(pilesLine)-1])
	fmt.Println("Number of piles: ", numberOfPiles)

	state := make([][]string, numberOfPiles)
	for i := len(serialisedState) - 2; i >= 0; i-- {
		line := serialisedState[i]
		//fmt.Printf("Stacking line: %v\n", line)
		for j := 0; j < numberOfPiles; j++ {
			startIndex := j * 4 // there is redundant whitespace at the end of lines
			//substring := line[startIndex : startIndex+3]
			//fmt.Printf("Substring: %v, ", substring)
			component := string(line[startIndex : startIndex+3][1])
			//fmt.Printf("Component: %v\n", component)
			if component != " " {
				state[j] = append(state[j], component)
			}
		}
	}
	return state
}

func second() {
}
