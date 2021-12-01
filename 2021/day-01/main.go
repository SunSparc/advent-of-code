package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalln("[ERROR] ioutil.ReadFile:", err)
	}
	input := strings.Split(string(fileBytes), "\n")
	measurementsLargerThanThePrevious := 0
	previousMeasurement := 0
	for _,measurement := range input {
		currentMeasurement, _ := strconv.Atoi(measurement)
		if previousMeasurement == 0 {
			log.Printf("%s: no previous measurement", measurement)
			previousMeasurement = currentMeasurement
			continue
		}
		if previousMeasurement < currentMeasurement {
			measurementsLargerThanThePrevious = measurementsLargerThanThePrevious + 1
		}
		previousMeasurement = currentMeasurement
	}
	log.Println("The number of measurements that are larger than the previous measurement:", measurementsLargerThanThePrevious)
}
