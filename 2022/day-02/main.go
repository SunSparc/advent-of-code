package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var score int
	//test := []string{
	//	"A Y",
	//	"B X",
	//	"C Z",
	//}

	for _, round := range loadInput() {

		shapes := strings.Split(round, " ")

		elfKey := shapes[0]
		outcomeKey := shapes[1]
		if elfKey == "A" && outcomeKey == DRAW {
			// ROCK(A) vs DRAW(Y=3) = ROCK(A=1)
			score = score + drawInt + 1
		}

		if elfKey == "B" && outcomeKey == LOSE {
			// PAPER(B) vs LOSE(X=0) = ROCK(A=1)
			score = score + loseInt + 1
		}

		if elfKey == "C" && outcomeKey == WIN {
			// SCISSORS(C) vs WIN(Z=6) = ROCK(A=1)
			score = score + winInt + 1
		}

		if elfKey == "B" && outcomeKey == DRAW {
			// PAPER(B) vs DRAW(Y=3) = PAPER(B=2)
			score = score + drawInt + 2
		}

		if elfKey == "C" && outcomeKey == LOSE {
			// SCISSORS(C) vs LOSE(X=0) = PAPER(B=2)
			score = score + loseInt + 2
		}

		if elfKey == "A" && outcomeKey == WIN {
			// ROCK(A) vs WIN(Z=6) = PAPER(B=2)
			score = score + winInt + 2
		}

		if elfKey == "C" && outcomeKey == DRAW {
			// SCISSORS(C) vs DRAW(Y=3) = SCISSORS(C=3)
			score = score + drawInt + 3
		}

		if elfKey == "A" && outcomeKey == LOSE {
			// ROCK(A) vs LOSE(X=0) = SCISSORS(C=3)
			score = score + loseInt + 3
		}

		if elfKey == "B" && outcomeKey == WIN {
			// PAPER(B) vs WIN(Z=6) = SCISSORS(C=3)
			score = score + winInt + 3
		}

	}
	log.Println("score:", score) // 15632
}

func getValue(shape string) int {
	switch shape {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	// 	WIN  = 6 // Z
	//	DRAW = 3 // Y
	//	LOSE = 0 // X
	default:
		log.Fatalln("[ERROR] unknown shape:", shape)
		return 0
	}
}

///////////////////////////////////

const (
	WIN     = "Z" // 6 // Z
	DRAW    = "Y" // 3 // Y
	LOSE    = "X" // 0 // X
	winInt  = 6   // Z
	drawInt = 3   // Y
	loseInt = 0   // X
)

var (
	outcomes = map[string]int{
		"Z": 6,
		"Y": 3,
		"X": 0,
	}
)

////////////////////

func loadInput() []string {
	fileBytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalln("[ERROR] ioutil.ReadFile:", err)
	}
	return strings.Split(string(fileBytes), "\n")
}
