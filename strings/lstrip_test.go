package strings

import (
	"testing"
)

func TestLStrip(t *testing.T) {
	type args struct {
		s     string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"xyzzyhelloxyzzy", "xyz"}, "helloxyzzy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LStrip(tt.args.s, tt.args.chars); got != tt.want {
				t.Errorf("LStrip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLStripWhiteSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"   hello   "}, "hello   "},
		{"", args{" \t\n\r\u000c\u000Babc \t\n\r\u000c\u000B"}, "abc \t\n\r\u000c\u000B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LStripWhiteSpace(tt.args.s); got != tt.want {
				t.Errorf("LStripWhiteSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
