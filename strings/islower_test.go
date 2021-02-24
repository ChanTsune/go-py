package strings

import "testing"

func TestIsLower(t *testing.T) {
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
		{"", args{"A"}, false},
		{"", args{"\n"}, false},
		{"", args{"abc"}, true},
		{"", args{"aBc"}, false},
		{"", args{"abc\n"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLower(tt.args.s); got != tt.want {
				t.Errorf("IsLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
