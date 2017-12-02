package main

//package AdventofCode2016

// http://adventofcode.com/2016/day/1

import (
	"fmt"
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

var previousLocations []Coordinate
var startingLocation Coordinate
var currentLocation Coordinate

func main() {
	directions := []string{"R1", "R3", "L2", "L5", "L2", "L1", "R3", "L4", "R2", "L2", "L4", "R2", "L1", "R1", "L2", "R3", "L1", "L4", "R2", "L5", "R3", "R4", "L1", "R2", "L1", "R3", "L4", "R5", "L4", "L5", "R5", "L3", "R2", "L3", "L3", "R1", "R3", "L4", "R2", "R5", "L4", "R1", "L1", "L1", "R5", "L2", "R1", "L2", "R188", "L5", "L3", "R5", "R1", "L2", "L4", "R3", "R5", "L3", "R3", "R45", "L4", "R4", "R72", "R2", "R3", "L1", "R1", "L1", "L1", "R192", "L1", "L1", "L1", "L4", "R1", "L2", "L5", "L3", "R5", "L3", "R3", "L4", "L3", "R1", "R4", "L2", "R2", "R3", "L5", "R3", "L1", "R1", "R4", "L2", "L3", "R1", "R3", "L4", "L3", "L4", "L2", "L2", "R1", "R3", "L5", "L1", "R4", "R2", "L4", "L1", "R3", "R3", "R1", "L5", "L2", "R4", "R4", "R2", "R1", "R5", "R5", "L4", "L1", "R5", "R3", "R4", "R5", "R3", "L1", "L2", "L4", "R1", "R4", "R5", "L2", "L3", "R4", "L4", "R2", "L2", "L4", "L2", "R5", "R1", "R4", "R3", "R5", "L4", "L4", "L5", "L5", "R3", "R4", "L1", "L3", "R2", "L2", "R1", "L3", "L5", "R5", "R5", "R3", "L4", "L2", "R4", "R5", "R1", "R4", "L3"} // 307
	// directions := []string{"R2", "L3"} // = 5 block
	// directions := []string{"R2", "R2", "R2"} // = 2 blocks
	// directions := []string{"R5", "L5", "R5", "R3"} // = 12 blocks
	// directions := []string{"R1", "L2", "L1", "L2", "L6", "R1", "L2", "R1", "L2", "R1", "L2"}
	fmt.Println("Directions:", directions)

	startingLocation = Coordinate{0, 0}
	currentLocation = startingLocation
	currentOrientation := NORTH

	for _, step := range directions {

		newDirection := step[0:1]
		numberOfBlocks, err := strconv.Atoi(step[1:len(step)])
		if err != nil {
			fmt.Println("strconv.Atoi error:", err)
		}
		if newDirection == "R" {
			currentOrientation = currentOrientation + 1

			if currentOrientation > WEST {
				currentOrientation = NORTH
			}
		} else {
			currentOrientation = currentOrientation - 1

			if currentOrientation < NORTH {
				currentOrientation = WEST
			}
		}

		switch currentOrientation {
		case NORTH:
			localLocation := currentLocation
			for x := 0; x < numberOfBlocks; x++ {
				findMatchingLocation(localLocation)
				previousLocations = append(previousLocations, localLocation)
				localLocation.y = localLocation.y + 1
			}
			currentLocation.y = currentLocation.y + numberOfBlocks
		case EAST:
			localLocation := currentLocation
			for x := 0; x < numberOfBlocks; x++ {
				findMatchingLocation(localLocation)
				previousLocations = append(previousLocations, localLocation)
				localLocation.x = localLocation.x + 1
			}
			currentLocation.x = currentLocation.x + numberOfBlocks
		case SOUTH:
			localLocation := currentLocation
			for x := 0; x < numberOfBlocks; x++ {
				findMatchingLocation(localLocation)
				previousLocations = append(previousLocations, localLocation)
				localLocation.y = localLocation.y - 1
			}
			currentLocation.y = currentLocation.y - numberOfBlocks
		case WEST:
			localLocation := currentLocation
			for x := 0; x < numberOfBlocks; x++ {
				findMatchingLocation(localLocation)
				previousLocations = append(previousLocations, localLocation)
				localLocation.x = localLocation.x - 1
			}
			currentLocation.x = currentLocation.x - numberOfBlocks
		}
	}
	fmt.Println("currentLocation:", currentLocation)

	fmt.Println("distance:", findDistance(startingLocation, currentLocation))
}

func findDistance(start, end Coordinate) int {
	first, second := 0, 0
	// Taxicab Equation: d = |x2 - x1| + |y2 - y1|
	if end.x > start.x {
		first = end.x - start.x
	} else {
		first = start.x - end.x
	}
	if end.y > start.y {
		second = end.y - start.y
	} else {
		second = start.y - end.y
	}
	return first + second
}

func findMatchingLocation(newLocation Coordinate) {
	for _, location := range previousLocations {
		if location == newLocation {
			fmt.Println("we have been here before", findDistance(startingLocation, location))
		}
	}
}

const (
	NORTH = 0 // x+1 // y+1
	EAST  = 1 // y+1 // x+1
	SOUTH = 2 // x-1 // y-1
	WEST  = 3 // y-1 // x-1
)
