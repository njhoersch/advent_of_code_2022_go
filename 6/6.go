package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	fmt.Println("Part 1: ")
	calc(1)
	fmt.Println("-----------------\nPart 2: ")
	calc(2)
}

func calc(part int) {
	file, err := os.Open("6input.txt")
	if err != nil {
		fmt.Println("Error opening file...")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	if part == 1 {
		for i := 0; i < len(line) - 1; i++ {
			str := line[i:i+4]
			if charsAreUnique(str) {
				fmt.Println("resut: ", i + 4)
				return
			}
		}
	} else {
		for i := 0; i < len(line) - 1; i++ {
			str := line[i:i+14]
			if charsAreUnique(str) {
				fmt.Println("resut: ", i + 14)
				return
			}
		}
	}

	fmt.Println("We were unable to find the marker.")
}

func charsAreUnique(str string) bool {
	seen := make(map[rune]bool)
	for _, char := range str {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}
