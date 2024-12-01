package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func RunDay() {
	listA, listB, err := ReadHistorianHysteriaInput()
	if err != nil {
		fmt.Printf("Day01 - [ERROR]: %v\n", err)
		return
	}

	diff := HistorianHysteriaDifference(listA, listB)
	sim := HistorianHysteriaSimilarity(listA, listB)
	fmt.Printf("Day01: %v, %v \n", diff, sim)
}

func ReadHistorianHysteriaInput() (listA []int64, listB []int64, err error) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		valueA, _ := strconv.ParseInt(parts[0], 10, 64)
		valueB, _ := strconv.ParseInt(parts[1], 10, 64)
		listA = append(listA, valueA)
		listB = append(listB, valueB)
	}
	return listA, listB, nil
}

func HistorianHysteriaDifference(listA []int64, listB []int64) int64 {
	slices.Sort(listA)
	slices.Sort(listB)
	var total int64 = 0
	for i := 0; i < len(listA); i++ {
		if listA[i] < listB[i] {
			total += listB[i] - listA[i]
		} else {
			total += listA[i] - listB[i]
		}
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
