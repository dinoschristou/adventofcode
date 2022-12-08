package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
	IsDir    bool
	Size     int
}

func (n *Node) AddChild(name string, size int, isDir bool) {
	for _, child := range n.Children {
		if child.Name == name {
			return
		}
	}
	newNode := &Node{
		Name:     name,
		Parent:   n,
		Children: []*Node{},
		IsDir:    isDir,
		Size:     size,
	}
	n.Children = append(n.Children, newNode)
	return
}

func (n *Node) PrettyPrint(indents int) {

	if n.IsDir {
		fmt.Printf("%v- %v (dir)\n", strings.Repeat(" ", indents), n.Name)
	} else {
		fmt.Printf("%v- %v  (file, size=%v)\n", strings.Repeat(" ", indents), n.Name, n.Size)
	}
	for _, child := range n.Children {
		child.PrettyPrint(indents + 2)
	}
}

func (n *Node) TotalSize() int {
	if n.IsDir {
		size := 0
		for _, child := range n.Children {
			size += child.TotalSize()
		}
		return size
	} else {
		return n.Size
	}
}

type NodeFilter func(*Node) bool

func (n *Node) FlattenDirectories() []*Node {
	dirs := []*Node{}
	if n.IsDir {
		dirs = append(dirs, n)
		for _, child := range n.Children {
			dirs = append(dirs, child.FlattenDirectories()...)
		}
	}
	return dirs
}

func (n *Node) DirectoriesOfNote(fn NodeFilter) []*Node {
	dirs := []*Node{}
	for _, child := range n.Children {
		if child.IsDir {
			if fn(child) {
				dirs = append(dirs, child)
			}
			dirs = append(dirs, child.DirectoriesOfNote(fn)...)
		}
	}
	return dirs
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	rootNode := &Node{
		Name:     "/",
		Parent:   nil,
		Children: []*Node{},
		IsDir:    true,
		Size:     0,
	}

	fmt.Printf("%v\n", rootNode.Name)

	currentNode := rootNode

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Evaluating: %v\n", line)
		if string(line[0]) == "$" {
			components := strings.Split(line, " ")
			if components[1] == "cd" {
				if components[2] == ".." {
					fmt.Printf("Going up one level\n")
					currentNode = currentNode.Parent
				} else if components[2] == "/" {
					fmt.Printf("Going to root\n")
					currentNode = rootNode
				} else {
					fmt.Printf("Going to %v\n", components[2])
					childAlreadyVisited := false
					for _, child := range currentNode.Children {
						if child.Name == components[2] {
							currentNode = child
							childAlreadyVisited = true
						}
					}
					if !childAlreadyVisited {
						newNode := &Node{
							Name:     components[2],
							Parent:   currentNode,
							Children: []*Node{},
							IsDir:    true,
							Size:     0,
						}
						currentNode.Children = append(currentNode.Children, newNode)
						currentNode = newNode
					}
				}
			} else if components[1] == "ls" {
				fmt.Printf("Listing %v\n", currentNode.Name)
			}
			// Parse command

		} else {
			components := strings.Split(line, " ")
			if components[0] == "dir" {
				currentNode.AddChild(components[1], 0, true)
			} else {
				size, _ := strconv.Atoi(components[0])
				currentNode.AddChild(components[1], size, false)
			}
		}
	}
	rootNode.PrettyPrint(0)

	dirs := rootNode.DirectoriesOfNote(func(n *Node) bool {
		return n.TotalSize() < 100000
	})

	sum := 0
	for d := range dirs {
		sum += dirs[d].TotalSize()
	}
	fmt.Printf("Total size of dirs of which each are <100000: %v\n", sum)

	targetSize := 30000000 - (70000000 - rootNode.TotalSize())
	fmt.Printf("Target size: %v\n", targetSize)
	filter := func(n *Node) bool {
		return n.TotalSize() > targetSize
	}
	allDirs := rootNode.FlattenDirectories()
	var candidateDir *Node

	for _, dir := range allDirs {
		if candidateDir == nil {
			candidateDir = dir
			continue
		}
		if filter(dir) && dir.TotalSize() < candidateDir.TotalSize() {
			candidateDir = dir
		}
	}
	fmt.Printf("Candidate dir: %v with size %v\n", candidateDir.Name, candidateDir.TotalSize())
}
