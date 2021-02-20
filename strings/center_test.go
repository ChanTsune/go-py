package strings

import "testing"

func TestCenter(t *testing.T) {
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
		{"", args{"0123", 7, ' '}, "  0123 "},
		{"", args{"abc", 10, ' '}, "   abc    "},
		{"", args{"abc", 6, ' '}, " abc  "},
		{"", args{"abc", 3, ' '}, "abc"},
		{"", args{"abc", 2, ' '}, "abc"},
		{"", args{"abc", 10, '*'}, "***abc****"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Center(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("Center() = %v, want %v", got, tt.want)
			}
		})
	}
}
