package main

import (
	"fmt"
	"math"
)

func main() {
	in := 1399072.0

	days, hours, minutes := convert(in)

	fmt.Printf("days %v, %02d:%02d", days, hours, minutes)
}

func convert(in float64) (int, int, int) {
	totalMinutes := (in /100)*6

	minutesInADay := 24*60

	days := math.Floor(totalMinutes/ float64(minutesInADay))+1
	minutesIntoTheDay := math.Mod(totalMinutes, float64(minutesInADay))

	hours := math.Floor(minutesIntoTheDay/60)
	minutes := math.Mod(minutesIntoTheDay, 60)

	return int(days), int(hours), int(minutes)
}