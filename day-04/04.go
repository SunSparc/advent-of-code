package day_04

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type RoomCheck struct {
	input []string
}

func NewRoomCheck(input []string) *RoomCheck {
	return &RoomCheck{
		input: input,
	}
}

func (this *RoomCheck) CheckAllRooms() int {
	sectorSum := 0
	for _, room := range this.input {
		roomResults := this.ParseRoom(room)
		roomName := strings.Split(roomResults[0].(string), "")
		sort.Strings(roomName)
		mapOfChars := countChars(roomResults[0].(string))
		sortedResults := string(sortRoomNameCharacters(mapOfChars))
		if sortedResults == roomResults[2] {
			sectorSum = sectorSum + roomResults[1].(int)
		}
	}
	return sectorSum
}

func (this *RoomCheck) ParseRoom(room string) []interface{} {
	splitInput := strings.Split(room[:len(room)-1], "[")
	checksum := splitInput[1]
	firstPartOfInput := strings.Split(splitInput[0], "-")
	sectorID, err := strconv.Atoi(firstPartOfInput[len(firstPartOfInput)-1])
	if err != nil {
		fmt.Println("string conversion error:", err)
	}
	roomName := strings.Join(firstPartOfInput[:len(firstPartOfInput)-1], "")
	return []interface{}{roomName, sectorID, checksum}
}

func countChars(name string) map[rune]int {
	charMap := make(map[rune]int)
	for _, char := range name {
		if a, ok := charMap[char]; ok {
			charMap[char] = a + 1
		} else {
			charMap[char] = 1
		}
	}
	return charMap
}

func sortRoomNameCharacters(mapOfChars map[rune]int) []rune {
	var sorted []rune

	highestNumber := getHighestNumber(mapOfChars)
	listsByNumber := make([][]rune, highestNumber+1)
	for key, value := range mapOfChars {
		listsByNumber[value] = append(listsByNumber[value], key)
	}
	for x:=len(listsByNumber)-1; x>0 ;x-- {
		if len(listsByNumber[x]) > 0 {
			sortList := sortList(listsByNumber[x])
			for y:=0;y<len(sortList);y++ {
				sorted = append(sorted, sortList[y])
			}
		}
	}

	return sorted[:5]
}

func getHighestNumber(mapOfChars map[rune]int) int {
	var highest int
	for _, value := range mapOfChars {
		if value > highest {
			highest = value
		}
	}
	return highest
}

func sortList(list []rune) []rune {
	runesAsInts := []int{}
	for _,value := range list {
		runesAsInts = append(runesAsInts, int(value))
	}
	sort.Ints(runesAsInts)

	sorted := []rune{}
	for _,value := range runesAsInts {
		sorted = append(sorted, rune(value))
	}

	return sorted
}
