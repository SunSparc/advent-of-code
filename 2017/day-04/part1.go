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
		if !hasDuplicate(parts) {
			valid = valid + 1
		}
	}
	fmt.Println("valid:", valid)
	// my answer is 451
}

func hasDuplicate(passphrase []string) bool {

	for len(passphrase) > 0 {
		var x string
		x, passphrase = passphrase[0], passphrase[1:]
		if contains(x, passphrase) {
			return true
		}
	}
	return false
}

func contains(word string, words []string) bool {
	for _, x := range words {
		if x == word {
			return true
		}
	}
	return false
}

// aa bb cc dd ee is valid.
// aa bb cc dd aa is not valid - the word aa appears more than once.
// aa bb cc dd aaa is valid - aa and aaa count as different words.
