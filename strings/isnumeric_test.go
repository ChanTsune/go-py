package strings

import "testing"

func TestIsNumeric(t *testing.T) {
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
		{"", args{"\u2460"}, true},
		// {"", args{"\xbc"}, true}, // TODO
		{"", args{"\u0660"}, true},
		{"", args{"0123456789"}, true},
		{"", args{"0123456789a"}, false},
		{"", args{"\U00010401"}, false},
		{"", args{"\U00010427"}, false},
		{"", args{"\U00010429"}, false},
		{"", args{"\U0001044E"}, false},
		{"", args{"\U0001F40D"}, false},
		{"", args{"\U0001F46F"}, false},
		{"", args{"\U00011065"}, true},
		{"", args{"\U0001D7F6"}, true},
		{"", args{"\U00011066"}, true},
		{"", args{"\U000104A0"}, true},
		{"", args{"\U0001F107"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
