package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatal("Invalid line")
		}

		totalScore += score(parse_first(parts[0]), parse_first(parts[1]))
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
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatal("Invalid line")
		}
		p1 := parse_move(parts[0])
		r := parse_result(parts[1])

		totalScore += score_second(p1, r)
	}
	fmt.Println(totalScore)
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

type Move int

const (
	rock Move = iota + 1
	paper
	scissors
)

type Result int

const (
	win  Result = 6
	draw Result = 3
	lose Result = 0
)

func parse_move(p string) Move {
	switch p {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	default:
		log.Fatal("Invalid move")
	}
	return 0
}

func parse_result(p string) Result {
	switch p {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	default:
		log.Fatal("Invalid result")
	}
	return 0
}

func parse_first(p string) Move {
	switch p {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		log.Fatal("Invalid move")
	}
	return 0
}

func score_second(p1 Move, r Result) int {
	score := 0
	if p1 == rock && r == lose {
		score += 0 + int(scissors)
	}
	if p1 == rock && r == draw {
		score += 3 + int(rock)
	}
	if p1 == rock && r == win {
		score += 6 + int(paper)
	}
	if p1 == paper && r == lose {
		score += 0 + int(rock)
	}
	if p1 == paper && r == draw {
		score += 3 + int(paper)
	}
	if p1 == paper && r == win {
		score += 6 + int(scissors)
	}
	if p1 == scissors && r == lose {
		score += 0 + int(paper)
	}
	if p1 == scissors && r == draw {
		score += 3 + int(scissors)
	}
	if p1 == scissors && r == win {
		score += 6 + int(rock)
	}
	return score
}

func score(p1 Move, p2 Move) int {
	score := 0
	if p1 == rock && p2 == scissors {
		score += 0
	}
	if p1 == rock && p2 == paper {
		score += 6
	}
	if p1 == rock && p2 == rock {
		score += 3
	}
	if p1 == paper && p2 == rock {
		score += 0
	}
	if p1 == paper && p2 == paper {
		score += 3
	}
	if p1 == paper && p2 == scissors {
		score += 6
	}
	if p1 == scissors && p2 == rock {
		score += 6
	}
	if p1 == scissors && p2 == paper {
		score += 0
	}
	if p1 == scissors && p2 == scissors {
		score += 3
	}
	return score + int(p2)
}
