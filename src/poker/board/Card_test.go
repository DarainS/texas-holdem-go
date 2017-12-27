package board

import (
	"reflect"
	"testing"
)

func TestCard_FromString(t *testing.T) {
	type fields struct {
		tag    uint8
		num    int
		symbol uint8
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Card
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				tag:    tt.fields.tag,
				num:    tt.fields.num,
				symbol: tt.fields.symbol,
			}
			if got := c.fromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Card.NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_ToString(t *testing.T) {
	type fields struct {
		tag    uint8
		num    int
		symbol uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				tag:    tt.fields.tag,
				num:    tt.fields.num,
				symbol: tt.fields.symbol,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Less(t *testing.T) {
	type fields struct {
		tag    uint8
		num    int
		symbol uint8
	}
	type args struct {
		that Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				tag:    tt.fields.tag,
				num:    tt.fields.num,
				symbol: tt.fields.symbol,
			}
			if got := c.Less(tt.args.that); got != tt.want {
				t.Errorf("Card.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Tag(t *testing.T) {
	type fields struct {
		tag    uint8
		num    int
		symbol uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				tag:    tt.fields.tag,
				num:    tt.fields.num,
				symbol: tt.fields.symbol,
			}
			if got := c.Tag(); got != tt.want {
				t.Errorf("Card.Tag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Num(t *testing.T) {
	type fields struct {
		tag    uint8
		num    int
		symbol uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				tag:    tt.fields.tag,
				num:    tt.fields.num,
				symbol: tt.fields.symbol,
			}
			if got := c.Num(); got != tt.want {
				t.Errorf("Card.Num() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Symbol(t *testing.T) {
	type fields struct {
		tag    uint8
		num    int
		symbol uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				tag:    tt.fields.tag,
				num:    tt.fields.num,
				symbol: tt.fields.symbol,
			}
			if got := c.Symbol(); got != tt.want {
				t.Errorf("Card.Symbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHands_FromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		h    Hands
		args args
		want Hands
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.FromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hands.NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHands_ToString(t *testing.T) {
	tests := []struct {
		name string
		h    Hands
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.String(); got != tt.want {
				t.Errorf("Hands.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHands_ToSingleString(t *testing.T) {
	tests := []struct {
		name string
		h    Hands
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ToSingleString(); got != tt.want {
				t.Errorf("Hands.ToSingleString() = %v, want %v", got, tt.want)
			}
		})
	}
}
