package board

import (
	"math/rand"
	"time"
)

type Deck []Card

func (deck *Deck) DealOne() Card {
	c := (*deck)[0]
	*deck = (*deck)[1:]
	return c
}

var (
	BaseDeck Deck
)

func init() {
	BaseDeck = Deck(AllCards)
}

func NewDeck() Deck {
	deck := make([]Card, 52)
	copy(deck, BaseDeck)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, j := range r.Perm(len(deck)) {
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}
