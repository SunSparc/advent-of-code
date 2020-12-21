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
		log.Fatal("[ERROR] ioutil.ReadFile:", err)
	}
	input := strings.Split(string(fileBytes), "\n")
loop:
	for _, v1 := range input {
		num1, _ := strconv.Atoi(v1)
		for _, v2 := range input {
			num2, _ := strconv.Atoi(v2)
			for _, v3 := range input {
				num3, _ := strconv.Atoi(v3)
				if num1+num2+num3 == 2020 {
					log.Printf("%d + %d + %d = %d", num1, num2, num3, num1+num2+num3)
					log.Printf("%d * %d * %d = %d", num1, num2, num3, num1*num2*num3)
					break loop
				}
			}
		}
	}
}
