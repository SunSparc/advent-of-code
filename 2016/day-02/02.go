package day02

type BathroomSecurity struct {
	start   int
	current int
}

func NewBathroomSecurity() *BathroomSecurity {
	return &BathroomSecurity{
		start:   5,
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
	if current == 1 || current == 2 || current == 4 || current == 5 || current == 9 {
		return
	}
	switch current {
	case 3:
		this.current = 1
	case 6:
		this.current = 2
	case 7:
		this.current = 3
	case 8:
		this.current = 4
	case A:
		this.current = 6
	case B:
		this.current = 7
	case C:
		this.current = 8
	case D:
		this.current = B
	}
}

func (this *BathroomSecurity) down() {
	current := this.current
	if current == 5 || current == 9 || current == A || current == C || current == D {
		return
	}
	switch current {
	case 1:
		this.current = 3
	case 2:
		this.current = 6
	case 3:
		this.current = 7
	case 4:
		this.current = 8
	case 6:
		this.current = A
	case 7:
		this.current = B
	case 8:
		this.current = C
	case B:
		this.current = D
	}
}

func (this *BathroomSecurity) left() {
	current := this.current
	if current == 1 || current == 2 || current == 5 || current == A || current == D {
		return
	}
	switch current {
	case 3:
		this.current = 2
	case 4:
		this.current = 3
	case 6:
		this.current = 5
	case 7:
		this.current = 6
	case 8:
		this.current = 7
	case 9:
		this.current = 8
	case B:
		this.current = A
	case C:
		this.current = B
	}

}

func (this *BathroomSecurity) right() {
	current := this.current
	if current == 1 || current == 4 || current == 9 || current == C || current == D {
		return
	}
	switch current {
	case 2:
		this.current = 3
	case 3:
		this.current = 4
	case 5:
		this.current = 6
	case 6:
		this.current = 7
	case 7:
		this.current = 8
	case 8:
		this.current = 9
	case A:
		this.current = B
	case B:
		this.current = C
	}
}

const (
	A = 10
	B = 11
	C = 12
	D = 13
)
