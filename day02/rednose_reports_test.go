package day02_test

import (
	"advent2024/day02"
	"testing"
)

func TestRedNosedReports(t *testing.T) {
	t.Run("7 6 4 2 1 = safe", func(t *testing.T) {
		report := []int64{7, 6, 4, 2, 1}
		result := day02.CheckPlantReport(report)

		if !result {
			t.Errorf("expected safe, but was unsafe")
		}
	})
	t.Run("1 2 7 8 9 = unsafe", func(t *testing.T) {
		report := []int64{1, 2, 7, 8, 9}
		result := day02.CheckPlantReport(report)

		if result {
			t.Errorf("expected unsafe, but was safe")
		}
	})
	t.Run("9 7 6 2 1 = unsafe", func(t *testing.T) {
		report := []int64{9, 7, 6, 2, 1}
		result := day02.CheckPlantReport(report)

		if result {
			t.Errorf("expected unsafe, but was safe")
		}
	})
	t.Run("1 3 2 4 5 = unsafe", func(t *testing.T) {
		report := []int64{1, 3, 2, 4, 5}
		result := day02.CheckPlantReport(report)

		if result {
			t.Errorf("expected unsafe, but was safe")
		}
	})
	t.Run("1 3 2 4 5 = safe with dampener", func(t *testing.T) {
		report := []int64{1, 3, 2, 4, 5}
		result := day02.CheckPlantReportWithDampening(report)

		if !result {
			t.Errorf("expected safe, but was unsafe")
		}
	})
	t.Run("8 6 4 4 1 = unsafe", func(t *testing.T) {
		report := []int64{8, 6, 4, 4, 1}
		result := day02.CheckPlantReport(report)

		if result {
			t.Errorf("expected unsafe, but was safe")
		}
	})
	t.Run("8 6 4 4 1 = safe with dampener", func(t *testing.T) {
		report := []int64{8, 6, 4, 4, 1}
		result := day02.CheckPlantReportWithDampening(report)

		if !result {
			t.Errorf("expected safe, but was unsafe")
		}
	})
	t.Run("1 3 6 7 9 = safe", func(t *testing.T) {
		report := []int64{1, 3, 6, 7, 9}
		result := day02.CheckPlantReport(report)

		if !result {
			t.Errorf("expected safe, but was unsafe")
		}
	})
	t.Run("3 safe reports counts 3", func(t *testing.T) {
		report1 := []int64{1, 3, 6, 7, 9}
		report2 := []int64{1, 3, 6, 7, 9}
		report3 := []int64{1, 3, 6, 7, 9}
		reports := [][]int64{report1, report2, report3}
		result1 := day02.CountSafeReports(reports, day02.CheckPlantReport)
		result2 := day02.CountSafeReports(reports, day02.CheckPlantReportWithDampening)

		if result1 != 3 {
			t.Errorf("expected 3, but got %v", result1)
		}
		if result2 != 3 {
			t.Errorf("expected 3, but got %v", result2)
		}
	})
}
