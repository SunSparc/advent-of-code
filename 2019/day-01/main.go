package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var totalMass int
	totalMass = 0
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		// floor(mass/3) -2 = ?
		totalMass = totalMass + number/3-2
	}
	fmt.Println("total mass:", totalMass)
}
