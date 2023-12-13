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
	var groupOfThree [3]string
	groupIndex := 0
	for _, v := range input {
		// get 3 lines
		groupOfThree[groupIndex] = v
		if groupIndex == 2 {
			commonLetter := findCommonLetter(groupOfThree[0], groupOfThree[1], groupOfThree[2])
			log.Printf("commonLetter: %v(%v)", commonLetter, string(commonLetter))
			numberOfLetter := convertLetterToNumber(commonLetter)
			answer = answer + numberOfLetter
			groupIndex = 0
		}
		groupIndex = groupIndex + 1
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

func findCommonLetter(string1, string2, string3 string) (common rune) {
	for _, l1 := range string1 {
		for _, l2 := range string2 {
			for _, l3 := range string3 {
				if l1 == l2 && l2 == l3 {
					log.Printf("l1: %v(%s), l2: %v(%s), l3: %v(%s)", l1, string(l1), l2, string(l2), l3, string(l3))
					//log.Println("found it:", l1, l2, l3)
					return l1
				}
			}
		}
	}
	log.Fatalln("no common letters found:", common)
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
