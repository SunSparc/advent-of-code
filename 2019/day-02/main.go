package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	inputString := strings.Split(string(bytes), ",")
	var input []int
	input = make([]int, len(inputString))
	for index, value := range inputString {
		number, _ := strconv.Atoi(value)
		input[index] = number
	}
	input[1] = 12
	input[2] = 2
	for x:=0;x+4<len(input);x=x+4{
		if input[x] == 99 {
			break
		} else if input[x] == 1 {
			// 1,2,3,4 = add value from position 2 with the value from position 3, put the result in position 4
			input[input[x+3]] = input[input[x+1]] + input[input[x+2]]
		} else if input[x] == 2 {
			input[input[x+3]] = input[input[x+1]] * input[input[x+2]]
		}
	}
	fmt.Println(input[0])
}
