package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}
type Size struct {
	Length, Width int
}

var theOneMap = map[Coordinate]bool{}

func main() {
	file, err := os.Open("../input.txt") // input.test
	if err != nil {
		fmt.Println("os.Open Error:", err)
	}
	scanner := bufio.NewScanner(file)

	// loop through all claims
	for scanner.Scan() {
		recordClaimCoordinates(splitClaimIntoComponents(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}

	fmt.Println("claim overlaps:", countClaimOverlaps())
}

func splitClaimIntoComponents(claim string) (Coordinate, Size) {
	// #1 @ 871,327: 16x20
	fourParts := strings.Split(claim, " ")

	coordinateString := strings.Replace(fourParts[2], ":", "", 1)
	coordinateStringSlice := strings.Split(coordinateString, ",")
	x, _ := strconv.Atoi(coordinateStringSlice[0])
	y, _ := strconv.Atoi(coordinateStringSlice[1])
	topLeftCoordinate := Coordinate{x + 1, y + 1} // "from the edge" - see puzzle.txt

	sizeSlice := strings.Split(fourParts[3], "x")
	length, _ := strconv.Atoi(sizeSlice[0])
	width, _ := strconv.Atoi(sizeSlice[1])
	size := Size{length, width}

	return topLeftCoordinate, size
}

func recordClaimCoordinates(topLeftCoordinate Coordinate, size Size) {
	rightSide := topLeftCoordinate.X + size.Length
	bottomSide := topLeftCoordinate.Y + size.Width
	for xCoordinate := topLeftCoordinate.X; xCoordinate < rightSide; xCoordinate++ {
		for yCoordinate := topLeftCoordinate.Y; yCoordinate < bottomSide; yCoordinate++ {
			if _, ok := theOneMap[Coordinate{X: xCoordinate, Y: yCoordinate}]; ok {
				theOneMap[Coordinate{X: xCoordinate, Y: yCoordinate}] = true
			} else {
				theOneMap[Coordinate{X: xCoordinate, Y: yCoordinate}] = false
			}
		}
	}
}

func countClaimOverlaps() int {
	numberOfOverlaps := 0
	for _, claimBool := range theOneMap {
		if claimBool {
			numberOfOverlaps = numberOfOverlaps + 1
		}
	}
	return numberOfOverlaps
}
