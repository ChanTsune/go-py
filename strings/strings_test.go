package strings

import (
	"testing"
)

func TestSwapCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestSwapCase",
			args: args{
				s: "hello_WORLD",
			},
			want: "HELLO_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapCase(tt.args.s); got != tt.want {
				t.Errorf("SwapCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRJust(t *testing.T) {
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
		{
			name: "TestRJust",
			args: args{
				s:        "0123",
				width:    5,
				fillChar: ' ',
			},
			want: " 0123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RJust(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("RJust() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{
			name: "TestRJust",
			args: args{
				s:        "0123",
				width:    5,
				fillChar: ' ',
			},
			want: "0123 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LJust(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("LJust() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{
			name: "TestRJust",
			args: args{
				s:        "0123",
				width:    7,
				fillChar: ' ',
			},
			want: "  0123 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Center(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("Center() = %v, want %v", got, tt.want)
			}
		})
	}
}
