package AdventofCode2016

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestAdventofCode2016Fixture(t *testing.T) {
	gunit.Run(new(AdventofCode2016Fixture), t)
}

type AdventofCode2016Fixture struct {
	*gunit.Fixture
}

// For example:

// Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
// R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
// R5, L5, R5, R3 leaves you 12 blocks away.

// How many blocks away is Easter Bunny HQ?

func (this *AdventofCode2016Fixture) TestFirst() {
	// do something
	this.So(1-1, should.Equal, 0)
	// directions := []string{"R2", "L3"} // should.Equal 5 block
	// directions := []string{"R2", "R2", "R2"} // should.Equal 2 blocks
	// directions := []string{"R5", "L5", "R5", "R3"} // should.Equal 12 blocks

}
