package strings

import "testing"

func TestRJust(t *testing.T) {
	type args struct {
		s        string
		width    int
		fillChar rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestRJust", args{"0123", 5, ' '}, " 0123"},
		{"", args{"abc", 10, ' '}, "       abc"},
		{"", args{"abc", 6, ' '}, "   abc"},
		{"", args{"abc", 3, ' '}, "abc"},
		{"", args{"abc", 2, ' '}, "abc"},
		{"", args{"abc", 10, '*'}, "*******abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RJust(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("RJust() = %v, want %v", got, tt.want)
			}
		})
	}
}
