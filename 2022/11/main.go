package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id              int
	items           []int
	operation       []string
	test            int
	trueThrow       int
	falseThrow      int
	inspectionCount int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rounds := 20
	puzzle := 1
	if len(os.Args) > 2 {
		if os.Args[1] >= "2" {
			puzzle = 2
			rounds, _ = strconv.Atoi(os.Args[2])
		}
	}

	scanner := bufio.NewScanner(file)
	headerRegex := regexp.MustCompile(`Monkey (\d+):`)
	startingItemsRegex := regexp.MustCompile(`\s+Starting items:\s(.*)`)
	operationRegex := regexp.MustCompile(`\s+Operation: new = ([a-z0-9]+) ([*|+]) ([a-z0-9]+)`)
	testRegex := regexp.MustCompile(`\s+Test: divisible by (\d+)`)
	trueThrowRegex := regexp.MustCompile(`\s+If true: throw to monkey (\d+)`)
	falseThrowRegex := regexp.MustCompile(`\s+If false: throw to monkey (\d+)`)
	monkeys := []*Monkey{}
	var currentMonkey *Monkey

	for scanner.Scan() {

		line := scanner.Text()
		if matches := headerRegex.FindStringSubmatch(line); matches != nil {
			id, _ := strconv.Atoi(matches[1])
			currentMonkey = &Monkey{id: id}
			monkeys = append(monkeys, currentMonkey)
		}
		if matches := startingItemsRegex.FindStringSubmatch(line); matches != nil {
			for i, item := range matches {
				if i > 0 {
					sanitiseString := strings.Replace(item, " ", "", -1)
					itemStrings := strings.Split(sanitiseString, ",")
					for _, itemString := range itemStrings {
						item, _ := strconv.Atoi(itemString)
						currentMonkey.items = append(currentMonkey.items, item)
					}
				}
			}
		}
		if matches := operationRegex.FindStringSubmatch(line); matches != nil {
			for i, item := range matches {
				if i > 0 {
					currentMonkey.operation = append(currentMonkey.operation, item)
				}
			}
		}
		if matches := testRegex.FindStringSubmatch(line); matches != nil {
			for i, item := range matches {
				if i > 0 {
					test, _ := strconv.Atoi(item)
					currentMonkey.test = test
				}
			}
		}
		if matches := trueThrowRegex.FindStringSubmatch(line); matches != nil {
			for i, item := range matches {
				if i > 0 {
					trueThrow, _ := strconv.Atoi(item)
					currentMonkey.trueThrow = trueThrow
				}
			}
		}
		if matches := falseThrowRegex.FindStringSubmatch(line); matches != nil {
			for i, item := range matches {
				if i > 0 {
					falseThrow, _ := strconv.Atoi(item)
					currentMonkey.falseThrow = falseThrow
				}
			}
		}
	}
	printMonkeys(monkeys)

	divisor := 1
	if puzzle == 2 {
		for _, m := range monkeys {
			divisor *= m.test
		}
		fmt.Printf("Divisor: %v\n", divisor)
		play(monkeys, rounds, 1, divisor)
	} else {
		play(monkeys, rounds, 3, divisor)
	}
	printMonkeys(monkeys)
}

func play(monkeys []*Monkey, rounds int, worrylevelReducer int, divisor int) {
	//scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			moves := monkey.round(i, worrylevelReducer, divisor)

			for _, move := range moves {
				monkeys[move.target].items = append(monkeys[move.target].items, move.item)
			}
		}
		//printMonkeys(monkeys)
		//scanner.Scan()
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCount < monkeys[j].inspectionCount
	})

	mc := len(monkeys)
	fmt.Printf("Final Results: %v * %v = %v\n", monkeys[mc-2].inspectionCount, monkeys[mc-1].inspectionCount, monkeys[mc-2].inspectionCount*monkeys[mc-1].inspectionCount)
}
func printMonkeys(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		fmt.Println(monkey)
	}
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d:\n Items: %v\n Operations: %v \n Test: %v\n True: %v\n False: %v\n Inspection Count: %v\n", m.id, m.items, m.operation, m.test, m.trueThrow, m.falseThrow, m.inspectionCount)
}

func (m *Monkey) worryLevel(i int) int {
	baseWorryLevel := m.items[i]
	leftOperand, rightOperand := 0, 0
	if m.operation[0] == "old" {
		leftOperand = baseWorryLevel
	} else {
		leftOperand, _ = strconv.Atoi(m.operation[0])
	}
	if m.operation[2] == "old" {
		rightOperand = baseWorryLevel
	} else {
		rightOperand, _ = strconv.Atoi(m.operation[2])
	}

	if m.operation[1] == "*" {
		return leftOperand * rightOperand
	} else {
		return leftOperand + rightOperand
	}
}

func (m *Monkey) testWorryLevel(i int, worrylevelReducer int, divisor int) (bool, int) {
	finalWorryLevel := 0
	if worrylevelReducer == 3 {
		finalWorryLevel = m.worryLevel(i) / worrylevelReducer
	} else {
		finalWorryLevel = m.worryLevel(i) % divisor
	}

	return finalWorryLevel%m.test == 0, finalWorryLevel
}

type Move struct {
	item   int
	target int
}

func (m *Monkey) round(i int, worrylevelReducer int, divisor int) []Move {
	moves := make([]Move, len(m.items))
	for i := 0; i < len(m.items); i++ {
		branch, newWorryLevel := m.testWorryLevel(i, worrylevelReducer, divisor)
		move := Move{
			item: newWorryLevel,
		}
		if branch {
			move.target = m.trueThrow
		} else {
			move.target = m.falseThrow
		}
		moves[i] = move
	}
	m.items = nil
	m.inspectionCount += len(moves)
	return moves
}
