package strings

import "testing"

func TestZFill(t *testing.T) {
	type args struct {
		s     string
		width int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"123", 2}, "123"},
		{"", args{"123", 3}, "123"},
		{"", args{"123", 4}, "0123"},
		{"", args{"+123", 3}, "+123"},
		{"", args{"+123", 4}, "+123"},
		{"", args{"+123", 5}, "+0123"},
		{"", args{"-123", 3}, "-123"},
		{"", args{"-123", 4}, "-123"},
		{"", args{"-123", 5}, "-0123"},
		{"", args{"", 3}, "000"},
		{"", args{"34", 1}, "34"},
		{"", args{"34", 4}, "0034"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZFill(tt.args.s, tt.args.width); got != tt.want {
				t.Errorf("ZFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
