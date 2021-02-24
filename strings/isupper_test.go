package strings

import "testing"

func TestIsUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{""}, false},
		{"", args{"a"}, false},
		{"", args{"A"}, true},
		{"", args{"\n"}, false},
		{"", args{"ABC"}, true},
		{"", args{"AbC"}, false},
		{"", args{"ABC\n"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUpper(tt.args.s); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
