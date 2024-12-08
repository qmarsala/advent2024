package day07

import (
	"advent2024/fileio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func RunDay(name string) {
	_, err := ReadBridgeRepairInputs()
	if err != nil {
		fmt.Printf("%v - [ERROR]: %v\n", name, err)
		return
	}
	part1 := "wip"
	part2 := "wip"
	fmt.Printf("%v: %v, %v \n", name, part1, part2)
}

type Equation struct {
	Result  int64
	Numbers []int64
}

func ReadBridgeRepairInputs() (input []Equation, err error) {
	err = fileio.ParseAllLines("./day07/input.txt", func(line string) {
		parts := strings.Split(line, ":")
		var key int64 = 0
		if num, err := strconv.ParseInt(parts[0], 10, 64); err == nil {
			key = num
		} else {
			fmt.Printf("key: %v\n", err)
			return
		}
		numberString := strings.Split(parts[1], " ")
		numbers := []int64{}
		for _, n := range numberString {
			if len(n) < 1 {
				continue
			}
			if num, err := strconv.ParseInt(n, 10, 64); err == nil {
				numbers = append(numbers, num)
			} else {
				fmt.Printf("val: %v\n", err)
				return
			}
		}
		input = append(input, Equation{
			Result:  key,
			Numbers: numbers,
		})
	})
	return input, err
}

func SumSolvableEquations(input []Equation) (sum int64) {
	for _, eq := range input {
		operationCount := len(eq.Numbers) - 1
		all := slices.Concat(GenerateCombinations([]rune{'+', '*'}, operationCount), GenerateCombinations([]rune{'*', '+'}, operationCount))
		for _, operations := range all {
			if len(operations) != operationCount {
				continue
			}
			result := ProcessOperators(eq.Numbers, operations)
			if result == eq.Result {
				sum += eq.Result
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
