package card

type Card struct {
	Tag    uint8
	Num    int
	Symbol uint8
}

var symbolTable = map[int]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}

func (c Card) FromString(s string) Card {
	c.Symbol = s[0]
	c.Tag = s[1]
	c.Num = symbolTable[int(c.Symbol)]
	return c
}

func (c Card) ToString() string {
	return string(c.Symbol) + string(c.Tag)
}

func (c Card) Less(that Card) bool {
	return c.Num<that.Num
}


type Hands [2]Card

func (h Hands) FromString(s string) Hands {
	if len(s) == 4 {
		c1 := Card{}.FromString(s[0:2])
		c2 := Card{}.FromString(s[2:4])
		if c1.Num>=c2.Num {
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
	s1 := string(h[0].Symbol + h[1].Symbol)
	if h[0].Num == h[1].Num {
		return s1
	} else if h[0].Tag == h[1].Tag {
		return s1 + "s"
	}
	return s1 + "o"
}

