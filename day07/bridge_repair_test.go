package day07_test

import (
	"advent2024/day07"
	"testing"
)

func TestBridgeRepair(t *testing.T) {
	t.Run("simple equation", func(t *testing.T) {
		input := []day07.Equation{
			{Result: 24372240066, Numbers: []int64{6, 1, 2, 89, 7, 2, 3, 1, 692, 4, 66}},
		}
		result := day07.SumSolvableEquations(input)
		if result != 2 {
			t.Errorf("expected 2, but got %v", result)
		}
	})

	t.Run("simple equations, three possibilities", func(t *testing.T) {
		input := []day07.Equation{
			{Result: 2, Numbers: []int64{1, 1}},
			{Result: 4, Numbers: []int64{2, 2}},
		}
		result := day07.SumSolvableEquations(input)
		if result != 6 {
			t.Errorf("expected 6, but got %v", result)
		}
	})

	t.Run("example", func(t *testing.T) {
		input := []day07.Equation{
			{Result: 190, Numbers: []int64{10, 19}},
			{Result: 3267, Numbers: []int64{81, 40, 27}},
			{Result: 83, Numbers: []int64{17, 5}},
			{Result: 156, Numbers: []int64{15, 6}},
			{Result: 7290, Numbers: []int64{6, 8, 6, 15}},
			{Result: 161011, Numbers: []int64{16, 10, 13}},
			{Result: 192, Numbers: []int64{17, 8, 14}},
			{Result: 21037, Numbers: []int64{9, 7, 18, 13}},
			{Result: 292, Numbers: []int64{11, 6, 16, 20}},
		}
		result := day07.SumSolvableEquations(input)
		if result != 3749 {
			t.Errorf("expected 3749, but got %v", result)
		}
	})
}
