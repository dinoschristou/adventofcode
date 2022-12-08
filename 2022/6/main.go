package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	messageMarker := false
	if len(os.Args) > 1 && os.Args[1] == "message" {
		messageMarker = true
	}
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileString := string(file)

	for pos, _ := range fileString {
		if pos < 3 || (messageMarker && pos < 13) {
			continue
		}
		subList := ""

		if messageMarker {
			subList = fileString[pos-13 : pos+1]
		} else {
			subList = fileString[pos-3 : pos+1]
		}

		if unique(subList) {
			fmt.Printf("Answer is: %v\n", pos+1)
			break
		}
	}
}

func unique(chars string) bool {
	seen := make(map[rune]struct{}, len(chars))
	for _, v := range chars {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}
