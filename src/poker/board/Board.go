package board

import "sort"

type Board struct {
	HandsList []Hands
	Deck      Deck
	ShowList  []Card
}

func (board Board) FromString(h Hands, s string) Board {
	for i := 0; i < len(s); i += 2 {
		board.ShowList = append(board.ShowList, Card{}.FromString(s[i:i+2]))
	}
	return board
}

func (board Board) ResolveValue(h Hands) int64 {
	cards := board.generateTempCardList(h)
	value := board.resolveFlush(cards)
	return value
}

func (board Board) resolvePair(h Hands) {

}

func (board Board) generateTempCardList(h Hands) []Card {
	result := append(board.ShowList, h[0], h[1])
	sort.Slice(result, func(i, j int) bool {
		return result[i].Num > result[j].Num
	})
	return result
}

func (board Board) resolveFlush(cardList []Card) int64 {
	tagMap := map[uint8]int{
		'c': 0, 's': 0, 'h': 0, 'd': 0,
	}
	for _, card := range cardList {
		tagMap[card.Tag] += 1
	}
	var tag uint8
	for tag2, num := range tagMap {
		if num >= 5 {
			tag = tag2
		}
	}
	if tag == 0 {
		return 0
	}
	res := int64(5)
	for _, card := range cardList {
		res = res*100 + int64(card.Num)
	}
	return res
}
