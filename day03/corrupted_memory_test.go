package day03_test

import (
	"advent2024/day03"
	"testing"
)

func TestCorruptedMemory(t *testing.T) {
	t.Run("mul(2,4)mul(5,5)mul(11,8)mul(8,5) = 161", func(t *testing.T) {
		input := "mul(2,4)mul(5,5)mul(11,8)mul(8,5)"
		result := day03.Calculate(input)

		if result != 161 {
			t.Errorf("expected 161, but got %v", result)
		}
	})

	t.Run("corrupted command input = 161", func(t *testing.T) {
		input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		result := day03.Calculate(input)

		if result != 161 {
			t.Errorf("expected 161, but got %v", result)
		}
	})

	t.Run("don't()mul(2,4)mul(5,5)mul(11,8)do()mul(8,5) = 35", func(t *testing.T) {
		input := "don't()mul(2,4)mul(5,5)mul(11,8)do()mul(8,5)"
		result := day03.CalculatePartTwo(input)

		if result != 40 {
			t.Errorf("expected 35, but got %v", result)
		}
	})
}
