package strings

import "testing"

func TestCapitalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{""}, ""},
		{"", args{"a"}, "A"},
		{"", args{" hello "}, " hello "},
		{"", args{"Hello "}, "Hello "},
		{"", args{"hello "}, "Hello "},
		{"", args{"aaaa"}, "Aaaa"},
		{"", args{"AaAa"}, "Aaaa"},
		{"", args{"\u1FF3\u1FF3\u1FFC\u1FFC"}, "\u1FFC\u1FF3\u1FF3\u1FF3"},
		{"", args{"\u24C5\u24CE\u24C9\u24BD\u24C4\u24C3"}, "\u24C5\u24E8\u24E3\u24D7\u24DE\u24DD"},
		{"", args{"\u24DF\u24E8\u24E3\u24D7\u24DE\u24DD"}, "\u24C5\u24E8\u24E3\u24D7\u24DE\u24DD"},
		{"", args{"\u2160\u2161\u2162"}, "\u2160\u2171\u2172"},
		{"", args{"\u2170\u2171\u2172"}, "\u2160\u2171\u2172"},
		{"", args{"\u019B\u1D00\u1D86\u0221\u1FB7"}, "\u019B\u1D00\u1D86\u0221\u1FB7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.s); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
