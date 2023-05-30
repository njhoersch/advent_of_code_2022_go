package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// DS to store trees
	trees := make([][]int, 0)

	// Open file
	file, err := os.Open("8input.txt")
	if err != nil {
		fmt.Println("File Opening Error: ", err)
		return
	}
	defer file.Close()

	// Read in trees
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, 0)

		for _, char := range line {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println("Str conv error: ", err)
				return
			}
			
			row = append(row, val)
		}

		trees = append(trees, row)
	}

	total := 0	

	// For every tree, check if it's visible. If visible -> increment total
	for i, row := range trees {
		for j := range row {
			if isTreeVisable(&trees, i, j) {
				total++
			}
		}
	}

	fmt.Println("Advent of code day 8 part 1 Result: ", total)

	highest := 0

	for i, row := range trees {
		for j := range row {
			score := calculateScenicScore(&trees, i, j)
			if score > highest {
				highest = score
			}
		}
	}

	fmt.Println("Advent of code day 8 part 2 Result: ", highest)
}

func isTreeVisable(trees *[][]int, row int, col int) bool {
	tSlice := *trees

	height := len(tSlice)
	length := len(tSlice[0])

	if row == 0 || row == length {
		return true
	}

	if col == 0 || col == height {
		return true
	}

	left, right, up, down := false, false, false, false

	// LEFT
	for i := col - 1; i >= 0; i-- {
		if tSlice[row][i] >= tSlice[row][col] {
			left = true
		}
	}

	// RIGHT
	for i := col + 1; i < length; i++ {
		if tSlice[row][i] >= tSlice[row][col] {
			right = true
		}
	}

	// UP
	for i := row - 1; i >= 0; i-- {
		if tSlice[i][col] >= tSlice[row][col] {
			up = true
		}
	}

	// DOWN
	for i := row + 1; i < height; i++ {
		if tSlice[i][col] >= tSlice[row][col] {
			down = true
		}
	}

	return !(left && right && up && down)
}

func calculateScenicScore(trees *[][]int, row int, col int) int {
	tSlice := *trees

	height := len(tSlice)
	length := len(tSlice[0])

	left, right, up, down := 0, 0, 0, 0

	// LEFT
	for i := col - 1; i >= 0; i-- {
		left++
		if tSlice[row][i] >= tSlice[row][col] {
			break
		}
	}

	// RIGHT
	for i := col + 1; i < length; i++ {
		right++
		if tSlice[row][i] >= tSlice[row][col] {
			break
		}
	}

	// UP
	for i := row - 1; i >= 0; i-- {
		up++
		if tSlice[i][col] >= tSlice[row][col] {
			break
		}
	}

	// DOWN
	for i := row + 1; i < height; i++ {
		down++
		if tSlice[i][col] >= tSlice[row][col] {
			break
		}
	}

	return (left * right * up * down)
}