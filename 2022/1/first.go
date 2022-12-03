package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	maxCalories := int64(0)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	currentElf := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentElf > maxCalories {
				maxCalories = currentElf
			}
			currentElf = 0
		} else {
			cal, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			currentElf += cal
		}

	}
	fmt.Println(maxCalories)
}
