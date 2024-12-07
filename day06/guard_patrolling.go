package day06

import (
	"advent2024/fileio"
	"fmt"
	"strings"
)

func RunDay(name string) {
	input, err := ReadGuardPatrolInputs()
	if err != nil {
		fmt.Printf("%v - [ERROR]: %v\n", name, err)
		return
	}
	part1 := CalculateTilesWalked(input)
	part2 := CalculateLoopablePositions(input)
	fmt.Printf("%v: %v, %v \n", name, part1, part2)
}

func ReadGuardPatrolInputs() (world World, err error) {
	tiles := [][]string{}
	err = fileio.ParseAllLines("./day06/input.txt", func(line string) {
		tiles = append(tiles, strings.Split(line, ""))
	})
	if err != nil {
		return World{}, err
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
	visitedLocations := Patrol(world)
	return len(visitedLocations) - 1
}

func Patrol(world World) map[Location]int {
	visitedLocations := map[Location]int{}
	if found, currentPosition, currentOrientation := FindGuard(world); found {
		visitedLocations[currentPosition] = 1
		finished := false
		for !finished {
			finished, currentPosition, currentOrientation = Move(world, currentPosition, currentOrientation)
			visitedLocations[currentPosition] += 1
		}
	}
	return visitedLocations
}

func CalculateLoopablePositions(world World) int {
	c := make(chan int)
	go CalculateLoopablePositionsInArea(world, 0, world.Width/2, 0, world.Height/2, c)
	go CalculateLoopablePositionsInArea(world, (world.Width / 2), world.Width, 0, world.Height/2, c)
	go CalculateLoopablePositionsInArea(world, 0, world.Width/2, world.Height/2, world.Height, c)
	go CalculateLoopablePositionsInArea(world, (world.Width / 2), world.Width, world.Height/2, world.Height, c)
	topLeft, topRight, bottomLeft, bottomRight := <-c, <-c, <-c, <-c
	return topLeft + topRight + bottomLeft + bottomRight
}

func CalculateLoopablePositionsInArea(world World, startX int, endX int, startY int, endY int, c chan int) int {
	loopCount := 0
	if found, startingPosition, startingOrientation := FindGuard(world); found {
		route := Patrol(world)
		visitedLocations := map[Location]int{}
		visitedLocations[startingPosition] = 1
		for y := startY; y < endY; y++ {
			for x := startX; x < endX; x++ {
				if x == startingPosition.X && y == startingPosition.Y || world.Tiles[y][x] == "#" {
					continue
				}
				_, westOfRoute := route[Location{X: x - 1, Y: y}]
				_, eastOfRoute := route[Location{X: x + 1, Y: y}]
				_, northOfRoute := route[Location{X: x, Y: y - 1}]
				_, southOfRoute := route[Location{X: x, Y: y + 1}]
				if !(westOfRoute || eastOfRoute || northOfRoute || southOfRoute) {
					continue
				}

				newWorld := CopyWorld(world)
				newWorld.Tiles[y][x] = "#"
				currentPosition := startingPosition
				currentOrientation := startingOrientation
				loopDetected := false
				finished := false
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
				visitedLocations[startingPosition] = 1
			}
		}
	}
	c <- loopCount
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
