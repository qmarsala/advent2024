package day04

import (
	"advent2024/fileio"
	"fmt"
	"strings"
)

func RunDay(name string) {
	input, err := ReadWordSearchInputs()
	if err != nil {
		fmt.Printf("%v - [ERROR]: %v\n", name, err)
		return
	}
	part1 := CountXMAS(input)
	part2 := CountXMAS2(input)
	fmt.Printf("%v: %v, %v \n", name, part1, part2)
}

func ReadWordSearchInputs() (input [][]rune, err error) {
	err = fileio.ParseAllLines("./day04/input.txt", func(line string) {
		input = append(input, []rune(strings.ToLower(line)))
	})
	if err != nil {
		return make([][]rune, 0), err
	}
	return input, nil
}

const (
	horizontal = iota
	vertical
	ascending
	descending
)

func CountXMAS(input [][]rune) int {
	xmasCount := 0
	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			currentChar := strings.ToLower(string(input[row][column]))
			if currentChar == "x" || currentChar == "s" {
				xmasCount += ReadXMAS(input, row, column, horizontal) +
					ReadXMAS(input, row, column, vertical) +
					ReadXMAS(input, row, column, ascending) +
					ReadXMAS(input, row, column, descending)
			}
		}
	}
	return xmasCount
}

func CountXMAS2(input [][]rune) int {
	xmasCount := 0
	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			if strings.ToLower(string(input[row][column])) == "a" {
				leftScore := ScoreXMAS(input, row, column, ascending)
				rightScore := ScoreXMAS(input, row, column, descending)
				if leftScore == 4 && rightScore == 4 {
					xmasCount += 1
				}
			}
		}
	}
	return xmasCount
}

func ReadXMAS(input [][]rune, startRow int, startColumn int, direction int) int {
	word := ""
	for i := 0; i < 4; i++ {
		switch direction {
		case horizontal:
			if startRow >= len(input) || startColumn+i >= len(input[startRow]) {
				return 0
			}
			word += string(input[startRow][startColumn+i])
		case vertical:
			if startRow+i >= len(input) || startColumn >= len(input[startRow]) {
				return 0
			}
			word += string(input[startRow+i][startColumn])
		case ascending:
			if startRow+i >= len(input) || startColumn+i >= len(input[startRow]) {
				return 0
			}
			word += string(input[startRow+i][startColumn+i])
		case descending:
			if startRow+i >= len(input) || startColumn-i < 0 {
				return 0
			}
			word += string(input[startRow+i][startColumn-i])
		default:
			return 0
		}
	}
	if strings.ToLower(word) == "xmas" || strings.ToLower(word) == "samx" {
		return 1
	} else {
		return 0
	}
}

func ScoreXMAS(input [][]rune, startRow int, startColumn int, direction int) int {
	scores := map[rune]int{
		'm': 1,
		's': 3,
	}
	score := 0
	switch direction {
	case ascending:
		if startRow-1 < 0 || startColumn+1 >= len(input[startRow]) {
			return 0
		}
		if v, ok := scores[input[startRow-1][startColumn+1]]; ok {
			score += v
		}
		if startRow+1 >= len(input) || startColumn-1 < 0 {
			return 0
		}
		if v, ok := scores[input[startRow+1][startColumn-1]]; ok {
			score += v
		}
	case descending:
		if startRow-1 < 0 || startColumn-1 < 0 {
			return 0
		}
		if v, ok := scores[input[startRow-1][startColumn-1]]; ok {
			score += v
		}
		if startRow+1 >= len(input) || startColumn+1 >= len(input[startRow]) {
			return 0
		}
		if v, ok := scores[input[startRow+1][startColumn+1]]; ok {
			score += v
		}
	default:
		return 0
	}
	return score
}
