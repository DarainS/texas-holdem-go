package board

type Card struct {
	tag      uint8
	num      int
	symbol   uint8
	IsActive bool
}

var SymbolTable = map[int]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}
var TagTable = map[int]int{
	'c': 'c', 's': 's', 'h': 'h', 'd': 'd',
}
var (
	AllCardsString []string        = make([]string, 0, 52)
	AllCards       []Card          = make([]Card, 0, 52)
	AllCardsTable  map[string]Card = make(map[string]Card)
)

func init() {
	for symbol := range SymbolTable {
		for tag := range TagTable {
			s := string(symbol) + string(tag)
			c := NewCardFromString(s)
			AllCardsString = append(AllCardsString, s)
			AllCards = append(AllCards, c)
			AllCardsTable[s] = c
		}
	}
}

func NewCardFromString(s string) Card {
	c := Card{}
	c.symbol = s[0]
	c.tag = s[1]
	c.num = SymbolTable[int(c.symbol)]
	return c
}


func (c Card) String() string {
	return string(c.symbol) + string(c.tag)
}

func (c Card) Less(that Card) bool {
	return c.num < that.num
}

func (c Card) Tag() uint8 {
	return c.tag
}

func (c Card) Num() int {
	return c.num
}

func (c Card) Symbol() uint8 {
	return c.symbol
}

type Hands []Card

func NewHandsFromTwoCard(c1, c2 Card) Hands {
	h := make(Hands, 2)
	h[0] = c1
	h[1] = c2
	//SortCards([2]Card(h))
	return h
}

func NewHands(s string) Hands {
	h := make(Hands, 2)
	if len(s) == 4 {
		c1 := NewCardFromString(s[0:2])
		c2 := NewCardFromString(s[2:4])
		if c1.num >= c2.num {
			h[0], h[1] = c1, c2
		} else {
			h[0], h[1] = c2, c1
		}
	}
	return h
}

func (h Hands) String() string {
	return h[0].String() + h[1].String()
}

func (h Hands) ToSingleString() string {
	s1 := string(h[0].symbol + h[1].symbol)
	if h[0].num == h[1].num {
		return s1
	} else if h[0].tag == h[1].tag {
		return s1 + "s"
	}
	return s1 + "o"
}
func CardsToString(cards []Card) string {
	res := ""
	for _, card := range cards {
		res += card.String()
	}
	return res
}

