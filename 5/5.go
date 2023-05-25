package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1: ")
	calc(1)
	fmt.Println("--------------------\nPart 2: ")
	calc(2)
}

func calc(part int) {
	stacks := [9][]string{}

	file, err := os.Open("5input.txt")
	if err != nil {
		fmt.Println("Error opening file.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == byte(' ') {
			break
		}

		stackIndex := 0
		for i := 1; i < len(line); i = i + 4 {
			if line[i] != byte(' ') {
				stacks[stackIndex] = append(stacks[stackIndex], string(line[i]))
			}
			stackIndex++
		}
	}

	for index, val := range stacks {
		stacks[index] = reverseSlice(val)
	}

	// we have our stacks, now we need to process instructions
	count := []int {}
	start := []int {}
	end := []int {}

	// scan in instructions
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {continue}
		words := strings.Split(line, " ")

		c, err := strconv.Atoi(words[1])
		if err != nil {
			fmt.Println("Error converting instruction num from string to int")
			return
		}

		s, err := strconv.Atoi(words[3])
		if err != nil {
			fmt.Println("Error converting instruction num from string to int")
			return
		}

		e, err := strconv.Atoi(words[5])
		if err != nil {
			fmt.Println("Error converting instruction num from string to int")
			return
		}

		count = append(count, c)
		start = append(start, s)
		end = append(end, e)
	}

	if part == 1 {
		for index := range count {
			for i := 0; i < count[index]; i++ {
				// get a value from start and place in end
				val := stacks[start[index] -1][len(stacks[start[index]-1]) - 1]
				stacks[start[index]-1] = stacks[start[index]-1][:len(stacks[start[index]-1]) - 1]
				stacks[end[index]-1] = append(stacks[end[index]-1], val)
			}
		}
	} else {
		for index := range count {
			// get slice by cutting count off of start slice
			temp := stacks[start[index] - 1][len(stacks[start[index] -1]) - count[index]:]
			stacks[start[index] - 1] = stacks[start[index] - 1][:(len(stacks[start[index] -1]) - count[index])]
			stacks[end[index] - 1] = append(stacks[end[index]-1], temp...)
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Println("Stack: ", i + 1, " Top Value: ", stacks[i][len(stacks[i]) - 1])
	}

}

func reverseSlice(slice []string) []string {
	newSlice := []string {}

	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}

	return newSlice
}