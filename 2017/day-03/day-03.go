package main

import (
	"fmt"
	"math"
)

type Day03 struct {
	direction                  string
	coordinateX                int
	coordinateY                int
	lastNumberOfSteps          int
	maxStepsInThisDirection    int
	stepsInTheCurrentDirection int
}

func NewDay03() *Day03 {
	return &Day03{
		direction:                  "right",
		coordinateX:                0,
		coordinateY:                0,
		lastNumberOfSteps:          0,
		maxStepsInThisDirection:    1,
		stepsInTheCurrentDirection: 0,
	}
}

func main() {
	day := NewDay03()
	for gridSquare := 1; gridSquare < 289326; gridSquare++ {

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
	fmt.Println("steps:", math.Abs(float64(day.coordinateX))+math.Abs(float64(day.coordinateY)))
	// 419 is the correct answer for my input
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
