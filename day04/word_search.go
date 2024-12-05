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
				print(input[row])
				if ReadXMAS(input, row, column, left) ||
					ReadXMAS(input, row, column, right) ||
					ReadXMAS(input, row, column, up) ||
					ReadXMAS(input, row, column, down) ||
					ReadXMAS(input, row, column, upRight) ||
					ReadXMAS(input, row, column, upLeft) ||
					ReadXMAS(input, row, column, downRight) ||
					ReadXMAS(input, row, column, downLeft) {
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
		case left:
			if startRow >= len(input) || startColumn-i < 0 {
				return false
			}
			word += string(input[startRow][startColumn-i])
		case right:
			if startRow >= len(input) || startColumn+i >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow][startColumn+i])
		case up:
			if startRow-i < 0 || startColumn >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow-i][startColumn])
		case down:
			if startRow+i >= len(input[startRow]) || startColumn >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow+i][startColumn])
		case upRight:
			if startRow-i < 0 || startColumn+i >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow-i][startColumn+i])
		case downRight:
			if startRow+i >= len(input[startRow]) || startColumn+i >= len(input[startRow]) {
				return false
			}
			word += string(input[startRow+i][startColumn+i])
		case upLeft:
			if startRow-i < 0 || startColumn-i < 0 {
				return false
			}
			word += string(input[startRow-i][startColumn-i])
		case downLeft:
			if startRow+i >= len(input) || startColumn-i < 0 {
				return false
			}
			word += string(input[startRow+i][startColumn-i])
		default:
			return false
		}
	}
	lowerWord := strings.ToLower(string(word))
	return lowerWord == "xmas"
}
