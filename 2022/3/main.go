package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func first() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += score(line)
	}
	fmt.Println(totalScore)
}

func second() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	totalScore := 0
	count := 0
	currGroup := []string{}

	for scanner.Scan() {
		if count > 0 && count%3 == 0 {
			totalScore += scoreGroup(currGroup)
			currGroup = []string{}
		}
		line := scanner.Text()
		currGroup = append(currGroup, line)
		count++
	}
	totalScore += scoreGroup(currGroup)
	fmt.Println(totalScore)
}

func intersect(a, b string) string {
	result := ""
	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				result += string(c1)
			}
		}
	}
	return result
}

func scoreGroup(currGroup []string) int {
	temp := intersect(currGroup[0], currGroup[1])
	badge := intersect(temp, currGroup[2])
	return priority(rune(badge[0]))
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Invalid number of arguments")
	}
	switch os.Args[1] {
	case "1":
		first()
	case "2":
		second()
	default:
		first()
	}

}

func score(line string) int {
	comp1 := line[:len(line)/2]
	comp2 := line[len(line)/2:]

	for _, c1 := range comp1 {
		for _, c2 := range comp2 {
			if c1 == c2 {
				return priority(c1)
			}
		}
	}
	return 0
}

func priority(c rune) int {
	if c >= 97 && c <= 122 {
		return int(c) - 96
	}
	return int(c) - 64 + 26
}
