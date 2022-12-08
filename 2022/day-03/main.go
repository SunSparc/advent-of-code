package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	log.Println("answer:", processInput(loadInput("input")))
}

func processInput(input []string) int32 {
	var answer int32
	for _, v := range input {
		parts := splitStringIntoTwo(v)
		commonLetter := findCommonLetter(parts[0], parts[1])

		numberOfLetter := convertLetterToNumber(commonLetter)
		answer = answer + numberOfLetter
	}
	return answer
}

func convertLetterToNumber(letter rune) (convertedLetter int32) {
	if letter >= 97 {
		return letter - 96
	}
	if letter <= 90 {
		return letter - 38
	}

	log.Fatalln("unknown letter")
	return -1
}

func findCommonLetter(string1, string2 string) (common rune) {
	for _, l1 := range string1 {
		for _, l2 := range string2 {
			if l1 == l2 {
				return l1
			}
		}
	}
	return common
}

func splitStringIntoTwo(stringToSplit string) []string {
	stringSize := len(stringToSplit)
	halfStringSize := stringSize / 2
	half1, half2 := stringToSplit[0:halfStringSize], stringToSplit[halfStringSize:stringSize]
	return []string{half1, half2}
}

///////////////////////////////////

func loadInput(filename string) []string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("[ERROR] ioutil.ReadFile:", err)
	}
	return strings.Split(string(fileBytes), "\n")
}
