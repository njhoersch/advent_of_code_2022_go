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
	fmt.Println("----------------------------------\nPart 2: ")
	calc(2)
}

func calc(part int) {
	file, err := os.Open("4input.txt")
	if err != nil {
		fmt.Println("Error opening file.")
		return
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sections := strings.Split(line, ",")
		s1 := strings.Split(sections[0], "-")
		s2 := strings.Split(sections[1], "-")

		n1, err := strconv.Atoi(s1[0])
		if err != nil {
			fmt.Println("Err converting str to int.") 
			return
		}

		n2, err := strconv.Atoi(s1[1])
		if err != nil {
			fmt.Println("Err converting str to int.") 
			return
		}

		n3, err := strconv.Atoi(s2[0])
		if err != nil {
			fmt.Println("Err converting str to int.") 
			return
		}

		n4, err := strconv.Atoi(s2[1])
		if err != nil {
			fmt.Println("Err converting str to int.") 
			return
		}

		if (part == 1){
			if (n1 >= n3 && n1 <= n4 && n2 >= n3 && n2 <= n4) || (n3 >= n1 && n3 <= n2 && n4 >= n1 && n4 <= n2) {
				total++
			}			
		} else {
			if (n3 >= n1 && n3 <= n2) || (n4 >= n1 && n4 <= n2) || (n1 >= n3 && n1 <= n4) || (n2 >= n3 && n2 <= n4) {
				total++
			}
		}
		
	}

	fmt.Println("Total: ", total)	
}