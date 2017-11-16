package main

import (
	"../card"
	"fmt"
)

func main() {
	h := card.Hands{}.FromString("AsTs")

	deck:=card.Deck{}.FromString(h,"JsQsKs")
	fmt.Println(deck.ShowList)
	fmt.Println(deck.ResolveValue(h))
}
