package day02

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

type Day02Fixture struct {
	*gunit.Fixture
	security *BathroomSecurity
}

func TestDay02Fixture(t *testing.T) {
	gunit.Run(new(Day02Fixture), t)
}

func (this *Day02Fixture) Setup() {
	this.security = NewBathroomSecurity()
}

func (this *Day02Fixture) Test_OneLine_OneCharacter_Blank() {
	buttons := this.security.PushButtons("")
	this.So(buttons, should.Equal, 0)
}
func (this *Day02Fixture) Test_OneLine_OneCharacter_Up() {
	buttons := this.security.PushButtons("U")
	this.So(buttons, should.Equal, 5)
}
func (this *Day02Fixture) Test_OneLine_OneCharacter_Down() {
	buttons := this.security.PushButtons("D")
	this.So(buttons, should.Equal, 5)
}
func (this *Day02Fixture) Test_OneLine_OneCharacter_Left() {
	buttons := this.security.PushButtons("L")
	this.So(buttons, should.Equal, 5)
}
func (this *Day02Fixture) Test_OneLine_OneCharacter_Right() {
	buttons := this.security.PushButtons("R")
	this.So(buttons, should.Equal, 6)
}

func (this *Day02Fixture) Test_OneLine_TwoCharacters_UpLeft() {
	buttons := this.security.PushButtons("UL")
	this.So(buttons, should.Equal, 5)
}
func (this *Day02Fixture) Test_OneLine_TwoCharacters_LeftUp() {
	buttons := this.security.PushButtons("LU")
	this.So(buttons, should.Equal, 5)
}
func (this *Day02Fixture) Test_OneLine_TwoCharacters_DownLeft() {
	buttons := this.security.PushButtons("DL")
	this.So(buttons, should.Equal, 5)
}
func (this *Day02Fixture) Test_OneLine_TwoCharacters_RightUp() {
	buttons := this.security.PushButtons("RU")
	this.So(buttons, should.Equal, 2)
}
func (this *Day02Fixture) Test_OneLine_TwoCharacters_RightDown() {
	buttons := this.security.PushButtons("RD")
	this.So(buttons, should.Equal, A)
}

func (this *Day02Fixture) Test_ManyLines_ManyCharacters() {
	var buttons int

	buttons = this.security.PushButtons("ULL")
	this.So(buttons, should.Equal, 5)

	buttons = this.security.PushButtons("RRDDD")
	this.So(buttons, should.Equal, D)

	buttons = this.security.PushButtons("LURDL")
	this.So(buttons, should.Equal, B)

	buttons = this.security.PushButtons("UUUUD")
	this.So(buttons, should.Equal, 3)
}

func (this *Day02Fixture) Test_TheWholeKeypadInOrder() {
	var buttons int

	buttons = this.security.PushButtons("RURU")
	this.So(buttons, should.Equal, 1)

	buttons = this.security.PushButtons("DL")
	this.So(buttons, should.Equal, 2)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 3)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 4)

	buttons = this.security.PushButtons("DLLL")
	this.So(buttons, should.Equal, 5)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 6)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 7)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 8)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 9)

	buttons = this.security.PushButtons("LLLD")
	this.So(buttons, should.Equal, A)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, B)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, C)

	buttons = this.security.PushButtons("LD")
	this.So(buttons, should.Equal, D)
}

func (this *Day02Fixture) Test_ThePuzzleInput() {
	var buttons int

	buttons = this.security.PushButtons("RDLRUUULRRDLRLLRLDDUDLULULDDULUDRRUURLRLLUULDURRULLRULDRRDLLULLRLLDRLDDRRRRLLRLURRRDRDULRDUULDDDULURUDDRRRUULUDRLLUUURLUDRUUUDRDUULLRLLUDDRURRDDDRDLUUURLRLLUDRURDUDUULDDLLRDURULLLURLDURLUUULDULDDULULLLRRUDLRUURDRDLLURLUDULDUUUURRLDLUDRULUDLDLLDRLDDDRRLLDUDLLRRDDDRLUDURLLLDRUDDLDDRRLUDRRDUDLRRLULDULURULDULUULDRLLDRUUDDRLLUDRULLRRRLRDLRLUDLRULDRDLRDRLRULUDUURRUUULLDDDDUDDLDDDDRRULRDLRDDULLDLDLLDLLDLLDRRDDDRDDLRRDDDRLLLLURRDLRRLDRURDDURDULDDRUURUDUDDDRDRDDRLRRLRULLDRLDLURLRLRUDURRRDLLLUDRLRDLLDDDLLUDRLDRRUUDUUDULDULLRDLUDUURLDDRUDR")
	this.So(buttons, should.Equal, C)

	buttons = this.security.PushButtons("URULDDLDDUDLLURLUUUUUULUDRRRDDUDURDRUURLLDRURLUULUDRDRLLDRLDULRULUURUURRLRRDRUUUDLLLLRUDDLRDLLDUDLLRRURURRRUDLRLRLLRULRLRLRDLRLLRRUDDRLRUDULDURDLDLLLRDRURURRULLLDLLRRDRLLDUUDLRUUDDURLLLDUUDLRDDURRDRRULLDRLRDULRRLLRLLLLUDDDRDRULRRULLRRUUDULRRRUDLLUUURDUDLLLURRDDUDLDLRLURDDRRRULRRUDRDRDULURULRUDULRRRLRUDLDDDDRUULURDRRDUDLULLRUDDRRRLUDLRURUURDLDURRDUUULUURRDULLURLRUUUUULULLDRURULDURDDRRUDLRLRRLLLLDDUURRULLURURRLLDRRDDUUDLLUURRDRLLLLRLUDUUUDLRLRRLDURDRURLRLRULURLDULLLRRUUUDLLRRDUUULULDLLDLRRRDUDDLRULLULLULLULRU")
	this.So(buttons, should.Equal, 1)

	buttons = this.security.PushButtons("DURUUDULRRLULLLDDUDDLRRDURURRRDDRRURDRURDRLULDUDUDUULULDDUURDDULRDUDUDRRURDRDDRLDRDRLDULDDULRULLDULURLUUDUDULRDDRRLURLLRRDLLDLDURULUDUDULDRLLRRRUDRRDDDRRDRUUURLDLURDLRLLDUULLRULLDDDDRULRRLRDLDLRLUURUUULRDUURURLRUDRDDDRRLLRLLDLRULUULULRUDLUDULDLRDDDDDRURDRLRDULRRULRDURDDRRUDRUDLUDLDLRUDLDDRUUULULUULUUUDUULDRRLDUDRRDDLRUULURLRLULRURDDLLULLURLUDLULRLRRDDDDDRLURURURDRURRLLLLURLDDURLLURDULURUUDLURUURDLUUULLLLLRRDUDLLDLUUDURRRURRUUUDRULDDLURUDDRRRDRDULURURLLDULLRDDDRRLLRRRDRLUDURRDLLLLDDDDLUUURDDDDDDLURRURLLLUURRUDLRLRRRURULDRRLULD")
	this.So(buttons, should.Equal, A)

	buttons = this.security.PushButtons("LLUUURRDUUDRRLDLRUDUDRLRDLLRDLLDRUULLURLRRLLUDRLDDDLLLRRRUDULDLLLDRLURDRLRRLURUDULLRULLLURRRRRDDDLULURUDLDUDULRRLUDDURRLULRRRDDUULRURRUULUURDRLLLDDRDDLRRULRDRDRLRURULDULRRDRLDRLLDRDURUUULDLLLRDRRRLRDLLUDRDRLURUURDLRDURRLUDRUDLURDRURLRDLULDURDDURUUDRLULLRLRLDDUDLLUUUURLRLRDRLRRRURLRULDULLLLDLRRRULLUUDLDURUUUDLULULRUDDLLDLDLRLDDUDURDRLLRRLRRDDUDRRRURDLRLUUURDULDLURULUDULRRLDUDLDDDUUDRDUULLDDRLRLLRLLLLURDDRURLDDULLULURLRDUDRDDURLLLUDLLLLLUDRDRDLURRDLUDDLDLLDDLUDRRDDLULRUURDRULDDDLLRLDRULURLRURRDDDRLUUDUDRLRRUDDLRDLDULULDDUDURRRURULRDDDUUDULLULDDRDUDRRDRDRDLRRDURURRRRURULLLRRLR")
	this.So(buttons, should.Equal, 8)

	buttons = this.security.PushButtons("URLULLLDRDDULRRLRLUULDRUUULDRRLLDDDLDUULLDRLULRRDRRDDDRRDLRRLLDDRDULLRRLLUDUDDLDRDRLRDLRDRDDUUDRLLRLULLULRDRDDLDDDRLURRLRRDLUDLDDDLRDLDLLULDDRRDRRRULRUUDUULDLRRURRLLDRDRRDDDURUDRURLUDDDDDDLLRLURULURUURDDUDRLDRDRLUUUULURRRRDRDULRDDDDRDLLULRURLLRDULLUUDULULLLLRDRLLRRRLLRUDUUUULDDRULUDDDRRRULUDURRLLDURRDULUDRUDDRUURURURLRDULURDDDLURRDLDDLRUDUUDULLURURDLDURRDRDDDLRRDLLULUDDDRDLDRDRRDRURRDUDRUURLRDDUUDLURRLDRRDLUDRDLURUDLLRRDUURDUDLUDRRL")
	this.So(buttons, should.Equal, 8)
}

// http://adventofcode.com/2016/day/2

// --- Day 2: Bathroom Security ---

// You arrive at Easter Bunny Headquarters under cover of darkness. However, you left in such a rush that you forgot to
// use the bathroom! Fancy office buildings like this one usually have keypad locks on their bathrooms, so you search
// the front desk for the code.

// "In order to improve security," the document you find says, "bathroom codes will no longer be written down. Instead,
// please memorize and follow the procedure below to access the bathrooms."

// The document goes on to explain that each button to be pressed can be found by starting on the previous button and
// moving to adjacent buttons on the keypad: U moves up, D moves down, L moves left, and R moves right. Each line of
// instructions corresponds to one button, starting at the previous button (or, for the first line, the "5" button);
// press whatever button you're on at the end of each line. If a move doesn't lead to a button, ignore it.

// You can't hold it much longer, so you decide to figure out the code as you walk to the bathroom. You picture a
// keypad like this:

// 1 2 3
// 4 5 6
// 7 8 9
// Suppose your instructions are:

// ULL
// RRDDD
// LURDL
// UUUUD

// You start at "5" and move up (to "2"), left (to "1"), and left (you can't, and stay on "1"), so the first button
// is 1.

// Starting from the previous button ("1"), you move right twice (to "3") and then down three times (stopping at "9"
// after two moves and ignoring the third), ending up with 9.

// Continuing from "9", you move left, up, right, down, and left, ending with 8.

// Finally, you move up four times (stopping at "2"), then down once, ending with 5.

// So, in this example, the bathroom code is 1985.

// Your puzzle input is the instructions from the document you found at the front desk. What is the bathroom code?

// Puzzle Input:

// RDLRUUULRRDLRLLRLDDUDLULULDDULUDRRUURLRLLUULDURRULLRULDRRDLLULLRLLDRLDDRRRRLLRLURRRDRDULRDUULDDDULURUDDRRRUULUDRLLUUURLUDRUUUDRDUULLRLLUDDRURRDDDRDLUUURLRLLUDRURDUDUULDDLLRDURULLLURLDURLUUULDULDDULULLLRRUDLRUURDRDLLURLUDULDUUUURRLDLUDRULUDLDLLDRLDDDRRLLDUDLLRRDDDRLUDURLLLDRUDDLDDRRLUDRRDUDLRRLULDULURULDULUULDRLLDRUUDDRLLUDRULLRRRLRDLRLUDLRULDRDLRDRLRULUDUURRUUULLDDDDUDDLDDDDRRULRDLRDDULLDLDLLDLLDLLDRRDDDRDDLRRDDDRLLLLURRDLRRLDRURDDURDULDDRUURUDUDDDRDRDDRLRRLRULLDRLDLURLRLRUDURRRDLLLUDRLRDLLDDDLLUDRLDRRUUDUUDULDULLRDLUDUURLDDRUDR
// URULDDLDDUDLLURLUUUUUULUDRRRDDUDURDRUURLLDRURLUULUDRDRLLDRLDULRULUURUURRLRRDRUUUDLLLLRUDDLRDLLDUDLLRRURURRRUDLRLRLLRULRLRLRDLRLLRRUDDRLRUDULDURDLDLLLRDRURURRULLLDLLRRDRLLDUUDLRUUDDURLLLDUUDLRDDURRDRRULLDRLRDULRRLLRLLLLUDDDRDRULRRULLRRUUDULRRRUDLLUUURDUDLLLURRDDUDLDLRLURDDRRRULRRUDRDRDULURULRUDULRRRLRUDLDDDDRUULURDRRDUDLULLRUDDRRRLUDLRURUURDLDURRDUUULUURRDULLURLRUUUUULULLDRURULDURDDRRUDLRLRRLLLLDDUURRULLURURRLLDRRDDUUDLLUURRDRLLLLRLUDUUUDLRLRRLDURDRURLRLRULURLDULLLRRUUUDLLRRDUUULULDLLDLRRRDUDDLRULLULLULLULRU
// DURUUDULRRLULLLDDUDDLRRDURURRRDDRRURDRURDRLULDUDUDUULULDDUURDDULRDUDUDRRURDRDDRLDRDRLDULDDULRULLDULURLUUDUDULRDDRRLURLLRRDLLDLDURULUDUDULDRLLRRRUDRRDDDRRDRUUURLDLURDLRLLDUULLRULLDDDDRULRRLRDLDLRLUURUUULRDUURURLRUDRDDDRRLLRLLDLRULUULULRUDLUDULDLRDDDDDRURDRLRDULRRULRDURDDRRUDRUDLUDLDLRUDLDDRUUULULUULUUUDUULDRRLDUDRRDDLRUULURLRLULRURDDLLULLURLUDLULRLRRDDDDDRLURURURDRURRLLLLURLDDURLLURDULURUUDLURUURDLUUULLLLLRRDUDLLDLUUDURRRURRUUUDRULDDLURUDDRRRDRDULURURLLDULLRDDDRRLLRRRDRLUDURRDLLLLDDDDLUUURDDDDDDLURRURLLLUURRUDLRLRRRURULDRRLULD
// LLUUURRDUUDRRLDLRUDUDRLRDLLRDLLDRUULLURLRRLLUDRLDDDLLLRRRUDULDLLLDRLURDRLRRLURUDULLRULLLURRRRRDDDLULURUDLDUDULRRLUDDURRLULRRRDDUULRURRUULUURDRLLLDDRDDLRRULRDRDRLRURULDULRRDRLDRLLDRDURUUULDLLLRDRRRLRDLLUDRDRLURUURDLRDURRLUDRUDLURDRURLRDLULDURDDURUUDRLULLRLRLDDUDLLUUUURLRLRDRLRRRURLRULDULLLLDLRRRULLUUDLDURUUUDLULULRUDDLLDLDLRLDDUDURDRLLRRLRRDDUDRRRURDLRLUUURDULDLURULUDULRRLDUDLDDDUUDRDUULLDDRLRLLRLLLLURDDRURLDDULLULURLRDUDRDDURLLLUDLLLLLUDRDRDLURRDLUDDLDLLDDLUDRRDDLULRUURDRULDDDLLRLDRULURLRURRDDDRLUUDUDRLRRUDDLRDLDULULDDUDURRRURULRDDDUUDULLULDDRDUDRRDRDRDLRRDURURRRRURULLLRRLR
// URLULLLDRDDULRRLRLUULDRUUULDRRLLDDDLDUULLDRLULRRDRRDDDRRDLRRLLDDRDULLRRLLUDUDDLDRDRLRDLRDRDDUUDRLLRLULLULRDRDDLDDDRLURRLRRDLUDLDDDLRDLDLLULDDRRDRRRULRUUDUULDLRRURRLLDRDRRDDDURUDRURLUDDDDDDLLRLURULURUURDDUDRLDRDRLUUUULURRRRDRDULRDDDDRDLLULRURLLRDULLUUDULULLLLRDRLLRRRLLRUDUUUULDDRULUDDDRRRULUDURRLLDURRDULUDRUDDRUURURURLRDULURDDDLURRDLDDLRUDUUDULLURURDLDURRDRDDDLRRDLLULUDDDRDLDRDRRDRURRDUDRUURLRDDUUDLURRLDRRDLUDRDLURUDLLRRDUURDUDLUDRRL

// --- Part Two ---
//
// You finally arrive at the bathroom (it's a several minute walk from the lobby so visitors can behold the many fancy
// conference rooms and water coolers on this floor) and go to punch in the code. Much to your bladder's dismay, the
// keypad is not at all like you imagined it. Instead, you are confronted with the result of hundreds of man-hours of
// bathroom-keypad-design meetings:
//
//     1
//   2 3 4
// 5 6 7 8 9
//   A B C
//     D
//
// You still start at "5" and stop when you're at an edge, but given the same instructions as above, the outcome is
// very different:
//
// - You start at "5" and don't move at all (up and left are both edges), ending at 5.
// - Continuing from "5", you move right twice and down three times (through "6", "7", "B", "D", "D"), ending at D.
// - Then, from "D", you move five more times (through "D", "B", "C", "C", "B"), ending at B.
// - Finally, after five more moves, you end at 3.
//
// So, given the actual keypad layout, the code would be 5DB3.
//
// Using the same instructions in your puzzle input, what is the correct bathroom code?
