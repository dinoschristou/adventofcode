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

type Pair struct {
	left  string
	right string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	pairs := []Pair{}
	p := Pair{}
	for i, line := range lines {
		if i%2 == 0 {
			p = Pair{
				left: line,
			}
		} else {
			p.right = line
			pairs = append(pairs, p)
		}
	}

	indices := []int{}

	for i, p := range pairs {
		if p.InOrder() {
			indices = append(indices, i+1)
		}
	}

	fmt.Printf("In order %v\n", indices)

	sum := 0
	for _, i := range indices {
		sum += i
	}
	fmt.Printf("Sum %v\n", sum)

	messages := []any{}
	for _, p := range pairs {
		messages = append(messages, parseMessage(p.left))
		messages = append(messages, parseMessage(p.right))
	}
	messages = append(messages, parseMessage("[[2]]"))
	messages = append(messages, parseMessage("[[6]]"))

	sort.Slice(messages, func(i, j int) bool {
		return compareMessages(messages[i].([]any), messages[j].([]any)) == -1
	})

	marker1 := 0
	marker2 := 0

	for i, m := range messages {
		s := fmt.Sprintf("%v", m)
		if s == "[[2]]" {
			marker1 = i + 1
		} else if s == "[[6]]" {
			marker2 = i + 1
		}
		fmt.Println(s)
	}

	fmt.Printf("Decoder signal is %v * %v = %v\n", marker1, marker2, marker1*marker2)
}

func (p Pair) InOrder() bool {

	l := parseMessage(p.left)
	r := parseMessage(p.right)

	ls := fmt.Sprintf("%v", l)
	rs := fmt.Sprintf("%v", r)
	sanitisedLeft := strings.ReplaceAll(p.left, ",", " ")
	sanitisedRight := strings.ReplaceAll(p.right, ",", " ")

	fmt.Printf("%v converted to \n%v\n", sanitisedLeft, l)
	fmt.Printf("%v converted to \n%v\n", sanitisedRight, r)

	if sanitisedLeft != ls || sanitisedRight != rs {
		log.Fatal("sanitised messages don't match\n")
	}

	return compareMessages(l, r) == -1

}

func compareMessages(l []any, r []any) int {
	for i, v := range l {
		// if right is shorter than left, it's not in order
		if i >= len(r) {
			return +1
		}
		// if left is an int
		if li, ok := v.(int); ok {
			// if right is an int
			if ri, ok := r[i].(int); ok {
				// simply compare ints
				if li < ri {
					return -1
				} else if li > ri {
					return +1
				}
				// if right is a list
			} else {
				// put left in a list and compare
				if c := compareMessages([]any{li}, r[i].([]any)); c != 0 {
					return c
				}
			}
		} else {
			// left is a list
			// if right is an int
			if ri, ok := r[i].(int); ok {
				// put right in a list and compare
				if c := compareMessages(v.([]any), []any{ri}); c != 0 {
					return c
				}
			} else {
				if c := compareMessages(v.([]any), r[i].([]any)); c != 0 {
					return c
				}
			}
		}
	}
	if len(l) == len(r) {
		return 0
	}
	return -1
}

func parseMessage(message string) []any {

	listStack := []any{}

	currentList := []any{}
	intCharacters := []rune{}

	for i, c := range message {
		if c == '[' && i > 0 {
			// push current list to stack and create a new list
			listStack = append(listStack, currentList)
			currentList = []any{}
		} else if c == ']' {
			// if we are finishing a character string, add it to the current list
			if len(intCharacters) > 0 {
				val, err := strconv.Atoi(string(intCharacters))
				if err != nil {
					log.Fatal(err)
				}
				currentList = append(currentList, val)
			}

			// pop the stack and set the current list to the top of the stack
			if len(listStack) > 0 {
				parentList := listStack[len(listStack)-1].([]any)
				listStack = listStack[:len(listStack)-1]
				parentList = append(parentList, currentList)
				currentList = parentList
			}

			intCharacters = []rune{}
		} else if c == ',' {
			if len(intCharacters) > 0 {
				val, err := strconv.Atoi(string(intCharacters))
				if err != nil {
					log.Fatal(err)
				}
				currentList = append(currentList, val)
			}
			intCharacters = []rune{}
		} else if c == ' ' {
			continue
		} else if c >= '0' && c <= '9' {
			intCharacters = append(intCharacters, c)
		}
	}
	return currentList
}
