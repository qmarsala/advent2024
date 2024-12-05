package day05_test

import (
	"advent2024/day05"
	"testing"
)

func TestPrintOrder(t *testing.T) {
	t.Run("ScoreFindsMiddle", func(t *testing.T) {
		orders := []day05.Order{
			{
				Value:  1,
				Before: 2,
			},
			{
				Value:  2,
				Before: 3,
			},
			{
				Value:  3,
				Before: 4,
			},
		}
		pages := []int64{1, 2, 3}
		result := day05.ScoreUpdates(orders, pages)
		if result != 2 {
			t.Errorf("expected 2, but got %v", result)
		}
	})
}
