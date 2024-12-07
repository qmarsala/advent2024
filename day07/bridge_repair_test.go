package day07_test

import (
	"testing"
)

func TestBridgeRepair(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		result := 0
		if result != 1 {
			t.Errorf("expected 1, but got %v", result)
		}
	})
}
