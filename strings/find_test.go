package strings

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		s       string
		sub     string
		options []Option
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"abcdefghiabc", "abc", []Option{}}, 0},
		{"", args{"abcdefghiabc", "abc", []Option{Start(1)}}, 9},
		{"", args{"abcdefghiabc", "def", []Option{Start(4)}}, -1},
		{"", args{"abc", "", []Option{Start(0)}}, 0},
		{"", args{"abc", "", []Option{Start(3)}}, 3},
		{"", args{"abc", "", []Option{Start(4)}}, -1},
		{"", args{"rrarrrrrrrrra", "a", []Option{}}, 2},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4)}}, 12},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4), End(6)}}, -1},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4)}}, 12},
		{"", args{"rrarrrrrrrrra", "a", []Option{End(6)}}, 2},
		{"", args{"", "", []Option{}}, 0},
		{"", args{"", "", []Option{Start(1), End(1)}}, -1},
		{"", args{"", "", []Option{Start(MaxInt), End(0)}}, -1},
		{"", args{"", "xx", []Option{}}, -1},
		{"", args{"", "xx", []Option{Start(1), End(1)}}, -1},
		{"", args{"", "xx", []Option{Start(MaxInt), End(0)}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.s, tt.args.sub, tt.args.options...); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
