package day06

import (
	"advent2024/fileio"
	"fmt"
	"strings"
)

func RunDay() {
	_, err := ReadGuardPatrolInputs()
	if err != nil {
		fmt.Printf("Day06 - [ERROR]: %v\n", err)
		return
	}
	part1 := 0
	part2 := 0
	fmt.Printf("Day06: %v, %v \n", part1, part2)
}

func ReadGuardPatrolInputs() (world [][]string, err error) {
	lines, err := fileio.ReadAllLines("./day06/input.txt")
	if err != nil {
		return make([][]string, 0), err
	}
	for _, l := range lines {
		world = append(world, strings.Split(l, ""))
	}
	return world, nil
}

func CalculateTilesWalked(world [][]string) int {
	return 0
}
