package day_03

type Triangles struct {
	list          [][3]int
	listMark      int
	currentColumn int
}

func NewTriangles(listOfTriangles [][3]int) *Triangles {
	return &Triangles{
		list:          listOfTriangles,
		listMark:      0,
		currentColumn: 0,
	}
}

func (this *Triangles) HowManyPossible() int {
	validTriangles := 0
	for sides := this.readNextTriplet(); sides != [3]int{0, 0, 0}; sides = this.readNextTriplet() {
		if sides[0]+sides[1] > sides[2] && sides[1]+sides[2] > sides[0] && sides[2]+sides[0] > sides[1] {
			validTriangles = validTriangles + 1
		}
	}
	return validTriangles
}

func (this *Triangles) readNextTriplet() [3]int {
	var nextTriplet [3]int

	keepGoing := true
	if this.listMark >= len(this.list) {
		keepGoing = this.incrementColumn()
	}
	if keepGoing {
		nextTriplet[0] = this.list[this.listMark][this.currentColumn]
		nextTriplet[1] = this.list[this.listMark+1][this.currentColumn]
		nextTriplet[2] = this.list[this.listMark+2][this.currentColumn]
		this.listMark = this.listMark + 3
	}

	return nextTriplet

}

func (this *Triangles) incrementColumn() bool {
	if this.currentColumn < 2 {
		this.currentColumn = this.currentColumn + 1
		this.listMark = 0
		return true
	}
	return false
}
