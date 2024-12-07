package day06

import (
	"advent2024/fileio"
	"fmt"
	"strings"
)

func RunDay() {
	input, err := ReadGuardPatrolInputs()
	if err != nil {
		fmt.Printf("Day06 - [ERROR]: %v\n", err)
		return
	}
	part1 := CalculateTilesWalked(input)
	part2 := CalculateLoopablePositions(input)
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

func CalculateTilesWalked(world World) int {
	visitedLocations := map[Location]int{}
	if found, currentPosition, currentOrientation := FindGuard(world); found {
		visitedLocations[currentPosition] = 1
		finished := false
		for !finished {
			finished, currentPosition, currentOrientation = Move(world, currentPosition, currentOrientation)
			visitedLocations[currentPosition] += 1
		}
	} else {
		return 0
	}
	return len(visitedLocations) - 1
}

func CalculateLoopablePositions(world World) int {
	visitedLocations := map[Location]int{}
	loopCount := 0
	if found, startingPosition, startingOrientation := FindGuard(world); found {
		visitedLocations[startingPosition] = 1
		loopDetected := false
		finished := false
		for y := 0; y < world.Height; y++ {
			for x := 0; x < world.Width; x++ {
				if x == startingPosition.X && y == startingPosition.Y || world.Tiles[y][x] == "#" {
					continue
				}
				newWorld := CopyWorld(world)
				newWorld.Tiles[y][x] = "#"
				currentPosition := startingPosition
				currentOrientation := startingOrientation
				for !(finished || loopDetected) {
					finished, currentPosition, currentOrientation = Move(
						newWorld,
						currentPosition,
						currentOrientation)
					visitedLocations[currentPosition] += 1
					loopDetected = visitedLocations[currentPosition] > 5
				}
				if loopDetected {
					loopCount += 1
				}
				visitedLocations = map[Location]int{}
				loopDetected = false
				finished = false
			}
		}
	} else {
		return 0
	}
	return loopCount
}

func CopyWorld(world World) World {
	newTiles := [][]string{}
	for _, r := range world.Tiles {
		newRow := make([]string, world.Width)
		copy(newRow, r)
		newTiles = append(newTiles, newRow)
	}
	return World{
		Height: world.Height,
		Width:  world.Width,
		Tiles:  newTiles,
	}
}

func FindGuard(world World) (found bool, loc Location, orientation Direction) {
	for y := 0; y < world.Height; y++ {
		for x := 0; x < world.Width; x++ {
			if world.Tiles[y][x] == "^" {
				return true, Location{X: x, Y: y}, north
			}
			if world.Tiles[y][x] == ">" {
				return true, Location{X: x, Y: y}, east
			}
			if world.Tiles[y][x] == "v" {
				return true, Location{X: x, Y: y}, south
			}
			if world.Tiles[y][x] == "v" {
				return true, Location{X: x, Y: y}, west
			}
		}
	}
	return false, Location{}, -1
}

// could we scan a whole row or column at once?
func Move(world World, currentPosition Location, currentOrientation Direction) (finished bool, endLoc Location, endOrientation Direction) {
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
	if newPosition.Y >= 0 && newPosition.Y < world.Height && newPosition.X >= 0 && newPosition.X < world.Width {
		if world.Tiles[newPosition.Y][newPosition.X] == "#" {
			return false, currentPosition, RotateGuard(currentOrientation)
		}
		return false, newPosition, currentOrientation
	} else {
		return true, newPosition, currentOrientation
	}
}

func RotateGuard(currentOrientation Direction) Direction {
	newOrientation := currentOrientation + 1
	if newOrientation > west {
		newOrientation = north
	}
	return newOrientation
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
