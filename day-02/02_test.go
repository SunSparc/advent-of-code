package Day2

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)


type Day2Fixture struct {
	*gunit.Fixture
	security *BathroomSecurity
}

func TestDay2Fixture(t *testing.T) {
	gunit.Run(new(Day2Fixture), t)
}

func (this *Day2Fixture) Setup() {
	this.security = NewBathroomSecurity()
}

var buttons int

func (this *Day2Fixture) Test_OneLine_OneCharacterInstructions() {
	buttons = this.security.PushButtons("")
	this.So(buttons, should.Equal, 0)

	buttons = this.security.PushButtons("U")
	this.So(buttons, should.Equal, 2)

	buttons = this.security.PushButtons("D")
	this.So(buttons, should.Equal, 8)

	buttons = this.security.PushButtons("L")
	this.So(buttons, should.Equal, 4)

	buttons = this.security.PushButtons("R")
	this.So(buttons, should.Equal, 6)
}

func (this *Day2Fixture) Test_OneLine_TwoCharacterInstructions() {
	buttons = this.security.PushButtons("UL")
	this.So(buttons, should.Equal, 1)
	buttons = this.security.PushButtons("LU")
	this.So(buttons, should.Equal, 1)

	buttons = this.security.PushButtons("DL")
	this.So(buttons, should.Equal, 7)

	buttons = this.security.PushButtons("RU")
	this.So(buttons, should.Equal, 3)

	buttons = this.security.PushButtons("RD")
	this.So(buttons, should.Equal, 9)
}

// -------------- Design ---------------
// number pad
// button pusher
// button pusher mover

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
