package day04_test

import (
	"advent2024/day04"
	"testing"
)

// can I convert 2d to 1d? the offset my search by 'row length'?

func TestCorruptedMemory(t *testing.T) {
	t.Run("xmas", func(t *testing.T) {
		input := [][]rune{
			{'x', 'm', 'a', 's'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("samx", func(t *testing.T) {
		input := [][]rune{
			{'s', 'a', 'm', 'x'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("xmas diagonal descending", func(t *testing.T) {
		input := [][]rune{
			{'x', '.', '.', '.'},
			{'.', 'm', '.', '.'},
			{'.', '.', 'a', '.'},
			{'.', '.', '.', 's'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("samx diagonal descending", func(t *testing.T) {
		input := [][]rune{
			{'s', '.', '.', '.'},
			{'.', 'a', '.', '.'},
			{'.', '.', 'm', '.'},
			{'.', '.', '.', 'x'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("samx diagonal ascending", func(t *testing.T) {
		input := [][]rune{
			{'.', '.', '.', 'x'},
			{'.', '.', 'm', '.'},
			{'.', 'a', '.', '.'},
			{'s', '.', '.', '.'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("xmas diagonal ascending", func(t *testing.T) {
		input := [][]rune{
			{'.', '.', '.', 's'},
			{'.', '.', 'a', '.'},
			{'.', 'm', '.', '.'},
			{'x', '.', '.', '.'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("xmas vertical", func(t *testing.T) {
		input := [][]rune{
			{'x', '.', '.', '.'},
			{'m', '.', '.', '.'},
			{'a', '.', '.', '.'},
			{'s', '.', '.', '.'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("samx vertical", func(t *testing.T) {
		input := [][]rune{
			{'s', '.', '.', '.'},
			{'a', '.', '.', '.'},
			{'m', '.', '.', '.'},
			{'x', '.', '.', '.'},
		}
		result := day04.CountXMAS(input)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})

	t.Run("xmas example", func(t *testing.T) {
		input := [][]rune{
			{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
			{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
			{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
			{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
			{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
			{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
			{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
			{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
			{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
			{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
		}
		result := day04.CountXMAS(input)
		if result != 18 {
			t.Errorf("expected 18, but got %v", result)
		}
	})

	t.Run("xmas example with .", func(t *testing.T) {
		input := [][]rune{
			{'.', '.', '.', '.', 'X', 'X', 'M', 'A', 'S', '.'},
			{'.', 'S', 'A', 'M', 'X', 'M', 'S', '.', '.', '.'},
			{'.', '.', '.', 'S', '.', '.', 'A', '.', '.', '.'},
			{'.', '.', 'A', '.', 'A', '.', 'M', 'S', '.', 'X'},
			{'X', 'M', 'A', 'S', 'A', 'M', 'X', '.', 'M', 'M'},
			{'X', '.', '.', '.', '.', '.', 'X', 'A', '.', 'A'},
			{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', 'S'},
			{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', 'A'},
			{'.', '.', 'M', '.', 'M', '.', 'M', '.', 'M', 'M'},
			{'.', 'X', '.', 'X', '.', 'X', 'M', 'A', 'S', 'X'},
		}
		result := day04.CountXMAS(input)
		if result != 18 {
			t.Errorf("expected 18, but got %v", result)
		}
	})

}
