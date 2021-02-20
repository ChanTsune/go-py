package strings

import (
	"strings"
	"testing"
)

func TestIsASCII_Simple(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{""}, true},
		{"", args{"\u0000"}, true},
		{"", args{"\u0000"}, true},
		{"", args{"\u0080"}, false},
		{"", args{"\u00E9"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsASCII(tt.args.s); got != tt.want {
				t.Errorf("IsASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsASCII(t *testing.T) {
	type args struct {
		s string
	}
	for p := range strings.Repeat(" ", 8) { // NOTE: this is simple 8 loop. TODO: change to more simply
		tests := []struct {
			name string
			args args
			want bool
		}{
			{"", args{strings.Repeat(" ", p) + ""}, true},
			{"", args{strings.Repeat(" ", p) + "\u0080"}, false},
			{"", args{strings.Repeat(" ", p) + "" + strings.Repeat(" ", 8)}, true},
			{"", args{strings.Repeat(" ", p) + "\u0080" + strings.Repeat(" ", 8)}, false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := IsASCII(tt.args.s); got != tt.want {
					t.Errorf("IsASCII() = %v, want %v", got, tt.want)
				}
			})
		}

	}
}
