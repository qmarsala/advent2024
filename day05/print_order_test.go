package day05_test

import (
	"advent2024/day05"
	"testing"
)

func TestPrintOrder(t *testing.T) {
	t.Run("Finds Middle", func(t *testing.T) {
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
		pages := [][]int64{{1, 2, 3}}
		result := day05.ScoreValidUpdates(orders, pages)
		if result != 2 {
			t.Errorf("expected 2, but got %v", result)
		}
	})

	t.Run("Scores Middle", func(t *testing.T) {
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
		pages := [][]int64{{1, 2, 3}, {1, 2, 3}}
		result := day05.ScoreValidUpdates(orders, pages)
		if result != 4 {
			t.Errorf("expected 4, but got %v", result)
		}
	})

	t.Run("Scores Middle ignoring invalid", func(t *testing.T) {
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
		pages := [][]int64{{1, 2, 3}, {1, 3, 2}}
		result := day05.ScoreValidUpdates(orders, pages)
		if result != 2 {
			t.Errorf("expected 2, but got %v", result)
		}
	})

	t.Run("Scores Middle When not valid", func(t *testing.T) {
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
				Before: 5,
			},
			{
				Value:  4,
				Before: 3,
			},
		}
		pages := [][]int64{{1, 2, 3}, {1, 3, 4}}
		result := day05.ScoreInValidUpdates(orders, pages)
		if result != 4 {
			t.Errorf("expected 4, but got %v", result)
		}
	})

	t.Run("Scores Middle Example Input", func(t *testing.T) {
		orders := []day05.Order{
			{
				Value:  47,
				Before: 53,
			},
			{
				Value:  97,
				Before: 13,
			},
			{
				Value:  97,
				Before: 61,
			},
			{
				Value:  97,
				Before: 47,
			},
			{
				Value:  75,
				Before: 29,
			},
			{
				Value:  61,
				Before: 13,
			},
			{
				Value:  75,
				Before: 53,
			},
			{
				Value:  29,
				Before: 13,
			},
			{
				Value:  97,
				Before: 29,
			},
			{
				Value:  53,
				Before: 29,
			},
			{
				Value:  61,
				Before: 53,
			},
			{
				Value:  97,
				Before: 53,
			},
			{
				Value:  61,
				Before: 29,
			},
			{
				Value:  47,
				Before: 13,
			},
			{
				Value:  75,
				Before: 47,
			},
			{
				Value:  97,
				Before: 75,
			},
			{
				Value:  47,
				Before: 61,
			},
			{
				Value:  75,
				Before: 61,
			},
			{
				Value:  47,
				Before: 29,
			},
			{
				Value:  75,
				Before: 13,
			},
			{
				Value:  53,
				Before: 13,
			},
		}
		pages := [][]int64{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}
		result := day05.ScoreValidUpdates(orders, pages)
		if result != 143 {
			t.Errorf("expected 143, but got %v", result)
		}
	})

	t.Run("Scores Middle Example Input when invalid", func(t *testing.T) {
		orders := []day05.Order{
			{
				Value:  47,
				Before: 53,
			},
			{
				Value:  97,
				Before: 13,
			},
			{
				Value:  97,
				Before: 61,
			},
			{
				Value:  97,
				Before: 47,
			},
			{
				Value:  75,
				Before: 29,
			},
			{
				Value:  61,
				Before: 13,
			},
			{
				Value:  75,
				Before: 53,
			},
			{
				Value:  29,
				Before: 13,
			},
			{
				Value:  97,
				Before: 29,
			},
			{
				Value:  53,
				Before: 29,
			},
			{
				Value:  61,
				Before: 53,
			},
			{
				Value:  97,
				Before: 53,
			},
			{
				Value:  61,
				Before: 29,
			},
			{
				Value:  47,
				Before: 13,
			},
			{
				Value:  75,
				Before: 47,
			},
			{
				Value:  97,
				Before: 75,
			},
			{
				Value:  47,
				Before: 61,
			},
			{
				Value:  75,
				Before: 61,
			},
			{
				Value:  47,
				Before: 29,
			},
			{
				Value:  75,
				Before: 13,
			},
			{
				Value:  53,
				Before: 13,
			},
		}
		pages := [][]int64{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}
		result := day05.ScoreInValidUpdates(orders, pages)
		if result != 123 {
			t.Errorf("expected 123, but got %v", result)
		}
	})

}
