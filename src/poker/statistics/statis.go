package statistics

import (
	. "../model"
	"sync"
)

var (
	ThreadNum = 4
)

func generateAllPossibleHandsFromBoard(board *Board) []Hands {
	deck := board.Deck
	cards := []Card(deck)
	l := len(cards)
	result := make([]Hands, 0, l*(l-1))
	for i1 := 0; i1 < l; i1++ {
		for i2 := i1 + 1; i2 < l; i2++ {
			h := NewHandsFromTwoCard((cards)[i1], (cards)[i2])
			result = append(result, h)
		}
	}
	return result
}

type HandsResultFilter func(result HandsResult) bool

func batchCalculate(board Board, handsList []Hands, result chan HandsResult, wg *sync.WaitGroup, filter HandsResultFilter) {
	for _, hands := range handsList {
		r := board.ResolveHandsResult(hands)
		if filter(r) {
			result <- r
		}
	}
	wg.Done()
}

func batchCalculateHandsResult(board *Board,handsList []Hands,threadNum int,filter HandsResultFilter) []HandsResult {
	wg := sync.WaitGroup{}
	totalLen:=len(handsList)
	l :=totalLen  / threadNum
	resultChannel := make(chan HandsResult, totalLen)

	for i := 0; i < threadNum-1; i++ {
		wg.Add(1)
		go batchCalculate(*board, handsList[l*i:l*(i+1)],  resultChannel, &wg,filter)
	}

	wg.Add(1)
	go batchCalculate(*board, handsList[l*(threadNum-1):], resultChannel, &wg,filter)

	wg.Wait()
	close(resultChannel)
	resultList := make([]HandsResult, 0, len(resultChannel))
	for r := range resultChannel {
		resultList = append(resultList, r)
	}
	return resultList
}

func FindHandsOver(ourHands Hands, board *Board) []HandsResult {
	baseResult := board.ResolveHandsResult(ourHands)
	allPossibleHands := generateAllPossibleHandsFromBoard(board)

	resultList:= batchCalculateHandsResult(board,allPossibleHands,4, func(result HandsResult) bool {
		return result.Value()>=baseResult.Value()
	})

	return resultList
}
