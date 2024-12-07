package day07_test

import (
	"advent2024/day07"
	"testing"
)

func TestBridgeRepair(t *testing.T) {
	t.Run("simple equation", func(t *testing.T) {
		input := map[int64][]int64{
			2: {1, 1},
		}
		result := day07.SumSolvableEquations(input)
		if result != 2 {
			t.Errorf("expected 2, but got %v", result)
		}
	})

	t.Run("example", func(t *testing.T) {
		input := map[int64][]int64{
			190:    {10, 19},
			3267:   {81, 40, 27},
			83:     {17, 5},
			156:    {15, 6},
			7290:   {6, 8, 6, 15},
			161011: {16, 10, 13},
			192:    {17, 8, 14},
			21037:  {9, 7, 18, 13},
			292:    {11, 6, 16, 20},
		}
		result := day07.SumSolvableEquations(input)
		if result != 3749 {
			t.Errorf("expected 3749, but got %v", result)
		}
	})
}
