package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileNode struct {
	Name     string
	IsDir    bool
	FileSize int64
	Children []*FileNode
}

func PrintFileTree(node *FileNode, indent string) {
	fmt.Print(indent + "-" + node.Name)
	if node.IsDir {
		fmt.Print(" (dir, ")
		fmt.Print(node.FileSize)
		fmt.Print(")")
	} else {
		fmt.Print(" (file, size=")
		fmt.Print(node.FileSize)
		fmt.Print(")")
	}
	fmt.Print("\n")
	for _, child := range node.Children {
		childIndent := indent + "  "
		PrintFileTree(child, childIndent)
	}
}

func main() {
	root := new(FileNode)
	root.Name = "/"
	root.IsDir = true

	err := buildFileTree(root)
	if err != nil {
		fmt.Println("An error occured while building the file tree: ", err)
	}

	calcDirSizes(root)
	listOfDirs := []int64 {}
	listOfDirs = findDirs(root, listOfDirs)

	for i, dir := range listOfDirs {
		fmt.Println(i, ": ", dir)
	}

	var total int64
	total = 0
	for _, dir := range listOfDirs {
		total += dir
	}

	PrintFileTree(root, "")
	fmt.Println(total)

}

func buildFileTree(root *FileNode) error {
	currentLocation := root
	
	file, err := os.Open("7input.txt")
	if err != nil {
		return errors.New("error opening file")
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			tokens := strings.Split(line, " ")
			if tokens[1] == "ls" {
				continue
			} else {
				currentLocation = handleCommand(tokens[2], root, currentLocation)
			}
		} else {
			tokens := strings.Split(line, " ")
			node := new(FileNode)
			node.IsDir = tokens[0] == "dir"
			if !node.IsDir {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					return errors.New("error converting file size string to int")
				}
				node.FileSize = int64(size)
			}

			node.Name = tokens[1]

			currentLocation.Children = append(currentLocation.Children, node)
		}
	}

	return nil
}

func handleCommand(command string, root *FileNode, current *FileNode) *FileNode {
	if command == ".." {
		return findTarget(root, current.Name, true)
	}

	return findTarget(root, command, false)
}

func findTarget(node *FileNode, target string, stepOut bool) *FileNode {
	if node.Name == target {
		return node
	}
	
	for _, child := range node.Children {
		tempNode := findTarget(child, target, stepOut)
		if tempNode != nil {
			if (stepOut) {
				return node
			}
			return tempNode
		}
	}

	return nil
}

func calcDirSizes(node *FileNode) int64 {
	if node.Name == "jqrdvnbp" {
		fmt.Println("asdf")
	}

	if !node.IsDir {
		return node.FileSize
	}

	var dirSize int64
	dirSize = 0

	for _, child := range node.Children {
		childSize := calcDirSizes(child)
		dirSize += childSize
	}

	node.FileSize = dirSize
	return dirSize
}


func findDirs(node *FileNode, dirs []int64) []int64 {
	if node.IsDir && node.FileSize <= 100000 {
		dirs = append(dirs, node.FileSize)
	}

	for _, child := range node.Children {
		dirs = findDirs(child, dirs)
	}

	return dirs
}