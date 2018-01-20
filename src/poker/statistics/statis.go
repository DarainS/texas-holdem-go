package statistics

import (
	base "../base"
	"sync"
)

var (
	ThredNum = 4
)

func generateAllPossibleHandsFromBoard(board *base.Board) *[]base.Hands {
	deck := board.Deck
	cards := []base.Card(deck)
	l := len(cards)
	result := make([]base.Hands, 0,l*(l-1))
	for i1 := 0; i1 < l; i1++ {
		for i2 := i1 + 1; i2 < l; i2++ {
			h := base.NewHandsFromTwoCard((cards)[i1], (cards)[i2])
			result = append(result, h)
		}
	}
	return &result
}

type HandsResultFilter func(result base.HandsResult) bool

func batchCalculate(board *base.Board, handsList []base.Hands, filter HandsResultFilter, result chan base.HandsResult, wg *sync.WaitGroup) {
	for _, hands := range handsList {
		r := board.ResolveHandsResult(hands)
		if filter(r) {
			result <- r
		}
	}
	wg.Done()
}

func FindHandsOver(ourHands base.Hands, board *base.Board) []base.HandsResult {

	allPossibleHands := *generateAllPossibleHandsFromBoard(board)

	l := len(allPossibleHands) / ThredNum

	resultChannel := make(chan base.HandsResult, len(allPossibleHands))

	r1 := board.ResolveHandsResult(ourHands)

	wg := &sync.WaitGroup{}

	for i := 0; i < ThredNum-1; i++ {
		wg.Add(1)
		go batchCalculate(board, allPossibleHands[l*i:l*(i+1)], func(result base.HandsResult) bool {
			return result.Value() >= r1.Value()
		}, resultChannel, wg)
	}

	wg.Add(1)
	go batchCalculate(board, allPossibleHands[l*(ThredNum-1):], func(result base.HandsResult) bool {
		return result.Value() >= r1.Value()
	}, resultChannel, wg)

	wg.Wait()

	close(resultChannel)

	resultList := make([]base.HandsResult,0, len(resultChannel))

	for r := range resultChannel {
		resultList = append(resultList, r)
	}

	return resultList
}
