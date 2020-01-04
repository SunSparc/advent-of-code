package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "130254-678275" // 548,021 digits in range

func main() {
	inputSplit := strings.Split(input, "-")
	rangeStart, _ := strconv.Atoi(inputSplit[0])
	rangeEnd, _ := strconv.Atoi(inputSplit[1])
	countOfPossiblePasswords := 0
	for x := rangeStart; x <= rangeEnd; x++ {
		digits := splitDigits(x)
		//digits = []int{1,1,1,1,1,1} // test part 1: should pass (should fail part 2)
		//digits = []int{2, 2, 3, 4, 5} // test part 1: should pass (should pass part 2)
		//digits = []int{1,2,3,7,8,9} // test part 1: should fail
		//digits = []int{1,1,2,2,3,3} // test part 2: should pass
		//digits = []int{1,2,3,4,4,4} // test part 2: should fail
		//digits = []int{1,1,1,1,2,2} // test part 2: should pass (1111 does not, 22 does)
		//digits = []int{2,2,3,3,3,3} // test part 2: should pass (3333 does not, 22 does)
		// check for adjacent digits
		if !foundDecreasingDigits(digits) && foundSameAdjacentDigits(digits) {
			countOfPossiblePasswords = countOfPossiblePasswords + 1
		}
		//break // for testing
	}
	fmt.Println("count of possible passwords:", countOfPossiblePasswords)
}

func foundDecreasingDigits(digits []int) bool {
	for i, digit := range digits {
		if i == 0 {
			continue
		}
		if digit < digits[i-1] {
			return true
		}
	}
	return false
}
func foundSameAdjacentDigits(digits []int) bool {
	for i, digit := range digits {
		if i > 0 && digit == digits[i-1] {
			if i > 1 && digit == digits[i-2] {
				continue
			}
			if i < len(digits)-1 && digit == digits[i+1] {
				continue
			}
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
