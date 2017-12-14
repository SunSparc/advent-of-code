package main

import (
	"fmt"
)

type Coordinate struct {
	x     int
	y     int
	value int
}

type Day03 struct {
	direction                  string
	coordinateX                int
	coordinateY                int
	lastNumberOfSteps          int
	maxStepsInThisDirection    int
	stepsInTheCurrentDirection int
	memoryMap                  []Coordinate
}

func NewDay03() *Day03 {
	return &Day03{
		direction:                  "right",
		coordinateX:                0,
		coordinateY:                0,
		lastNumberOfSteps:          0,
		maxStepsInThisDirection:    1,
		stepsInTheCurrentDirection: 0,
		memoryMap:                  []Coordinate{},
	}
}

func main() {
	day := NewDay03()
	for gridSquare := 1; gridSquare <= 60; gridSquare++ {
		square := Coordinate{x: day.coordinateX, y: day.coordinateY}
		square.value = day.checkNeighbors(square)
		day.memoryMap = append(day.memoryMap, square)

		switch day.direction {
		case "up":
			day.coordinateY = day.coordinateY + 1
			break
		case "down":
			day.coordinateY = day.coordinateY - 1
			break
		case "left":
			day.coordinateX = day.coordinateX - 1
			break
		case "right":
			day.coordinateX = day.coordinateX + 1
			break
		}

		day.stepsInTheCurrentDirection = day.stepsInTheCurrentDirection + 1

		if day.stepsInTheCurrentDirection == day.maxStepsInThisDirection {
			day.changeDirection()
			day.stepsInTheCurrentDirection = 0
		}

	}
	for index, value := range day.memoryMap {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
	// my correct answer is 295229
}

func (this *Day03) checkNeighbors(center Coordinate) int {
	topLeft := Coordinate{x: center.x - 1, y: center.y + 1}
	topMid := Coordinate{x: center.x, y: center.y + 1}
	topRight := Coordinate{x: center.x + 1, y: center.y + 1}
	midLeft := Coordinate{x: center.x - 1, y: center.y}
	midRight := Coordinate{x: center.x + 1, y: center.y}
	bottomLeft := Coordinate{x: center.x - 1, y: center.y - 1}
	bottomMid := Coordinate{x: center.x, y: center.y - 1}
	bottomRight := Coordinate{x: center.x + 1, y: center.y - 1}

	sumOfNeighbors := 0
	if value := contains(this.memoryMap, topLeft); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, topMid); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, topRight); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, midLeft); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, midRight); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, bottomLeft); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, bottomMid); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if value := contains(this.memoryMap, bottomRight); value > 0 {
		sumOfNeighbors = sumOfNeighbors + value
	}
	if sumOfNeighbors == 0 {
		sumOfNeighbors = 1
	}

	return sumOfNeighbors
}

func contains(slice []Coordinate, check Coordinate) int {
	for _, value := range slice {
		if check.x == value.x && check.y == value.y {
			return value.value
		}
	}
	return 0
}

func (this *Day03) changeDirection() {
	if this.direction == "right" {
		this.direction = "up"
	} else if this.direction == "up" {
		this.direction = "left"
	} else if this.direction == "left" {
		this.direction = "down"
	} else if this.direction == "down" {
		this.direction = "right"
	}

	if this.lastNumberOfSteps < this.maxStepsInThisDirection {
		this.lastNumberOfSteps = this.maxStepsInThisDirection
	} else {
		this.maxStepsInThisDirection = this.maxStepsInThisDirection + 1
	}

}
