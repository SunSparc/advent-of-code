package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var score int
	for _, round := range loadInput() {
		shapes := strings.Split(round, " ")
		me := getValue(shapes[1])

		if (shapes[0] == "A" && shapes[1] == "X") ||
			(shapes[0] == "B" && shapes[1] == "Y") ||
			(shapes[0] == "C" && shapes[1] == "Z") {
			// draw = 3
			log.Printf("[DRAW] score(%v) + draw(3) + shape(%v)", score, me)
			score = score + DRAW + me
		} else if (shapes[0] == "A" && shapes[1] == "Y") ||
			(shapes[0] == "B" && shapes[1] == "Z") ||
			(shapes[0] == "C" && shapes[1] == "X") {
			// win = 6
			log.Printf("[WIN] score(%v) + win(6) + shape(%v)", score, me)
			score = score + WIN + me
		} else if (shapes[0] == "A" && shapes[1] == "Z") ||
			(shapes[0] == "B" && shapes[1] == "X") ||
			(shapes[0] == "C" && shapes[1] == "Y") {
			// lose = 0
			log.Printf("[LOSE] score(%v) + lose(0) + shape(%v)", score, me)
			score = score + LOSE + me
		}
	}
	log.Println("score:", score) // 15632
}

func getValue(shape string) int {
	switch shape {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		log.Fatalln("[ERROR] unknown shape:", shape)
		return 0
	}
}

// rock vs rock = draw
// rock vs paper = lose
// rock vs scissors = win

// paper vs paper = draw
// paper vs scissors = lose
// paper vs rock = win

// scissors vs scissors = draw
// scissors vs rock = lose
// scissors vs paper = win

/////////////////////////////////

// A vs X = draw
// B vs Y = draw
// C vs Z = draw

// A vs Y = win
// B vs Z = win
// C vs X = win

// A vs Z = lose
// B vs X = lose
// C vs Y = lose

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0
)

////////////////////

func loadInput() []string {
	fileBytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalln("[ERROR] ioutil.ReadFile:", err)
	}
	return strings.Split(string(fileBytes), "\n")
}
