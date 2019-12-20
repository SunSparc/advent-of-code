package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coordinate struct {
	x, y int
	runningTotal int
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	// input has two lines, representing two wires
	//bytes = []byte("R8,U5,L5,D3\nU7,R6,D4,L4")
	//bytes = []byte("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83") // = 159
	//bytes = []byte("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7") // = 135
	fileString := string(bytes)
	wires := strings.Split(fileString, "\n")
	wireOne := strings.Split(wires[0], ",")
	wireTwo := strings.Split(wires[1], ",")

	wireOnePath := pathFinder(wireOne)
	wireTwoPath := pathFinder(wireTwo)

	matchesDistance, matchesSteps := findMatchingCoordinates(wireOnePath, wireTwoPath)
	fmt.Println("closest distance to central port:", closestToZero(matchesDistance))
	fmt.Println("fewest combined steps:", closestToZero(matchesSteps))
}

func closestToZero(list []int) int {
	closest := 0
	for _, x := range list {
		if x == 0 {
			continue
		}
		if x < closest {
			closest = x
		}
		if closest == 0 {
			closest = x
		}
	}
	return closest
}

func findMatchingCoordinates(wireOnePath, wireTwoPath []coordinate) ([]int,[]int) {
	var runningTotalDistanceOfMatches []int
	var combinedStepsToMatches []int
	for _,w1p := range wireOnePath {
		for _,w2p := range wireTwoPath {
			if w1p.x == w2p.x && w1p.y == w2p.y {
				fmt.Println("match:", w1p, w2p)
				if w1p.x < 0 {
					w1p.x = -w1p.x
				}
				if w1p.y < 0 {
					w1p.y = -w1p.y
				}
				runningTotalDistanceOfMatches = append(runningTotalDistanceOfMatches, w1p.x + w1p.y)
				combinedStepsToMatches = append(combinedStepsToMatches, w1p.runningTotal+w2p.runningTotal)
			}
		}
	}
	return runningTotalDistanceOfMatches, combinedStepsToMatches
}

func pathFinder(input []string) []coordinate {
	path := []coordinate{{0, 0, 0}}
	for _, vector := range input {
		direction := vector[0]
		magnitude, _ := strconv.Atoi(vector[1:])
		if direction == Down {
			path = travelDown(path, magnitude)
		}
		if direction == Left {
			path = travelLeft(path, magnitude)
		}
		if direction == Right {
			path = travelRight(path, magnitude)
		}
		if direction == Up {
			path = travelUp(path, magnitude)
		}
	}
	return path
}

func travelUp(path []coordinate, magnitude int) []coordinate {
	for x := 0; x < magnitude; x++ {
		path = append(path, coordinate{
			x: path[len(path)-1].x,
			y: path[len(path)-1].y + 1,
			runningTotal: path[len(path)-1].runningTotal+1,
		})
	}
	return path
}
func travelLeft(path []coordinate, magnitude int) []coordinate {
	for x := 0; x < magnitude; x++ {
		path = append(path, coordinate{
			x: path[len(path)-1].x-1,
			y: path[len(path)-1].y,
			runningTotal: path[len(path)-1].runningTotal+1,
		})
	}
	return path
}
func travelRight(path []coordinate, magnitude int) []coordinate {
	for x := 0; x < magnitude; x++ {
		path = append(path, coordinate{
			x: path[len(path)-1].x + 1,
			y: path[len(path)-1].y,
			runningTotal: path[len(path)-1].runningTotal+1,
		})
	}
	return path
}
func travelDown(path []coordinate, magnitude int) []coordinate {
	for x := 0; x < magnitude; x++ {
		path = append(path, coordinate{
			x: path[len(path)-1].x,
			y: path[len(path)-1].y - 1,
			runningTotal: path[len(path)-1].runningTotal+1,
		})
	}
	return path
}



const (
	Down  = 68 // -y
	Left  = 76 // -x
	Right = 82 // +x
	Up    = 85 // +y
)
