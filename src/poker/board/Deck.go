package board

import (
	"math/rand"
	"time"
	"fmt"
)

type Deck []Card


func (deck Deck)DealOne()Card  {
	c:=&deck[0]
	deck=deck[1:]
	return *c
}

func NewDeck()Deck  {
	deck:=Deck{}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(deck)) {
		val := deck[i]
		fmt.Println(val)
	}
	return deck
}

