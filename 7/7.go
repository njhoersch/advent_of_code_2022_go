package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type FileNode struct {
	Name     string
	IsDir    bool
	FileSize int
	Children []*FileNode
}

func PrintFileTree(node *FileNode, indent string) {
	fmt.Println(indent + node.Name)
	for _, child := range node.Children {
		childIndent := indent + "  "
		PrintFileTree(child, childIndent)
	}
}

func main() {
	root := new(FileNode)
	err := buildFileTree(root)

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
		if line[0] == byte("$") {

		}
	}



	return nil
}

