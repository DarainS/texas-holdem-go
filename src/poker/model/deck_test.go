package board_test

import (
	. "./"
	"testing"
	. "gopkg.in/check.v1"
	"fmt"
)


func Test(t *testing.T) {
	TestingT(t)
}

type MySuite Board

var _ = Suite(&MySuite{})

func (s *MySuite) TestFindHandsOver(c *C) {

}
