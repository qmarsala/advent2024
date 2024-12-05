package day04_test

import (
	"advent2024/day04"
	"testing"
)

// can I convert 2d to 1d? the offset my search by 'row length'?

func TestWordSearch(t *testing.T) {
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
		day04.ReadXMAS(input, 0, 3, 1)
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
			{'x', 's', 's', 'x', 's'},
			{'m', 'a', 'm', 'a', 'x'},
			{'a', 'm', 'm', 's', 's'},
			{'s', 'x', 'm', 'a', 's'},
			{'s', 'x', 'm', 'a', 's'},
		}
		result := day04.CountXMAS(input)
		if result != 5 {
			t.Errorf("expected 5, but got %v", result)
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
			{'m', 'm', 'm', 's', 'x', 'x', 'm', 'a', 's', 'm'},
			{'m', 's', 'a', 'm', 'x', 'm', 's', 'm', 's', 'a'},
			{'a', 'm', 'x', 's', 'x', 'm', 'a', 'a', 'm', 'm'},
			{'m', 's', 'a', 'm', 'a', 's', 'm', 's', 'm', 'x'},
			{'x', 'm', 'a', 's', 'a', 'm', 'x', 'a', 'm', 'm'},
			{'x', 'x', 'a', 'm', 'm', 'x', 'x', 'a', 'm', 'a'},
			{'s', 'm', 's', 'm', 's', 'a', 's', 'x', 's', 's'},
			{'s', 'a', 'x', 'a', 'm', 'a', 's', 'a', 'a', 'a'},
			{'m', 'a', 'm', 'm', 'm', 'x', 'm', 'm', 'm', 'm'},
			{'m', 'x', 'm', 'x', 'a', 'x', 'm', 'a', 's', 'x'},
		}
		result := day04.CountXMAS(input)
		if result != 18 {
			t.Errorf("expected 18, but got %v", result)
		}
	})

	t.Run("x-mas", func(t *testing.T) {
		input := [][]rune{
			{'m', 'm', 'm', 's', 'x', 'x', 'm', 'a', 's', 'm'},
			{'m', 's', 'a', 'm', 'x', 'm', 's', 'm', 's', 'a'},
			{'a', 'm', 'x', 's', 'x', 'm', 'a', 'a', 'm', 'm'},
			{'m', 's', 'a', 'm', 'a', 's', 'm', 's', 'm', 'x'},
			{'x', 'm', 'a', 's', 'a', 'm', 'x', 'a', 'm', 'm'},
			{'x', 'x', 'a', 'm', 'm', 'x', 'x', 'a', 'm', 'a'},
			{'s', 'm', 's', 'm', 's', 'a', 's', 'x', 's', 's'},
			{'s', 'a', 'x', 'a', 'm', 'a', 's', 'a', 'a', 'a'},
			{'m', 'a', 'm', 'm', 'm', 'x', 'm', 'm', 'm', 'm'},
			{'m', 'x', 'm', 'x', 'a', 'x', 'm', 'a', 's', 'x'},
		}
		result := day04.CountXMAS2(input)
		if result != 9 {
			t.Errorf("expected 9, but got %v", result)
		}
	})

}
