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
	var totalFuel, totalFuelPlusFuel int
	totalFuel = 0
	for scanner.Scan() {
		moduleMass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		// floor(mass/3) -2 = ?
		massFuel, fuelFuel := calculateFuel(moduleMass)
		totalFuel = totalFuel + massFuel
		totalFuelPlusFuel = totalFuelPlusFuel + fuelFuel
	}
	fmt.Println("total mass fuel:", totalFuel)
	fmt.Println("total fuel fuel:", totalFuelPlusFuel)
	fmt.Println("grand total:", totalFuel + totalFuelPlusFuel)
}

func calculateFuel(mass int) (int,int) {
	massFuel := mass/3-2
	if mass < 0 || massFuel < 0 {
		return mass, 0
	}
	fuelFuel := 0
	fuelCounter := massFuel
	for {
		fuelCounter = fuelCounter/3-2
		if fuelCounter < 0 {
			break
		}
		fuelFuel = fuelFuel + fuelCounter
	}
	return massFuel, fuelFuel
}