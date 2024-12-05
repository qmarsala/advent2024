package day04

import (
	"advent2024/fileio"
	"fmt"
	"strings"
)

func RunDay() {
	input, err := ReadWordSearchInputs()
	if err != nil {
		fmt.Printf("Day04 - [ERROR]: %v\n", err)
		return
	}
	part1 := CountXMAS(input)
	part2 := CountXMAS2(input)
	fmt.Printf("Day04: %v, %v \n", part1, part2)
}

func ReadWordSearchInputs() (input [][]rune, err error) {
	lines, err := fileio.ReadAllLines("./day04/input.txt")
	if err != nil {
		return make([][]rune, 0), nil
	}
	for _, l := range lines {
		input = append(input, []rune(strings.ToLower(l)))
	}
	return input, nil
}

const (
	left = iota
	right
	up
	down
	upRight
	downRight
	upLeft
	downLeft
)

func CountXMAS(input [][]rune) int {
	xmasCount := 0
	for row := 0; row < len(input); row++ {
		for column := 0; column < len(input[row]); column++ {
			if strings.ToLower(string(input[row][column])) == "x" {
				xmasCount += ReadXMAS(input, row, column, left) +
					ReadXMAS(input, row, column, right) +
					ReadXMAS(input, row, column, up) +
					ReadXMAS(input, row, column, down) +
					ReadXMAS(input, row, column, upRight) +
					ReadXMAS(input, row, column, upLeft) +
					ReadXMAS(input, row, column, downRight) +
					ReadXMAS(input, row, column, downLeft)
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
				leftScore := ScoreXMAS(input, row, column, upRight) +
					ScoreXMAS(input, row, column, downLeft)
				rightScore := ScoreXMAS(input, row, column, downRight) +
					ScoreXMAS(input, row, column, upLeft)
				if leftScore == 4 && rightScore == 4 {
					xmasCount += 1
				}
			}
		}
	}
	return xmasCount
}

func ReadXMAS(input [][]rune, startRow int, startColumn int, direction int) int {
	currentMatch := 0
	expectedXmasChars := map[int]rune{
		0: 'x',
		1: 'm',
		2: 'a',
		3: 's',
	}
	for i := 0; i < 4; i++ {
		switch direction {
		case left:
			if startRow >= len(input) || startColumn-i < 0 {
				return 0
			}
			if input[startRow][startColumn-i] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case right:
			if startRow >= len(input) || startColumn+i >= len(input[startRow]) {
				return 0
			}
			if input[startRow][startColumn+i] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case up:
			if startRow-i < 0 || startColumn >= len(input[startRow]) {
				return 0
			}
			if input[startRow-i][startColumn] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case down:
			if startRow+i >= len(input[startRow]) || startColumn >= len(input[startRow]) {
				return 0
			}
			if input[startRow+i][startColumn] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case upRight:
			if startRow-i < 0 || startColumn+i >= len(input[startRow]) {
				return 0
			}
			if input[startRow-i][startColumn+i] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case downRight:
			if startRow+i >= len(input[startRow]) || startColumn+i >= len(input[startRow]) {
				return 0
			}
			if input[startRow+i][startColumn+i] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case upLeft:
			if startRow-i < 0 || startColumn-i < 0 {
				return 0
			}
			if input[startRow-i][startColumn-i] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		case downLeft:
			if startRow+i >= len(input) || startColumn-i < 0 {
				return 0
			}
			if input[startRow+i][startColumn-i] == expectedXmasChars[currentMatch] {
				currentMatch += 1
			}
		default:
			return 0
		}
	}
	if currentMatch == len(expectedXmasChars) {
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

	switch direction {
	case upRight:
		if startRow-1 < 0 || startColumn+1 >= len(input[startRow]) {
			return 0
		}
		if v, ok := scores[input[startRow-1][startColumn+1]]; ok {
			return v
		}
	case downRight:
		if startRow+1 >= len(input[startRow]) || startColumn+1 >= len(input[startRow]) {
			return 0
		}
		if v, ok := scores[input[startRow+1][startColumn+1]]; ok {
			return v
		}
	case upLeft:
		if startRow-1 < 0 || startColumn-1 < 0 {
			return 0
		}
		if v, ok := scores[input[startRow-1][startColumn-1]]; ok {
			return v
		}
	case downLeft:
		if startRow+1 >= len(input) || startColumn-1 < 0 {
			return 0
		}
		if v, ok := scores[input[startRow+1][startColumn-1]]; ok {
			return v
		}
	default:
		return 0
	}
	return 0
}
