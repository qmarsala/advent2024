package day01

import (
	"advent2024/fileio"
	"advent2024/math"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func RunDay(name string) {
	input, err := ReadHistorianHysteriaInput()
	if err != nil {
		fmt.Printf("%v - [ERROR]: %v\n", name, err)
		return
	}

	diff := HistorianHysteriaDifference(input.ListA, input.ListB)
	sim := HistorianHysteriaSimilarity(input.ListA, input.ListB)
	fmt.Printf("%v: %v, %v \n", name, diff, sim)
}

type HistorianHysteriaInput struct {
	ListA []int64
	ListB []int64
}

func ReadHistorianHysteriaInput() (input HistorianHysteriaInput, err error) {
	err = fileio.ParseAllLines("./day01/input.txt", func(line string) {
		parts := strings.Split(line, "   ")
		valueA, _ := strconv.ParseInt(parts[0], 10, 64)
		valueB, _ := strconv.ParseInt(parts[1], 10, 64)
		input.ListA = append(input.ListA, valueA)
		input.ListB = append(input.ListB, valueB)
	})
	if err != nil {
		return HistorianHysteriaInput{}, err
	}
	return input, nil
}

func HistorianHysteriaDifference(listA []int64, listB []int64) int64 {
	slices.Sort(listA)
	slices.Sort(listB)
	var total int64 = 0
	for i := 0; i < len(listA); i++ {
		total += math.GetDiff(listA[i], listB[i])
	}
	return total
}

func HistorianHysteriaSimilarity(listA []int64, listB []int64) int64 {
	listBCounts := CountUniques(listB)
	var total int64 = 0
	for _, v := range listA {
		total += v * listBCounts[v]
	}
	return total
}

func CountUniques(numbers []int64) map[int64]int64 {
	uniqueCounts := map[int64]int64{}
	for _, v := range numbers {
		uniqueCounts[v] += 1
	}
	return uniqueCounts
}
