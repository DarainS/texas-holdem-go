package card

import "sort"

type Deck struct {
	HandsList  []Hands
	ShowList   []Card
	InDeckList []Card
}

func (deck Deck)FromString(h Hands,s string)Deck {
	for i:=0;i<len(s);i+=2 {
		deck.ShowList=append(deck.ShowList,Card{}.FromString(s[i:i+2]))
	}
	return deck
}

func (deck Deck)ResolveValue(h Hands) int64 {
	cards:=deck.generateTempCardList(h)
	value:=deck.resolveFlush(cards)
	return value
}

func (deck Deck)resolvePair(h Hands) {

}

func (deck Deck)generateTempCardList(h Hands)[]Card  {
	result:=append(deck.ShowList,h[0],h[1])
	sort.Slice(result, func(i, j int) bool {
		return result[i].Num>result[j].Num
	})
	return result
}

func (deck Deck)resolveFlush(cardList []Card) int64 {
	tagMap:=map[uint8]int{
		'c':0,'s':0,'h':0,'d':0,
	}
	for _,card := range cardList {
		tagMap[card.Tag]+=1
	}
	var tag uint8
	for tag2,num := range tagMap{
		if  num>=5{
			tag=tag2
		}
	}
	if tag==0 {
		return 0
	}
	res:=int64(5)
	for _,card := range cardList{
		res=res*100+int64(card.Num)
	}
	return res
}

