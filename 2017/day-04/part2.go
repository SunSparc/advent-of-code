package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valid := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if !hasAnagram(parts) {
			valid = valid + 1
		}
	}
	fmt.Println("valid:", valid)
	// my answer was 223
}

func hasAnagram(passphrase []string) bool {

	for len(passphrase) > 0 {
		var x string
		x, passphrase = passphrase[0], passphrase[1:]
		if wordsMatch(x, passphrase) {
			return true
		}
	}
	return false
}

func wordsMatch(word string, words []string) bool {
	for _, y := range words {
		if wordsContainSameLetters(y, word) {
			return true
		} else {
		}
	}
	return false
}

func wordsContainSameLetters(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	for _, char := range word1 {
		if !containedInWord(char, word2) || len(word2) <= 0 {
			return false
		}
		word2 = removeCharacterFromWord(char, word2)
	}
	return true
}

func containedInWord(letter int32, word string) bool {
	for _, x := range word {
		if x == letter {
			return true
		}
	}

	return false
}

func removeCharacterFromWord(char int32, word string) string {
	var newWord []int32
	found := false
	for _, letter := range word {
		if letter == char && !found {
			found = true
			continue
		}
		newWord = append(newWord, letter)
	}
	return string(newWord)
}

// abcde fghij is a valid passphrase.
// abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
// a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
// iiii oiii ooii oooi oooo is valid.
// oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
