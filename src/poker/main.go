package main

import "fmt"
import . "./base"
import . "./statistics"

func main() {
	board := NewDefaultBoard()
	handsList := board.DealHandsCards(3)
	board.DealShowCards(5)
	rlist := FindHandsOver(handsList[0], board)
	for _, r := range rlist {
		fmt.Println(r.String())
	}
}
