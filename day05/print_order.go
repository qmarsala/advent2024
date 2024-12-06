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
	part2 := ScoreInValidUpdates(orders, updatedPages)
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
		if valid, _ := ValidateUpdate(orders, p); !valid {
			continue
		}
		score += GetMiddle(p)
	}

	return score
}

func ScoreInValidUpdates(orders []Order, updatedPages [][]int64) int {
	score := 0
	for _, p := range updatedPages {
		if valid, containsOrders := ValidateUpdate(orders, p); valid || len(containsOrders) < 1 {
			continue
		} else {
			fixedOrder := SortUpdates(containsOrders, p)
			score += GetMiddle(fixedOrder)
		}
	}

	return score
}

func GetMiddle(list []int64) int {
	len := len(list)
	if len >= 3 {
		return int(list[(len-1)/2])
	} else {
		return int(list[0])
	}
}

func SortUpdates(orders []Order, page []int64) []int64 {
	newList := []int64{}
	for i, p := range page {
		index := GetBeforeIndex(orders, newList, p, i)
		newList = slices.Insert(newList, index, p)
	}
	if v, _ := ValidateUpdate(orders, newList); v {
		return newList
	} else {
		return SortUpdates(orders, newList)
	}
}

func GetBeforeIndex(orders []Order, list []int64, value int64, defaultIndex int) int {
	for _, o := range orders {
		if o.Value != value {
			continue
		}

		index := slices.Index(list, o.Before)
		if index > -1 {
			return index
		}
	}
	return defaultIndex
}

func ValidateUpdate(order []Order, updatedPage []int64) (valid bool, containsOrders []Order) {
	valid = true
	containsOrders = []Order{}
	for _, o := range order {
		if slices.Contains(updatedPage, o.Value) && slices.Contains(updatedPage, o.Before) {
			containsOrders = append(containsOrders, o)
			valid = valid && slices.Index(updatedPage, o.Value) < slices.Index(updatedPage, o.Before)
		}
	}
	return valid, containsOrders
}
