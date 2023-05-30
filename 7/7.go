package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileNode struct {
	Name 		string
	Dir 		bool
	Size 		int
	Parent 		*FileNode
	Children 	[]*FileNode
}

func PrintFileTree(node *FileNode, indent string) {
	fmt.Print(indent + "-" + node.Name)
	if node.Dir {
		fmt.Print(" (dir, size=")
		fmt.Print(node.Size)
		fmt.Print(")")
	} else {
		fmt.Print(" (file, size=")
		fmt.Print(node.Size)
		fmt.Print(")")
	}
	fmt.Print("\n")
	for _, child := range node.Children {
		childIndent := indent + "  "
		PrintFileTree(child, childIndent)
	}
}

func main() {
	// setup read file
	file, err := os.Open("7input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	// setup scanner
	scanner := bufio.NewScanner(file)

	// setup root node
	root := new(FileNode)
	root.Name = "/"
	root.Dir = true

	// track current dir
	currentDir := root

	// Scan through all the commands and build file tree
	for scanner.Scan() {
		line := scanner.Text()

		if (strings.HasPrefix(line, "$")) { // this is a command
			tokens := strings.Split(line, " ")

			if (tokens[1] == "cd") { // we need to change dirs
				if (tokens[2] == "/") { // change back to root
					currentDir = root
				} else if tokens[2] != ".." { // need to step into
					for _, child := range currentDir.Children {
						if child.Name == tokens[2] {
							currentDir = child
						}
					}
				} else { // need to step out of
					currentDir = currentDir.Parent
				}

				if currentDir == nil {
					fmt.Println("An error occured while changing dirs.")
					return
				}
			}
		} else { // not a command, need to add files/dirs to file tree
			tokens := strings.Split(line, " ")
			node := new(FileNode)
			node.Name = tokens[1]
			node.Parent = currentDir
			node.Dir = tokens[0] == "dir"

			if (!node.Dir) {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					fmt.Println("There was an error converting a file size from string to int.")
					return
				}
				node.Size = size
			}

			currentDir.Children = append(currentDir.Children, node)
		}
	}

	calcDirSizes(root)
	// PrintFileTree(root, "")
	fmt.Println("Advent of Code Day 7 Part 1 Result: ", calculateTotalSize(root))
}

func calcDirSizes(node *FileNode) int {
	if !node.Dir {
		return node.Size
	}

	dirTotal := 0

	for _, child := range node.Children {
		childSize := calcDirSizes(child)
		dirTotal += childSize
	}

	node.Size = dirTotal
	return dirTotal
}

func calculateTotalSize(node *FileNode) int {
	var totalSize int

	if node.Dir {
		for _, child := range node.Children {
			totalSize += calculateTotalSize(child)
		}

		if node.Size <= 100000 {
			totalSize += node.Size
		}
	}

	return totalSize
}