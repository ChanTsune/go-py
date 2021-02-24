package strings

import "testing"

func TestJoin(t *testing.T) {
	type args struct {
		s        string
		iterable []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{" ", []string{"a", "b", "c", "d"}}, "a b c d"},
		{"", args{"", []string{"a", "b", "c", "d"}}, "abcd"},
		{"", args{"", []string{"", "b", "", "d"}}, "bd"},
		{"", args{"", []string{"a", "", "c", ""}}, "ac"},
		{"", args{" ", []string{"w", "x", "y", "z"}}, "w x y z"},
		{"", args{"a", []string{"abc"}}, "abc"},
		{"", args{"a", []string{"z"}}, "z"},
		{"", args{".", []string{"a", "b", "c"}}, "a.b.c"},
		{"", args{" ", []string{"a", "b", "c"}}, "a b c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.s, tt.args.iterable); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
