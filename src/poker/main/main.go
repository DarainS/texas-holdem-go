package main

import (
	"../board"
	"fmt"
)

func main() {
	h:= board.Hands{}.FromString("AsTs")
	board:=board.Board{}.FromString(h,"JsQsKs")
	fmt.Println(board.ShowList)
	fmt.Println(board.ResolveValue(h))
}
