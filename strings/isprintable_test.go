package strings

import "testing"

func TestIsPrintable(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{""}, true},
		{"", args{" "}, true},
		{"", args{"abcdefg"}, true},
		{"", args{"abcdefg\n"}, false},
		{"", args{"\u0374"}, true},
		{"", args{"\u0378"}, false},
		// {"", args{"\ud800"}, false},
		{"", args{"\U0001F46F"}, true},
		{"", args{"\U000E0020"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrintable(tt.args.s); got != tt.want {
				t.Errorf("IsPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
