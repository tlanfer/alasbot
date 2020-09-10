package main

import (
	"fmt"
	"math"
)

func main() {
	in := 1386258.0

	out := convert(in)
	fmt.Println(out)
}

func convert(in float64) string {
	minutes := (in /100)*6

	minutesInADay := 24*60

	days := math.Floor(minutes / float64(minutesInADay))+1
	minutesIntoTheDay := math.Mod(minutes, float64(minutesInADay))

	hours := math.Floor(minutesIntoTheDay/60)
	minutesIntoTheHour := math.Mod(minutesIntoTheDay, 60)

	return fmt.Sprintf("days %v, %v:%v", days, hours, minutesIntoTheHour)
}