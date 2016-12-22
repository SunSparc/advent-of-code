package day02


type BathroomSecurity struct {
	start int
	current int
}

func NewBathroomSecurity() *BathroomSecurity {
	return &BathroomSecurity{
		start: 5,
		current: 5,
	}
}


func (this *BathroomSecurity) PushButtons(instructions string) int {
	if instructions == "" {
		return 0
	}
	for _, direction := range instructions {
		switch string(direction) {
		case "U":
			this.up()
		case "D":
			this.down()
		case "L":
			this.left()
		case "R":
			this.right()
		}
	}
	return this.current
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
