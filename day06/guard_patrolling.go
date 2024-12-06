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

func ReadGuardPatrolInputs() (world World, err error) {
	lines, err := fileio.ReadAllLines("./day06/input.txt")
	tiles := [][]string{}
	if err != nil {
		return World{}, err
	}
	for _, l := range lines {
		tiles = append(tiles, strings.Split(l, ""))
	}
	if len(tiles) < 2 {
		return World{}, fmt.Errorf("not enough input rows")
	}
	return World{
		Tiles:  tiles,
		Width:  len(tiles[0]),
		Height: len(tiles),
	}, nil
}

type World struct {
	Tiles  [][]string
	Height int
	Width  int
}

type Location struct {
	X int
	Y int
}

type Direction int

const (
	north = iota
	east
	south
	west
)

func CalculateTilesWalked(world [][]string) int {
	if found, _, _ := FindGuard(world); !found {
		return 0
	}
	return 0
}

func FindGuard(world [][]string) (found bool, loc Location, orientation Direction) {
	for i := 0; i < len(world); i++ {
		for j := 0; j < len(world[i]); j++ {
			if world[i][j] == "^" {
				return true, Location{X: j, Y: i}, north
			}
			if world[i][j] == ">" {
				return true, Location{X: j, Y: i}, east
			}
			if world[i][j] == "v" {
				return true, Location{X: j, Y: i}, south
			}
			if world[i][j] == "v" {
				return true, Location{X: j, Y: i}, west
			}
		}
	}
	return false, Location{}, -1
}

func Move(world [][]string, currentPosition Location, currentOrientation Direction) (finished bool, endLoc Location, endOrientation Direction) {
	newPosition := Location{X: 0, Y: 0}
	switch currentOrientation {
	case north:
		newPosition = MoveNorth(currentPosition)
	case east:
		newPosition = MoveEast(currentPosition)
	case south:
		newPosition = MoveSouth(currentPosition)
	case west:
		newPosition = MoveWest(currentPosition)
	}
	fmt.Println(newPosition)
	//todo: validate pos in bounds
	//todo: turn if new pos would be #
	return false, Location{}, -1
}

func MoveNorth(currentPosition Location) (endLoc Location) {
	return Location{X: currentPosition.X, Y: currentPosition.Y - 1}
}

func MoveEast(currentPosition Location) (endLoc Location) {
	return Location{X: currentPosition.X + 1, Y: currentPosition.Y}
}

func MoveSouth(currentPosition Location) (endLoc Location) {
	return Location{X: currentPosition.X, Y: currentPosition.Y + 1}
}

func MoveWest(currentPosition Location) (endLoc Location) {
	return Location{X: currentPosition.X - 1, Y: currentPosition.Y}
}
