package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	run(lines)
}

func run(lines []string) {
	cycles := make(map[int]int)

	cycle := 0
	register := 1

	for _, line := range lines {
		//fmt.Printf("Line: %v\n", line)
		if line == "noop" {
			cycle++
			cycles[cycle] = register
		} else {
			components := strings.Split(line, " ")
			if components[0] == "addx" {
				increment, _ := strconv.Atoi(components[1])
				cycle++
				//fmt.Printf("Cycle %v, register remains %v\n", cycle, register)
				cycles[cycle] = register
				cycle++
				register = register + increment
				//fmt.Printf("Cycle %v, register increments to %v\n", cycle, register)
				cycles[cycle] = register
			}
		}
	}
	//print(cycles)
	printInterestingCycles(cycles, []int{20, 60, 100, 140, 180, 220})

	framebuffer := make([]string, 240)
	for i := range framebuffer {
		framebuffer[i] = "."
	}
	printFrameBuffer(framebuffer)

	cycles[0] = 1
	for cycle := 1; cycle <= 240; cycle++ {
		spritePosition := cycles[cycle-1]
		pixel := cycle - 1
		printSprite(spritePosition)

		fmt.Printf("During cycle %v: CRT draws pixel in position %v\n", cycle, pixel)

		if overlap(spritePosition, pixel%40) {
			framebuffer[pixel] = "#"
		}

		printFrameBuffer(framebuffer)
		//scanner := bufio.NewScanner(os.Stdin)
		//scanner.Scan()

	}

	printFrameBuffer(framebuffer)
}

func overlap(spritePosition int, pixel int) bool {
	return spritePosition-1 <= pixel && spritePosition+1 >= pixel
}
func printSprite(x int) {
	fmt.Printf("Sprite position: %v\n", x)
	for i := 0; i < 39; i++ {
		if i >= x-1 && i <= x+1 {
			fmt.Printf("# ")
		} else {
			fmt.Printf(". ")
		}
	}
	fmt.Printf("\n")
}
func printFrameBuffer(framebuffer []string) {
	for i, pixel := range framebuffer {
		if i%40 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%v ", pixel)
	}
	fmt.Printf("\n")
}

func print(cycles map[int]int) {
	keys := []int{}

	for k := range cycles {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, cycles[k])
	}
}

func printInterestingCycles(cycles map[int]int, interestingCycles []int) {
	sum := 0
	for _, cycle := range interestingCycles {
		if val, ok := cycles[cycle-1]; !ok {
			continue
		} else {
			fmt.Printf("%v: %v\n", cycle, val)
			sum += val * cycle
		}
	}

	fmt.Printf("Sum: %v\n", sum)
}
