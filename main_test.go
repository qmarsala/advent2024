package main

import (
	"testing"
)

func TestHistorianHysteriaDifference(t *testing.T) {
	t.Run("1 2 3, 3 2 1 = 0", func(t *testing.T) {
		listA := []int64{1, 2, 3}
		listB := []int64{3, 2, 1}
		result := HistorianHysteriaDifference(listA, listB)

		if result != 0 {
			t.Error("Expected 0, got ", result)
		}
	})

	t.Run("2 2 3, 3 2 1 = 1", func(t *testing.T) {
		listA := []int64{2, 2, 3}
		listB := []int64{3, 2, 1}
		result := HistorianHysteriaDifference(listA, listB)

		if result != 1 {
			t.Error("Expected 1, got ", result)
		}
	})

	t.Run("3 2 2, 1 1 1 = 4", func(t *testing.T) {
		listA := []int64{3, 2, 2}
		listB := []int64{1, 1, 1}
		result := HistorianHysteriaDifference(listA, listB)

		if result != 4 {
			t.Error("Expected 4, got ", result)
		}
	})
}

func TestHistorianHysteriaSimilarity(t *testing.T) {
	t.Run("similarity example = 31", func(t *testing.T) {
		listA := []int64{3, 4, 2, 1, 3, 3}
		listB := []int64{4, 3, 5, 3, 9, 3}
		result := HistorianHysteriaSimilarity(listA, listB)

		if result != 31 {
			t.Error("Expected 31, got ", result)
		}
	})
}
