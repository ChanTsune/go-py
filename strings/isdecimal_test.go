package strings

import "testing"

func TestIsDecimal(t *testing.T) {
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
		{"", args{"0"}, true},
		{"", args{"\u2460"}, false}, // CIRCLED DIGIT ONE
		{"", args{"\xbc"}, false},   // VULGAR FRACTION ONE QUARTER
		{"", args{"\u0660"}, true},  // ARABIC-INDIC DIGIT ZERO
		{"", args{"0123456789"}, true},
		{"", args{"0123456789a"}, false},
		{"", args{"\U00010401"}, false},
		{"", args{"\U00010427"}, false},
		{"", args{"\U00010429"}, false},
		{"", args{"\U0001044E"}, false},
		{"", args{"\U0001F40D"}, false},
		{"", args{"\U0001F46F"}, false},
		{"", args{"\U00011065"}, false},
		{"", args{"\U0001F107"}, false},
		{"", args{"\U0001D7F6"}, true},
		{"", args{"\U00011066"}, true},
		{"", args{"\U000104A0"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDecimal(tt.args.s); got != tt.want {
				t.Errorf("IsDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
