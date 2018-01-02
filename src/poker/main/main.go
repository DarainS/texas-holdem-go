package main

import (
	"../board"
	"fmt"
)

func main() {
	h := board.NewHands("AsTs")
	board := board.NewBoard(h, "JsQsKs")
	//fmt.Println(board.ShowList)
	r := board.ResolveHandsResult(h)
	fmt.Println(r.String())
}
