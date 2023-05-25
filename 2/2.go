package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	calc(1)
	fmt.Println("Part 2:")
	calc(2)
}

func calc(part int) {
	total := 0
	file, err := os.Open("2input.txt")
	if err != nil {
		fmt.Println("There was an error opening the strategy file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, " ")
		var val int
		var err error
		if part == 1 {
			val, err = calcScoreForRound(letters[0], letters[1])
		} else {
			val, err = calcScoreForRoundPartTwo(letters[0], letters[1])
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		total = total + val
	}

	fmt.Println(total)
}

func calcScoreForRound(opp string, my string) (int, error) {
	switch opp {
	case "A":
		switch my {
		case "X":
			return 4, nil
		case "Y":
			return 8, nil
		case "Z":
			return 3, nil
		}
	case "B":
		switch my {
		case "X":
			return 1, nil
		case "Y":
			return 5, nil
		case "Z":
			return 9, nil
		}
	case "C":
		switch my {
		case "X":
			return 7, nil
		case "Y":
			return 2, nil
		case "Z":
			return 6, nil
		}
	}

	fmt.Println("There was an error in evaluating the score for the round.")
	return -1, errors.New("there was an error in evaluating the score for the round")
}

func calcScoreForRoundPartTwo(opp string, my string) (int, error) {
	switch opp {
	case "A":
		switch my {
		case "X":
			return 3, nil
		case "Y":
			return 4, nil
		case "Z":
			return 8, nil
		}
	case "B":
		switch my {
		case "X":
			return 1, nil
		case "Y":
			return 5, nil
		case "Z":
			return 9, nil
		}
	case "C":
		switch my {
		case "X":
			return 2, nil
		case "Y":
			return 6, nil
		case "Z":
			return 7, nil
		}
	}

	fmt.Println("There was an error in evaluating the score for the round.")
	return -1, errors.New("there was an error in evaluating the score for the round")
}
