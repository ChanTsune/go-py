package strings

import "testing"

func TestIsAlNum(t *testing.T) {
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
		{"", args{"123abc456"}, true},
		{"", args{"a1b3c"}, true},
		{"", args{"aBc000 "}, false},
		{"", args{"abc\n"}, false},
		{"", args{"\U00010401"}, true},
		{"", args{"\U00010427"}, true},
		{"", args{"\U00010429"}, true},
		{"", args{"\U0001044E"}, true},
		{"", args{"\U0001D7F6"}, true},
		{"", args{"\U00011066"}, true},
		{"", args{"\U000104A0"}, true},
		{"", args{"\U0001F107"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlNum(tt.args.s); got != tt.want {
				t.Errorf("IsAlNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
