package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1: ")
	partOne()
	fmt.Println("---------------------------------------\nPart 2: ")
	partTwo()
}

func partOne() {
	file, err := os.Open("3input.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}
	defer file.Close()

	priorities := []byte {}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) /2
		first := line[:mid]
		second := line[mid:]

		outerloop:
		for i := 0; i < len(first); i++ {
			for j := 0; j < len(second); j++ {
				if first[i] == second[j] {
					priorities = append(priorities, first[i])
					break outerloop
				}
			}
		}
	}

	total := 0
	for i := 0; i < len(priorities); i++ {
		val := priorities[i]
		if val > byte(92) {
			total = total + int(val - 96)
		} else {
			total = total + int(val - 38)
		}
	}

	fmt.Println("Total: ", total)
}

func partTwo() {
	file, err := os.Open("3input.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}
	defer file.Close()

	priorities := []byte {}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		outerloop:
		for i := 0; i < len(first); i++ {
			for j := 0; j < len(second); j++ {
				for k := 0; k < len(third); k++ {
					if first[i] == second[j] && first[i] == third[k] {
						priorities = append(priorities, first[i])
						break outerloop
					}
				}
			}
		}
	}

	total := 0
	for i := 0; i < len(priorities); i++ {
		val := priorities[i]
		if val > byte(92) {
			total = total + int(val - 96)
		} else {
			total = total + int(val - 38)
		}
	}

	fmt.Println("Total: ", total)
}