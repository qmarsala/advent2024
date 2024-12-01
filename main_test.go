package main

import (
	"testing"
)

func TestHistorianHysteria(t *testing.T) {
	t.Run("1 2 3, 3 2 1 = 0", func(t *testing.T) {
		listA := []int64{1, 2, 3}
		listB := []int64{3, 2, 1}
		result := HistorianHysteria(listA, listB)

		if result != 0 {
			t.Error("Expected 0, got ", result)
		}
	})

	t.Run("2 2 3, 3 2 1 = 1", func(t *testing.T) {
		listA := []int64{2, 2, 3}
		listB := []int64{3, 2, 1}
		result := HistorianHysteria(listA, listB)

		if result != 1 {
			t.Error("Expected 1, got ", result)
		}
	})

	t.Run("3 2 2, 1 1 1 = 4", func(t *testing.T) {
		listA := []int64{2, 2, 3}
		listB := []int64{3, 2, 1}
		result := HistorianHysteria(listA, listB)

		if result != 1 {
			t.Error("Expected 4, got ", result)
		}
	})
}
