package day03

import (
	"advent2024/fileio"
	"fmt"
	"strconv"
	"strings"
)

func RunDay(name string) {
	input, err := ReadCorruptedMemoryInputs()
	if err != nil {
		fmt.Printf("%v - [ERROR]: %v\n", name, err)
		return
	}
	part1 := Calculate(input)
	part2 := CalculatePartTwo(input)
	fmt.Printf("%v: %v, %v \n", name, part1, part2)
}

func ReadCorruptedMemoryInputs() (input string, err error) {
	var sb strings.Builder
	err = fileio.ParseAllLines("./day03/input.txt", func(line string) {
		sb.WriteString(line)
	})
	if err != nil {
		return "", err
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
			if skip > 0 {
				i += skip - 1
			}
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
	expectedMulChars := map[int]rune{
		0: 'm',
		1: 'u',
		2: 'l',
		3: '(',
		4: ',',
		5: ')',
	}
	found, values, charactersReadForMul := ReadToken(input, expectedMulChars, true)
	if found {
		return Mul{values[0], values[1]}, charactersReadForMul
	}
	return Mul{0, 0}, charactersReadForMul
}

func ReadDoOrDont(input string, currentState bool) (bool, int) {
	expectedDoChars := map[int]rune{
		0: 'd',
		1: 'o',
		2: '(',
		3: ')',
	}
	do, _, charactersReadForDo := ReadToken(input, expectedDoChars, false)
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
	dont, _, charactersReadForDont := ReadToken(input, expectedDontChars, false)
	if dont {
		return false, charactersReadForDont
	}
	return currentState, charactersReadForDont
}

func ReadToken(input string, expectedChars map[int]rune, captureParams bool) (found bool, values []int64, charactersRead int) {
	if len(input) < len(expectedChars) {
		return false, nil, len(input)
	}
	tokenLocation := 0
	for i := 0; i < len(input); i++ {
		if rune(input[i]) != expectedChars[tokenLocation] {
			return false, nil, i
		}
		if captureParams {
			valid, params, skip := CaptureParams(input[i:])
			if valid {
				values = params
				tokenLocation += 1
			}
			if skip > 0 {
				i += skip - 1
			}
		}

		tokenLocation += 1
		if tokenLocation >= len(expectedChars) {
			return true, values, i
		}
	}
	return false, nil, len(input)
}

func CaptureParams(input string) (validParams bool, values []int64, charsRead int) {
	if input[0] != '(' {
		return false, nil, 0
	}

	var capturedValues = make([]rune, 0)
	for i := 1; i < len(input); i++ {
		if input[i] == ',' {
			num, _ := strconv.ParseInt(string(capturedValues), 10, 64)
			values = append(values, num)
			capturedValues = make([]rune, 0)
			continue
		} else if input[i] == ')' {
			num, _ := strconv.ParseInt(string(capturedValues), 10, 64)
			values = append(values, num)
			if len(values) == 2 {
				return true, values, i
			} else {
				return false, nil, i
			}
		}

		_, err := strconv.ParseInt(string(input[i]), 10, 64)
		if err != nil {
			return false, nil, i
		} else {
			capturedValues = append(capturedValues, rune(input[i]))
		}
	}
	return false, nil, len(input)
}
