package day01_test

import (
	"advent2024/day01"
	"testing"
)

func TestHistorianHysteriaDifference(t *testing.T) {
	t.Run("difference example = 11", func(t *testing.T) {
		listA := []int64{3, 4, 2, 1, 3, 3}
		listB := []int64{4, 3, 5, 3, 9, 3}
		result := day01.HistorianHysteriaDifference(listA, listB)

		if result != 11 {
			t.Error("Expected 11 got ", result)
		}
	})
}

func TestHistorianHysteriaSimilarity(t *testing.T) {
	t.Run("similarity example = 31", func(t *testing.T) {
		listA := []int64{3, 4, 2, 1, 3, 3}
		listB := []int64{4, 3, 5, 3, 9, 3}
		result := day01.HistorianHysteriaSimilarity(listA, listB)

		if result != 31 {
			t.Error("Expected 31, got ", result)
		}
	})
}
