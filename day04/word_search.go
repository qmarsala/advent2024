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
	part2 := 0
	fmt.Printf("Day04: %v, %v \n", part1, part2)
}

func ReadWordSearchInputs() (input [][]rune, err error) {
	lines, err := fileio.ReadAllLines("./day04/input.txt")
	if err != nil {
		return make([][]rune, 0), nil
	}
	for _, l := range lines {
		input = append(input, []rune(l))
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
			if strings.ToLower(string(input[row][column])) == "x" || strings.ToLower(string(input[row][column])) == "s" {
				if ReadXMAS(input, row, column, horizontal) ||
					ReadXMAS(input, row, column, vertical) ||
					ReadXMAS(input, row, column, ascending) ||
					ReadXMAS(input, row, column, descending) {
					xmasCount += 1
				}
			}
		}
	}
	return xmasCount
}

func ReadXMAS(input [][]rune, startRow int, startColumn int, direction int) bool {
	word := ""
	for i := 0; i < 4; i++ {
		switch direction {
		case horizontal:
			if startRow >= len(input) || startColumn+i >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow][startColumn+i])
		case vertical:
			if startRow+i >= len(input) || startColumn >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow+i][startColumn])
		case ascending:
			if startRow+i >= len(input) || startColumn+i >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow+i][startColumn+i])
		case descending:
			if startRow+i >= len(input) || startColumn-i < 0 {
				return false
			}
			word += string(input[startRow+i][startColumn-i])
		default:
			return false
		}
	}
	lowerWord := strings.ToLower(string(word))
	return lowerWord == "xmas" || lowerWord == "samx"
}
