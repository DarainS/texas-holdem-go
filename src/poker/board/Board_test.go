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
		name  string
		board Board
		args  args
		want  Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(tt.args.h, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard.NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
