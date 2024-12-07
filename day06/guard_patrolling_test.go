package day06_test

import (
	"advent2024/day06"
	"testing"
)

func TestGuardPatrol(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		tiles := [][]string{
			{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", "#", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", "#", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", "#", ".", ".", "^", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
			{"#", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		}
		world := day06.World{
			Tiles:  tiles,
			Height: len(tiles),
			Width:  len(tiles[0]),
		}
		result := day06.CalculateTilesWalked(world)
		if result != 41 {
			t.Errorf("expected 41, but got %v", result)
		}
	})

	t.Run("detect loop", func(t *testing.T) {
		tiles := [][]string{
			{".", "#", ".", "."},
			{".", ".", ".", "#"},
			{".", "^", ".", "."},
			{".", ".", "#", "."},
		}
		world := day06.World{
			Tiles:  tiles,
			Height: len(tiles),
			Width:  len(tiles[0]),
		}
		result := day06.CalculateLoopablePositions(world)
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})
}
