package day03

import (
	"advent2024/fileio"
	"fmt"
	"strconv"
	"strings"
)

func RunDay() {
	input, err := ReadCorruptedMemoryInputs()
	if err != nil {
		fmt.Printf("Day03 - [ERROR]: %v\n", err)
		return
	}
	part1 := Calculate(input)
	part2 := CalculatePartTwo(input)
	fmt.Printf("Day03: %v, %v \n", part1, part2)
}

func ReadCorruptedMemoryInputs() (input string, err error) {
	lines, err := fileio.ReadAllLines("./day03/input.txt")
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for _, l := range lines {
		sb.WriteString(l)
	}
	return sb.String(), nil
}

func Calculate(input string) int64 {
	muls := ParseInput(input)
	var total int64 = 0
	for _, v := range muls {
		total += v.X * v.Y
	}
	return total
}

func CalculatePartTwo(input string) int64 {
	muls := ParseInputDoDont(input)
	var total int64 = 0
	for _, v := range muls {
		total += v.X * v.Y
	}
	return total
}

func ParseInput(input string) []Mul {
	var muls = make([]Mul, 0)
	for i := 0; i < len(input); i++ {
		if input[i] == 'm' {
			mul, skip := ReadMul(input[i:])
			muls = append(muls, mul)
			i += skip - 1
			continue
		}
	}
	return muls
}

func ParseInputDoDont(input string) []Mul {
	var muls = make([]Mul, 0)
	process := true
	for i := 0; i < len(input); i++ {
		skip := 0
		if input[i] == 'd' {
			newState, s := ReadDoOrDont(input[i:], process)
			skip = s
			process = newState
		} else if input[i] == 'm' && process {
			mul, s := ReadMul(input[i:])
			skip = s
			muls = append(muls, mul)
		}
		if skip > 0 {
			i += skip - 1
		}
	}
	return muls
}

type Mul struct {
	X int64
	Y int64
}

func ReadMul(input string) (mul Mul, charactersRead int) {
	if len(input) < 8 {
		return Mul{0, 0}, len(input)
	}
	m := false
	u := false
	l := false
	open := false
	separator := false
	var number1 int64 = 0
	var number2 int64 = 0
	var chars = make([]rune, 0)
	for i := 0; i < len(input); i++ {
		if i > 0 && !m || m && input[i] == 'm' {
			return Mul{0, 0}, i
		}
		if i > 1 && !u || u && input[i] == 'u' {
			return Mul{0, 0}, i
		}
		if i > 2 && !l || l && input[i] == 'l' {
			return Mul{0, 0}, i
		}
		if i > 3 && !open || open && input[i] == '(' {
			return Mul{0, 0}, i
		}
		if separator && input[i] == ',' {
			return Mul{0, 0}, i
		}
		if input[i] == 'm' {
			m = true
			continue
		}
		if input[i] == 'u' {
			u = true
			continue
		}
		if input[i] == 'l' {
			l = true
			continue
		}
		if input[i] == '(' {
			open = true
			continue
		}
		if input[i] == ',' {
			separator = true
			n, err := strconv.ParseInt(string(chars), 10, 64)
			chars = make([]rune, 0)
			if err != nil {
				return Mul{0, 0}, i
			}
			number1 = n
			continue
		}
		if input[i] == ')' {
			n, err := strconv.ParseInt(string(chars), 10, 64)
			if err != nil {
				return Mul{0, 0}, i
			}
			number2 = n
			return Mul{number1, number2}, i
		}
		chars = append(chars, rune(input[i]))
	}
	return Mul{0, 0}, len(input)
}

func ReadDoOrDont(input string, currentState bool) (bool, int) {
	expectedDoChars := map[int]rune{
		0: 'd',
		1: 'o',
		2: '(',
		3: ')',
	}
	do, charactersReadForDo := ReadToken(input, expectedDoChars)
	if do {
		return true, charactersReadForDo
	}

	expectedDontChars := map[int]rune{
		0: 'd',
		1: 'o',
		2: 'n',
		3: '\'',
		4: 't',
		5: '(',
		6: ')',
	}
	dont, charactersReadForDont := ReadToken(input, expectedDontChars)
	if dont {
		return false, charactersReadForDont
	}
	return currentState, charactersReadForDont
}

func ReadToken(input string, expectedChars map[int]rune) (found bool, charactersRead int) {
	if len(input) < len(expectedChars) {
		return false, len(input)
	}
	tokenLocation := 0
	for i := 0; i < len(input); i++ {
		if rune(input[i]) != expectedChars[tokenLocation] {
			return false, i
		}

		tokenLocation += 1
		if tokenLocation >= len(expectedChars) {
			return true, i
		}
	}
	return false, len(input)
}
