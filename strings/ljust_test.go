package strings

import "testing"

func TestLJust(t *testing.T) {
	type args struct {
		s        string
		width    int
		fillChar rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abc", 10, ' '}, "abc       "},
		{"", args{"abc", 6, ' '}, "abc   "},
		{"", args{"abc", 3, ' '}, "abc"},
		{"", args{"abc", 2, ' '}, "abc"},
		{"", args{"abc", 10, '*'}, "abc*******"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LJust(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("LJust() = %v, want %v", got, tt.want)
			}
		})
	}
}
