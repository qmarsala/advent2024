package day07

import (
	"advent2024/fileio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func RunDay(name string) {
	input, err := ReadBridgeRepairInputs()
	if err != nil {
		fmt.Printf("%v - [ERROR]: %v\n", name, err)
		return
	}
	part1 := SumSolvableEquations(input)
	part2 := 0
	fmt.Printf("%v: %v, %v \n", name, part1, part2)
}

func ReadBridgeRepairInputs() (input map[int64][]int64, err error) {
	input = map[int64][]int64{}
	err = fileio.ParseAllLines("./day07/input.txt", func(line string) {
		parts := strings.Split(line, ":")
		var key int64 = 0
		if num, err := strconv.ParseInt(parts[0], 10, 64); err == nil {
			key = num
		} else {
			return
		}
		numbers := strings.Split(parts[1], " ")
		for _, n := range numbers {
			if num, err := strconv.ParseInt(n, 10, 64); err == nil {
				input[key] = append(input[key], num)
			}
		}
	})
	return input, err
}

func SumSolvableEquations(input map[int64][]int64) (sum int64) {
	for k, v := range input {
		all := slices.Concat(
			GenerateCombinations([]rune{'+', '*'}, len(v)-1),
			GenerateCombinations([]rune{'*', '+'}, len(v)-1))
		for _, c := range all {
			if len(c) != len(v)-1 {
				continue
			}
			if ProcessOperators(v, c) == k {
				sum += k
				break
			}
		}
	}
	return
}

func GenerateCombinations(alphabet []rune, length int) [][]rune {
	combinations := AddOperator([][]rune{}, []rune{}, alphabet, length)
	return combinations
}

func AddOperator(combinations [][]rune, combo []rune, alphabet []rune, length int) [][]rune {
	if length <= 0 {
		return combinations
	}

	var newCombo = []rune{}
	for _, ch := range alphabet {
		newCombo = append(combo, ch)
		combinations = append(AddOperator(combinations, newCombo, alphabet, length-1), newCombo)
	}
	return combinations
}

func ProcessOperators(numbers []int64, operators []rune) (result int64) {
	result = numbers[0]
	for i, o := range operators {
		switch o {
		case '+':
			result += numbers[i+1]
		case '*':
			result *= numbers[i+1]
		}
	}
	return
}
