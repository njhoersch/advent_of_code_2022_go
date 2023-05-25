package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1: ")
	partOne()
}

func partOne() {
	stacks, err := readInStartingCrates()
	if err != nil {
		fmt.Println("There was an error reading in the starting crates.")
		return
	}

	for i := 0; i < 9; i++ {
		fmt.Println("___________________\nStack: ", i)
		for j := 0; j < len(stacks[i]); j++ {
			fmt.Println("Value: ", stacks[i][j])
		}
	}
}

// interface returned is [9]Stack
func readInStartingCrates() ([9][]string, error) {
	stacks := [9][]string{}

	file, err := os.Open("5input.txt")
	if err != nil {
		fmt.Println("Error opening file.")
		return stacks, errors.New("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == byte(' ') {
			return stacks, nil
		}

		stackIndex := 0
		for i := 1; i < len(line); i = i + 4 {
			if line[i] != byte(' ') {
				stacks[stackIndex] = append(stacks[stackIndex], string(line[i]))
			}
			stackIndex++
		}

		// for i := 0; i < 9; i++ {
		// 	stacks[i] = reverseSlice(stacks[i])
		// }

	}

	return stacks, nil

}

func reverseSlice(slice []string) []string {
	newStack := []string{}
	for i := len(slice) - 1; i >= 0; i-- {
		newStack = append(newStack, slice[i])
	}
	return newStack
}
