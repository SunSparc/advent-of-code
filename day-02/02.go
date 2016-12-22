package day02


type Keypad [3][3]int

type BathroomSecurity struct {
	keypad Keypad
	start int
	current int
}

func NewBathroomSecurity() *BathroomSecurity {
	keypad := Keypad{
		[3]int{1,2,3},
		[3]int{4,5,6},
		[3]int{7,8,9},
	}
	return &BathroomSecurity{
		keypad: keypad,
		start: 5,
		current: 5,
	}
}


func (this *BathroomSecurity) PushButtons(instructions string) int {
	if instructions == "U" {
		this.up()
		return this.current
	} else if instructions == "D" {
		this.down()
		return this.current
	} else if instructions == "L" {
		this.left()
		return this.current
	} else if instructions == "R" {
		this.right()
		return this.current
	} else if instructions == "UL" {
		this.up()
		this.left()
		return this.current
	} else if instructions == "LU" {
		this.left()
		this.up()
		return this.current
	} else if instructions == "DL" {
		this.down()
		this.left()
		return this.current
	} else if instructions == "RU" {
		this.right()
		this.up()
		return this.current
	} else if instructions == "RD" {
		this.right()
		this.down()
		return 9
	}
	return 0
}

func (this *BathroomSecurity) up() {
	current := this.current
	new := current-3
	if  new > 0 {
		this.current = new
	}
}

func (this *BathroomSecurity) down() {
	current := this.current
	new := current+3
	if new <= 9 {
		this.current = new
	}
}

func (this *BathroomSecurity) left() {
	current := this.current
	if current == 1 || current == 4 || current == 7 {
		return
	}
	this.current = current-1
}

func (this *BathroomSecurity) right() {
	current := this.current
	if current == 3 || current == 6 || current == 9 {
		return
	}
	this.current = current+1
}
