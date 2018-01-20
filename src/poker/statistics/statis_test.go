package statistics_test

import (
	"testing"
	. "../base"
	. "../statistics"
	. "gopkg.in/check.v1"
	"fmt"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	board *Board
}

var _ = Suite(&MySuite{})

func (s *MySuite) TestFindHandsOver(c *C) {
	s.board = NewDefaultBoard()
	handsList := s.board.DealHandsCards(3)
	s.board.DealShowCards(5)
	rlist := FindHandsOver(handsList[0], s.board)
	for _, r := range rlist {
		fmt.Println(r.String())
	}
}
