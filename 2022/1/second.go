package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	elfCalories := []int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	currentElf := (0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfCalories = append(elfCalories, currentElf)
			currentElf = 0
		} else {
			cal, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentElf += cal
		}
	}
	sort.Ints(elfCalories)
	top3 := elfCalories[len(elfCalories)-3:]
	fmt.Println(top3)
	sum := 0

	for _, c := range top3 {
		sum += c
	}
	fmt.Println(sum)
}
