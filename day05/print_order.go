package day05

import (
	"advent2024/fileio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func RunDay() {
	orders, updatedPages, err := ReadPrintOrderInputs()
	if err != nil {
		fmt.Printf("Day05 - [ERROR]: %v\n", err)
		return
	}
	part1 := ScoreValidUpdates(orders, updatedPages)
	part2 := 0
	fmt.Printf("Day05: %v, %v \n", part1, part2)
}

type Order struct {
	Value  int64
	Before int64
}

func ReadPrintOrderInputs() (orders []Order, updatedPages [][]int64, err error) {
	lines, err := fileio.ReadAllLines("./day05/input.txt")
	if err != nil {
		return make([]Order, 0), make([][]int64, 0), err
	}
	readingOrder := true
	for _, l := range lines {
		if l == "" {
			readingOrder = false
			continue
		}
		if readingOrder {
			parts := strings.Split(l, "|")
			newOrder := Order{
				Value:  0,
				Before: 1,
			}
			if value, err := strconv.ParseInt(parts[0], 10, 64); err != nil {
				return nil, nil, err
			} else {
				newOrder.Value = value
			}
			if value, err := strconv.ParseInt(parts[1], 10, 64); err != nil {
				return nil, nil, err
			} else {
				newOrder.Before = value
			}
			orders = append(orders, newOrder)
		} else {
			pages := strings.Split(l, ",")
			values := []int64{}
			for _, p := range pages {
				if value, err := strconv.ParseInt(p, 10, 64); err != nil {
					return nil, nil, err
				} else {
					values = append(values, value)
				}
			}
			updatedPages = append(updatedPages, values)
		}
	}
	return orders, updatedPages, nil
}

func ScoreValidUpdates(orders []Order, updatedPages [][]int64) int {
	score := 0
	for _, p := range updatedPages {
		if !ValidateUpdate(orders, p) {
			continue
		}
		len := len(p)
		if len >= 3 {
			score += int(p[(len-1)/2])
		} else {
			score += int(p[0])
		}
	}

	return score
}

func ScoreInValidUpdates(orders []Order, updatedPages [][]int64) int {
	score := 0
	for _, p := range updatedPages {
		if !ValidateUpdate(orders, p) {
			continue
		}
		len := len(p)
		if len >= 3 {
			score += int(p[(len-1)/2])
		} else {
			score += int(p[0])
		}
	}

	return score
}

func ValidateUpdate(order []Order, updatedPage []int64) bool {
	valid := true
	for _, o := range order {
		if slices.Contains(updatedPage, o.Value) && slices.Contains(updatedPage, o.Before) {
			valid = valid && slices.Index(updatedPage, o.Value) < slices.Index(updatedPage, o.Before)
		}
	}
	return valid
}
