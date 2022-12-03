package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalln("[ERROR] ioutil.ReadFile:", err)
	}
	input := strings.Split(string(fileBytes), "\n")

	var elfCalorieTotals []int
	var winningElf, winningElf2, winningElf3 int
	var winningTotal, winningTotal2, winningTotal3 int
	currentElf := 1
	currentElfCalories := 0
	for _, caloriesString := range input {
		log.Printf("caloriesString: %s", caloriesString)
		if caloriesString == "" {
			currentElf = currentElf + 1
			elfCalorieTotals = append(elfCalorieTotals, currentElfCalories)
			if currentElfCalories > winningTotal {
				winningTotal = currentElfCalories
				winningElf = currentElf
			} else if currentElfCalories > winningTotal2 {
				winningTotal2 = currentElfCalories
				winningElf2 = currentElf
			} else if currentElfCalories > winningTotal3 {
				winningTotal3 = currentElfCalories
				winningElf3 = currentElf
			}
			currentElfCalories = 0
			continue
		}
		caloriesNumeric, err := strconv.Atoi(caloriesString)
		if err != nil {
			log.Println("[ERROR] strconv.Atoi:", err)
		}
		currentElfCalories = currentElfCalories + caloriesNumeric
	}

	log.Printf("elfCalorieTotals: %v", elfCalorieTotals)
	log.Printf("winning elf: %v", winningElf)
	log.Printf("winning total: %v", winningTotal)
	log.Printf("winning elf2: %v", winningElf2)
	log.Printf("winning total2: %v", winningTotal2)
	log.Printf("winning elf3: %v", winningElf3)
	log.Printf("winning total3: %v", winningTotal3)
	log.Printf("winning totals combined: %v", winningTotal+winningTotal2+winningTotal3) // 201464 was too low
	sort.Ints(elfCalorieTotals)
	log.Printf("elfCalorieTotals: %v", elfCalorieTotals[len(elfCalorieTotals)-1]+elfCalorieTotals[len(elfCalorieTotals)-2]+elfCalorieTotals[len(elfCalorieTotals)-3])
}
