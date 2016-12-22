package day_03

type Triangles struct{}

func NewTriangles() *Triangles {
	return &Triangles{}
}

func (this *Triangles) HowManyPossible(triangles [][3]int) int {
	validTriangles := 0
	for _, sides := range triangles {
		if sides[0]+sides[1] > sides[2] && sides[1]+sides[2] > sides[0] && sides[2]+sides[0] > sides[1] {
			validTriangles = validTriangles + 1
		}
	}
	return validTriangles
}
