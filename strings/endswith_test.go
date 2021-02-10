package strings

import (
	"testing"
)

func TestEndsWith(t *testing.T) {
	type args struct {
		s       string
		suffix  string
		options []Option
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"hello", "lo", []Option{}}, true},
		{"", args{"hello", "he", []Option{}}, false},
		{"", args{"hello", "", []Option{}}, true},
		{"", args{"hello", "hello world", []Option{}}, false},
		{"", args{"helloworld", "worl", []Option{}}, false},
		{"", args{"helloworld", "worl", []Option{Start(3), End(9)}}, true},
		{"", args{"helloworld", "world", []Option{Start(3), End(12)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(1), End(7)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(2), End(7)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(3), End(7)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(4), End(7)}}, false},
		{"", args{"helloworld", "lowo", []Option{Start(3), End(8)}}, false},
		{"", args{"ab", "ab", []Option{Start(0), End(1)}}, false},
		{"", args{"ab", "ab", []Option{Start(0), End(0)}}, false},
		{"", args{"", "", []Option{Start(0), End(1)}}, true},
		{"", args{"", "", []Option{Start(0), End(0)}}, true},
		{"", args{"", "", []Option{Start(1), End(0)}}, false},
		{"", args{"hello", "lo", []Option{Start(-2)}}, true},
		{"", args{"hello", "he", []Option{Start(-2)}}, false},
		{"", args{"hello", "", []Option{Start(-3), End(-3)}}, true},
		{"", args{"hello", "hello world", []Option{Start(-10), End(-2)}}, false},
		{"", args{"helloworld", "worl", []Option{Start(-6)}}, false},
		{"", args{"helloworld", "worl", []Option{Start(-5), End(-1)}}, true},
		{"", args{"helloworld", "worl", []Option{Start(-5), End(9)}}, true},
		{"", args{"helloworld", "world", []Option{Start(-7), End(12)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(-99), End(-3)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(-8), End(-3)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(-7), End(-3)}}, true},
		{"", args{"helloworld", "lowo", []Option{Start(3), End(-4)}}, false},
		{"", args{"helloworld", "lowo", []Option{Start(-8), End(-2)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndsWith(tt.args.s, tt.args.suffix, tt.args.options...); got != tt.want {
				t.Errorf("EndsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndsWithAny(t *testing.T) {
	type args struct {
		s        string
		suffixes []string
		options  []Option
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"hello", []string{"he", "ha"}, []Option{}}, false},
		{"", args{"hello", []string{"lo", "llo"}, []Option{}}, true},
		{"", args{"hello", []string{"hellox", "hello"}, []Option{}}, true},
		{"", args{"hello", []string{}, []Option{}}, false},
		{"", args{"helloworld", []string{"hellowo", "rld", "lowo"}, []Option{Start(3)}}, true},
		{"", args{"helloworld", []string{"hellowo", "ello", "rld"}, []Option{Start(3), End(-1)}}, false},
		{"", args{"hello", []string{"hell", "ell"}, []Option{Start(0), End(-1)}}, true},
		{"", args{"hello", []string{"he", "hel"}, []Option{Start(0), End(1)}}, false},
		{"", args{"hello", []string{"he", "hell"}, []Option{Start(0), End(4)}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndsWithAny(tt.args.s, tt.args.suffixes, tt.args.options...); got != tt.want {
				t.Errorf("EndsWithAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
