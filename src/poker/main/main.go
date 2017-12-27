package main

import (
	"../board"
	"fmt"
)

func main() {
	h := board.NewHands("AsTs")
	board := board.NewBoard(h, "JsQsKs")
	//fmt.Println(board.ShowList)
	board.ResolveHandsValue(h)
	fmt.Println(board.String())
}
