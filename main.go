package main

import (
	"advent2024/day01"
	"advent2024/day02"
	"advent2024/day03"
	"advent2024/day04"
	"advent2024/day05"
	"advent2024/day06"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Merry Christmas!")
	day01.RunDay()
	day02.RunDay()
	day03.RunDay()
	day04.RunDay()
	day05.RunDay()
	start := time.Now()
	day06.RunDay()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("%v", elapsed)
}
