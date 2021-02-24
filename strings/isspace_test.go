package strings

import "testing"

func TestIsSpace(t *testing.T) {
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
		{"", args{" "}, true},
		{"", args{"\t"}, true},
		{"", args{"\r"}, true},
		{"", args{"\n"}, true},
		{"", args{" \t\r\n"}, true},
		{"", args{" \t\r\na"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSpace(tt.args.s); got != tt.want {
				t.Errorf("IsSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
