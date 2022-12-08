package main

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestMainFixture(t *testing.T) {
	gunit.Run(new(Main), t)
}

type Main struct {
	*gunit.Fixture
}

func (this *Main) Setup() {
}

// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw

// The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp,
//   which means its first compartment contains the items vJrwpWtwJgWr,
//   while the second compartment contains the items hcsFMMfFFhFp.
//   The only item type that appears in both compartments is lowercase p.
// The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL.
//   The only item type that appears in both compartments is uppercase L.
// The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg;
//   the only common item type is uppercase P.
// The fourth rucksack's compartments only share item type v.
// The fifth rucksack's compartments only share item type t.
// The sixth rucksack's compartments only share item type s.

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.

func (this *Main) TestInputCanBeProcessed() {
	input := loadInput("input_test")
	this.So(input, should.HaveSameTypeAs, []string{""})
	answer := processInput(input)
	this.So(answer, should.Equal, 157)
}

func (this *Main) TestSplitStringIntoTwoEqualParts() {
	testString := "vJrwpWtwJgWrhcsFMMfFFhFp"
	twoHalves := splitStringIntoTwo(testString)
	this.So(twoHalves, should.Resemble, []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"})
}

func (this *Main) TestFindCommonLetter() {
	string1 := "vJrwpWtwJgWr"
	string2 := "hcsFMMfFFhFp"
	common := findCommonLetter(string1, string2)
	this.So(common, should.Equal, 'p')
	string1 = "jqHRNqRjqzjGDLGL"
	string2 = "rsFMfFZSrLrFZsSL"
	common = findCommonLetter(string1, string2)
	this.So(common, should.Equal, 'L')
	string1 = "PmmdzqPrV"
	string2 = "vPwwTWBwg"
	common = findCommonLetter(string1, string2)
	this.So(common, should.Equal, 'P')
	string1 = "wMqvLMZHhHMvwLH"
	string2 = "jbvcjnnSBnvTQFn"
	common = findCommonLetter(string1, string2)
	this.So(common, should.Equal, 'v')
	string1 = "ttgJtRGJ"
	string2 = "QctTZtZT"
	common = findCommonLetter(string1, string2)
	this.So(common, should.Equal, 't')
	string1 = "CrZsJsPPZsGz"
	string2 = "wwsLwLmpwMDw"
	common = findCommonLetter(string1, string2)
	this.So(common, should.Equal, 's')
}

func (this *Main) TestMatchLettersToNumbers() {
	// 16 (p) -> ascii 112-96
	// 38 (L) -> ascii 76-38
	// 42 (P) -> ascii 80-38
	// 22 (v) -> ascii 118-96
	// 20 (t) -> ascii 116-96
	// 19 (s) -> ascii 115-96
	// ascii: A=65, Z=90, a=97, z=122
	this.So(convertLetterToNumber('p'), should.Equal, 16)
	this.So(convertLetterToNumber('L'), should.Equal, 38)
	this.So(convertLetterToNumber('P'), should.Equal, 42)
	this.So(convertLetterToNumber('v'), should.Equal, 22)
	this.So(convertLetterToNumber('t'), should.Equal, 20)
	this.So(convertLetterToNumber('s'), should.Equal, 19)
}
