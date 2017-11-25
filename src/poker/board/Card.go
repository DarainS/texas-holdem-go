package board

type Card struct {
	tag    uint8
	num    int
	symbol uint8
}

var symbolTable = map[int]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

func (c Card) FromString(s string) Card {
	c.symbol = s[0]
	c.tag = s[1]
	c.num = symbolTable[int(c.symbol)]
	return c
}

func (c Card) ToString() string {
	return string(c.symbol) + string(c.tag)
}

func (c Card) Less(that Card) bool {
	return c.num <that.num
}

func (c Card) Tag() uint8  {
	return c.tag
}

func (c Card) Num() int  {
	return c.num
}

func (c Card) Symbol() uint8  {
	return c.symbol
}

type Hands [2]Card

func (h Hands) FromString(s string) Hands {
	if len(s) == 4 {
		c1 := Card{}.FromString(s[0:2])
		c2 := Card{}.FromString(s[2:4])
		if c1.num >=c2.num {
			h[0],h[1]=c1,c2
		}else {
			h[0],h[1]=c2,c1
		}
	}
	return h
}

func (h Hands) ToString() string {
	return h[0].ToString() + h[1].ToString()
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

