package board

import (
	"fmt"
	"sort"
	"strconv"
)

type Board struct {
	hand     Hands
	Deck     Deck
	ShowList []Card
}

func Test() {
	h := NewHands("AsTs")
	board := NewBoard(h, "JsQsKs")
	//fmt.Println(board.ShowList)
	handsResult := board.ResolveHandsValue(h)
	fmt.Println(handsResult.String())
}

type HandsResult struct {
	hands     Hands
	cards     []Card
	value     int64
	level     int
	levelText string
}

var (
	LevelMap = map[int]string{
		0: "高牌",
		1: "一对",
		2: "两对",
		3: "三条",
		4: "顺子",
		5: "同花",
		6: "葫芦",
		7: "四条",
		8: "同花顺",
	}
)

func (r *HandsResult) String() string {
	return r.levelText + " " + CardsToString(r.cards) + " " + strconv.FormatInt(r.value, 10)
}

func CardsToString(cards []Card) string {
	res := ""
	for _, card := range cards {
		res += card.String()
	}
	return res
}

func NewBoard(h Hands, s string) Board {
	board := Board{}
	board.hand = h
	for i := 0; i < len(s); i += 2 {
		board.ShowList = append(board.ShowList, NewCard(s[i:i+2]))
	}
	board.Deck = NewDeck()
	return board
}

func (board *Board) DealHands(playerNum int) []Hands {
	handsList := make([]Hands, playerNum)
	for i := 0; i < cap(handsList); i++ {
		h := Hands{}
		h[0] = board.Deck.DealOne()
		h[1] = board.Deck.DealOne()
		handsList[i] = h
	}
	return handsList
}

func ResolveValue(cards []Card) int64 {
	SortCards(cards)
	numMap := generateCardNumMap(cards)
	tagMap := generateTagNumMap(cards)
	var value int64
	if value = resolveStraightFlushAndFlush(cards, numMap, tagMap); value > 0 {
		return value
	}
	if value := resolveQuads(cards, numMap); value > 0 {
		return value
	}
	if value := resolveFullHouse(cards, numMap); value > 0 {
		return value
	}
	if value := resolveStraight(cards, numMap); value > 0 {
		return value
	}
	if value := resolveSet(cards, numMap); value > 0 {
		return value
	}
	if value = resolvePair(cards, numMap); value > 0 {
		return value
	}
	if value = resolveHigh(cards); value > 0 {
		return value
	}
	return -1
}

func (board *Board) ResolveHandsValue(h Hands) HandsResult {
	cards := board.generateTempCardList(h)
	res := HandsResult{}
	res.hands = h
	res.cards = cards
	res.value = ResolveValue(cards)
	s := strconv.FormatInt(res.value, 10)
	if len(s) == 11 {
		res.level = int(s[0]) - int('0')
	} else {
		res.level = 0
	}
	res.levelText = LevelMap[res.level]
	return res
}

func resolveStraightFlushAndFlush(cards []Card, numMap map[int]int, tagMap map[uint8]int) int64 {
	tag := testFlush(tagMap)
	if tag == 0 {
		return 0
	}
	straightCards := make([]Card, 7)
	for _, card := range cards {
		if card.tag == tag {
			straightCards = append(straightCards, card)
		}
	}
	SortCards(straightCards)
	if num := testStraight(straightCards, generateCardNumMap(straightCards)); num > 0 {
		return calculateValue(8, []int{num, num - 1, num - 2, num - 3, num - 4})
	} else {
		return resolveFlush(cards, tagMap)
	}
	return 0
}

func resolveQuads(cards []Card, numMap map[int]int) int64 {
	f, h := 0, 0
	for i := 14; i >= 2; i-- {
		if numMap[i] >= 4 && f == 0 {
			f = i
		}
		if numMap[i] >= 1 && f != i {
			h = i
			break
		}
	}
	if f == 0 || h == 0 {
		return 0
	}
	return calculateValue(7, []int{f, h})
}

func resolveFullHouse(cards []Card, numMap map[int]int) int64 {
	f, h := 0, 0
	for i := 14; i >= 2; i-- {
		if numMap[i] >= 3 && f == 0 {
			f = i
		}
		if numMap[i] >= 2 && f != i {
			h = i
			break
		}
	}
	if f == 0 || h == 0 {
		return 0
	}
	return calculateValue(6, []int{f, h})
}

func testFlush(tagMap map[uint8]int) uint8 {
	for tag, num := range tagMap {
		if num >= 5 {
			return tag
		}
	}
	return 0
}

func resolveFlush(cardList []Card, tagMap map[uint8]int) int64 {
	var tag uint8
	for tag2, num := range tagMap {
		if num >= 5 {
			tag = tag2
			break
		}
	}
	if tag == 0 {
		return 0
	}
	res := int64(5)
	for _, card := range cardList[0:5] {
		res = res*100 + int64(card.num)
	}
	return res
}

func testStraight(cards []Card, numMap map[int]int) int {
	for i := 14; i >= 5; i-- {
		isStraight := true
		for j := 0; i <= 4; j++ {
			if !(numMap[i-j] >= 1) {
				isStraight = false
				break
			}
		}
		if isStraight {
			return i
		}
	}
	return 0
}

func resolveStraight(cards []Card, numMap map[int]int) int64 {
	if i := testStraight(cards, numMap); i > 0 {
		return calculateValue(4, []int{i, i - 1, i - 2, i - 3, i - 4})
	}
	return 0
}

func resolveSet(cards []Card, numMap map[int]int) int64 {
	f, h1, h2 := 0, 0, 0
	for i := 14; i >= 2; i-- {
		if numMap[i] >= 3 {
			f = i
			break
		}
	}
	if f == 0 {
		return 0
	}
	for i := 14; i >= 2; i-- {
		if numMap[i] >= 1 {
			if f != i {
				if h1 == 0 {
					h1 = i
				} else {
					h2 = i
					break
				}
			}
		}
	}
	return calculateValue(3, []int{f, h1, h2})
}

func resolvePair(cards []Card, numMap map[int]int) int64 {
	p1, p2, t := 0, 0, 0
	for i := 14; i >= 2; i-- {
		if numMap[i] == 2 {
			if p1 == 0 {
				p1 = i
			} else {
				p2 = i
				break
			}
		}
	}
	if p1 == 0 {
		return 0
	}
	ticker := make([]int, 3)
	if p2 == 0 {
		for i := 14; i >= 2; i-- {
			if len(ticker) == 3 {
				break
			}
			if numMap[i] == 1 {
				ticker = append(ticker, i)
			}
		}
		return calculateValue(1, ticker)
	}

	if p2 != 0 {
		for i := 14; i >= 2; i-- {
			if numMap[i] == 1 {
				t = i
				break
			}
		}
		return calculateValue(2, []int{t})
	}
	return 0
}

func resolveHigh(cards []Card) int64 {
	res := []int{0, 0, 0, 0, 0}
	for index, card := range cards[0:5] {
		res[index] = card.num
	}
	return calculateValue(0, res)
}

func calculateValue(level int, nums []int) int64 {
	result := int64(level)
	for i := 0; i < 5; i++ {
		if i < len(nums) {
			result = result*100 + int64(nums[i])
		} else {
			result *= 100
		}
	}
	return result
}

func SortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].num > cards[j].num
	})
}

func (board Board) generateTempCardList(h Hands) []Card {
	result := append(board.ShowList, h[0], h[1])
	SortCards(result)
	return result
}

func generateCardNumMap(cards []Card) map[int]int {
	numMap := make(map[int]int)
	for i := 1; i <= 14; i++ {
		numMap[i] = 0
	}
	for _, card := range cards {
		numMap[card.Num()] += 1
		if card.num == 14 {
			numMap[1] += 1
		}
	}
	return numMap
}

func generateTagNumMap(cards []Card) map[uint8]int {
	tagMap := map[uint8]int{
		'c': 0, 's': 0, 'h': 0, 'd': 0,
	}
	for _, card := range cards {
		tagMap[card.tag] += 1
	}
	return tagMap
}

func (board *Board) calculateValue(value ...int) int64 {
	res := int64(0)
	res = int64(value[0])
	for _, v := range value {
		res = res*100 + int64(v)
	}
	for i := len(value); i < 5; i++ {
		res *= 100
	}
	return res
}
