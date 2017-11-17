package board

import (
	"reflect"
	"testing"
)

func TestBoard_FromString(t *testing.T) {
	type args struct {
		h Hands
		s string
	}
	tests := []struct {
		name string
		deck Board
		args args
		want Board
	}{
	// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.deck.FromString(tt.args.h, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
