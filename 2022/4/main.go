package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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

func first() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += fullOverlap(line)
	}
	fmt.Println(totalScore)
}

func ranges(s string) (int, int, int, int) {
	r := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	matches := r.FindStringSubmatch(s)
	e1s, _ := strconv.Atoi(matches[1])
	e1e, _ := strconv.Atoi(matches[2])
	e2s, _ := strconv.Atoi(matches[3])
	e2e, _ := strconv.Atoi(matches[4])
	return e1s, e1e, e2s, e2e
}

func partialOverlap(s string) int {
	e1s, e1e, e2s, e2e := ranges(s)

	if e1s <= e2e && e2s <= e1e {
		return 1
	}
	return 0
}
func fullOverlap(s string) int {
	e1s, e1e, e2s, e2e := ranges(s)

	if e1s <= e2s && e1e >= e2e {
		return 1
	}
	if e2s <= e1s && e2e >= e1e {
		return 1
	}
	return 0
}

func second() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += partialOverlap(line)
	}
	fmt.Println(totalScore)
}
