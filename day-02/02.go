package day02

// http://adventofcode.com/2016/day/2

type BathroomSecurity struct {}

func NewBathroomSecurity() *BathroomSecurity {
	return &BathroomSecurity{}
}

func (this *BathroomSecurity) PushButtons(instructions string) int {
	if instructions == "U" {
		return 2
	} else if instructions == "D" {
		return 8
	} else if instructions == "L" {
		return 4
	} else if instructions == "R" {
		return 6
	} else if instructions == "UL" {
		return 1
	} else if instructions == "LU" {
		return 1
	} else if instructions == "DL" {
		return 7
	} else if instructions == "RU" {
		return 3
	} else if instructions == "RD" {
		return 9
	}
	return 0
}