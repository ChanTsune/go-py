package strings

import "testing"

func TestIsAlpha(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{""}, false},
		{"", args{"a"}, true},
		{"", args{"A"}, true},
		{"", args{"\n"}, false},
		{"", args{"abc"}, true},
		{"", args{"aBc123"}, false},
		{"", args{"abc\n"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.args.s); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}
