package day02

import (
	"advent2024/fileio"
	"fmt"
	"strconv"
	"strings"
)

func RunDay() {
	reports, err := ReadRedNosedReportInputs()
	if err != nil {
		fmt.Printf("Day02 - [ERROR]: %v\n", err)
		return
	}
	part1 := CountSafeReports(reports, func(r []int64) bool {
		safe, _ := CheckPlantReport(r)
		return safe
	})
	part2 := CountSafeReports(reports, CheckPlantReportWithDampening)
	fmt.Printf("Day02: %v, %v \n", part1, part2)
}

func ReadRedNosedReportInputs() (reports [][]int64, err error) {
	lines, err := fileio.ReadAllLines("./day02/input.txt")
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		values := strings.Split(line, " ")
		report := []int64{}
		for _, v := range values {
			value, _ := strconv.ParseInt(v, 10, 64)
			report = append(report, value)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func CountSafeReports(reports [][]int64, strategy func([]int64) bool) (count int64) {
	for _, r := range reports {
		if safe := strategy(r); safe {
			count += 1
		}
	}
	return
}

func CheckPlantReport(report []int64) (safe bool, problemIndex int) {
	current_level := report[0]
	is_desc := current_level > report[1]
	for i := 1; i < len(report); i++ {
		diff := GetDiff(current_level, report[i])
		if diff < 1 || diff > 3 {
			return false, i - 1
		}
		if (is_desc && current_level < report[i]) || (!is_desc && current_level > report[i]) {
			return false, i - 1
		}
		current_level = report[i]
	}
	return true, -1
}

func CheckPlantReportWithDampening(report []int64) (safe bool) {
	safe, problemIndex := CheckPlantReport(report)
	if !safe {
		new_report := RemoveIndex(report, problemIndex)
		safe, _ = CheckPlantReport(new_report)
	}
	return
}

func GetDiff(a int64, b int64) int64 {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func RemoveIndex(s []int64, index int) []int64 {
	ret := make([]int64, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}