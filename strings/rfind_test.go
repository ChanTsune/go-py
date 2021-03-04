package strings

import "testing"

func TestRFind(t *testing.T) {
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
		{"", args{"abcdefghiabc", "abc", []Option{}}, 9},
		{"", args{"abcdefghiabc", "", []Option{}}, 12},
		{"", args{"abcdefghiabc", "abcd", []Option{}}, 0},
		{"", args{"abcdefghiabc", "abcz", []Option{}}, -1},
		{"", args{"abc", "", []Option{Start(0)}}, 3},
		{"", args{"abc", "", []Option{Start(3)}}, 3},
		{"", args{"abc", "", []Option{Start(4)}}, -1},
		{"", args{"rrarrrrrrrrra", "a", []Option{}}, 12},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4)}}, 12},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4), End(6)}}, -1},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4)}}, 12},
		{"", args{"rrarrrrrrrrra", "a", []Option{End(6)}}, 2},
		{"", args{"ab", "xxx", []Option{Start(MaxInt), End(0)}}, -1},
		{"", args{"<......\u043C...", "<", []Option{}}, 0},
		{"", args{"hello", "l", []Option{}}, 3},
		{"", args{"hello", "l", []Option{Start(-2)}}, 3},
		{"", args{"hello", "l", []Option{End(-2)}}, 2},
		{"", args{"hello", "h", []Option{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RFind(tt.args.s, tt.args.sub, tt.args.options...); got != tt.want {
				t.Errorf("RFind() = %v, want %v", got, tt.want)
			}
		})
	}
}
