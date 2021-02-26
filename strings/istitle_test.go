package strings

import "testing"

func TestIsTitle(t *testing.T) {
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
		{"", args{"A Titlecased Line"}, true},
		{"", args{"A\nTitlecased Line"}, true},
		{"", args{"A Titlecased, Line"}, true},
		{"", args{"Not a capitalized String"}, false},
		{"", args{"Not\ta Titlecase String"}, false},
		{"", args{"Not--a Titlecase String"}, false},
		{"", args{"NOT"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTitle(tt.args.s); got != tt.want {
				t.Errorf("IsTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
