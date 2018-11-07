package board_test

import (
	. "./"
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func TestBoard(t *testing.T) { TestingT(t) }

type BoardSuite struct {
	board *Board
}

var _ = Suite(&BoardSuite{})

func (s *BoardSuite) TestResolveHandsResult(c *C) {
	h:=NewHands("6s5h")
	b:=NewBoard("KhQhThTc7s")
	hresult:=b.ResolveHandsResult(h)
	c.Assert(hresult.Value(),Equals,107131211)
}

func (s *BoardSuite) TestAllCardLevel(c *C) {
	allLevel:=map[int]int{0:0,1:0,2:0,3:0,4:0,5:0,6:0,7:0,8:0}
	for len(allLevel)>0  {
		b:=NewDefaultBoard()
		h:=b.DealHandsCards(1)
		b.DealShowCards(5)
		r:=b.ResolveHandsResult(h[0])
		_,ok:=allLevel[r.Level()]
		if ok {
			fmt.Println(r.String())
			delete(allLevel,r.Level())
		}
	}
}

func (s *BoardSuite) TestDealCard(c *C) {
	b:=NewDefaultBoard()
	hlist:=b.DealHandsCards(5)
	fmt.Println(hlist)
	b.DealShowCards(5)
	for _,h:=range hlist {
		hresult:=b.ResolveHandsResult(h)
		fmt.Println("3"+hresult.String())
	}
}