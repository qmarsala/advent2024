package day04

import "fmt"

func RunDay() {
	_, err := ReadWordSearchInputs()
	if err != nil {
		fmt.Printf("Day04 - [ERROR]: %v\n", err)
		return
	}
	part1 := 0
	part2 := 0
	fmt.Printf("Day04: %v, %v \n", part1, part2)
}

func ReadWordSearchInputs() (input [][]rune, err error) {
	return make([][]rune, 0), nil
}
