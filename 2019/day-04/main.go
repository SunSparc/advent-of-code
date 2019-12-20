package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "130254-678275"

func main() {
	inputSplit := strings.Split(input, "-")
	rangeStart, _ := strconv.Atoi(inputSplit[0])
	rangeEnd, _ := strconv.Atoi(inputSplit[1])
	countOfPpossiblePasswords := 0
	for x:=rangeStart;x<=rangeEnd;x++ {
		digits := splitDigits(x)
		//digits = []int{1,1,1,1,1,1}
		//digits = []int{2,2,3,4,5,0}
		//digits = []int{1,2,3,7,8,9}
		// check for adjacent digits
		if !foundDecreasingDigits(digits) && foundAdjacentDigits(digits) {
			countOfPpossiblePasswords = countOfPpossiblePasswords + 1
		}
		break
	}
	fmt.Println("count of possible passwords:", countOfPpossiblePasswords)
}

func foundDecreasingDigits(digits []int) bool {
	for i, digit := range digits {
		if i == 0 { continue }
		if digit < digits[i-1] {
			return true
		}
	}
	return false
}
func foundAdjacentDigits(digits []int) bool {
	for i, digit := range digits {
		if i > 0 && digit == digits[i-1] {
			return true
		}
		if i < len(digits)-1 && digit == digits[i+1] {
			return true
		}
	}
	return false
}
func splitDigits(x int) []int {
	digitsString := strings.Split(strconv.Itoa(x), "")
	var digitsInt []int
	for _, digitS := range digitsString {
		digitI, _ := strconv.Atoi(digitS)
		digitsInt = append(digitsInt, digitI)
	}
	return digitsInt
}