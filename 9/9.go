package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("9input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hx, hy, tx, ty := 0, 0, 0, 0
	tailMap := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		direction := tokens[0]
		spaces, err := strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Println("Error converting from string to num: spaces: ", err)
			return
		}

		for i := 0; i < spaces; i++ {
			// move head
			switch direction {
			case "R":
				hx++
			case "L":
				hx--
			case "U":
				hy++
			case "D":
				hy--
			}

			// check if tail needs to move
			move := needMoveTail(hx, hy, tx, ty)
			if move {
				tx, ty = moveTail(hx, hy, tx, ty)
			}

			// update map
			str := strconv.Itoa(tx) + "," + strconv.Itoa(ty)
			tailMap[str] = tx + ty
		}
	}

	fmt.Println("Advent of Code Day 9 Part 1 Result: ", len(tailMap))
}

func needMoveTail(hx int, hy int, tx int, ty int) bool {
	return math.Abs(float64(hx-tx)) > 1 || math.Abs(float64(hy-ty)) > 1
}

// returns the new x and y position of the tail
func moveTail(hx int, hy int, tx int, ty int) (int, int) {
	if hx == tx { // move tail on y
		up := (hy-ty > 0)
		if up {
			return tx, ty + (int(math.Abs(float64(hy-ty))) - 1)
		}

		return tx, ty - (int(math.Abs(float64(hy-ty))) - 1)
	}

	if hy == ty { // move tail on x
		right := (hx - tx) > 0
		if right {
			return tx + int(math.Abs(float64(hx-tx))-1), ty
		}

		return tx - int(math.Abs(float64(hx-tx))-1), ty
	}

	// need to move diagonal
	a, b := tx, ty

	if hx > tx {
		a++
	} else {
		a--
	}

	if hy > ty {
		b++
	} else {
		b--
	}

	return a, b
}
