package strings

import "testing"

func TestRIndex(t *testing.T) {
	type args struct {
		s       string
		sub     string
		options []Option
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"", args{"abcdefghiabc", "", []Option{}}, 12, false},
		{"", args{"abcdefghiabc", "def", []Option{}}, 3, false},
		{"", args{"abcdefghiabc", "abc", []Option{}}, 9, false},
		{"", args{"abcdefghiabc", "abc", []Option{Start(0), End(-1)}}, 0, false},
		{"", args{"abcdefghiabc", "hib", []Option{}}, -1, true},
		{"", args{"defghiabc", "def", []Option{Start(1)}}, -1, true},
		{"", args{"defghiabc", "abc", []Option{Start(0), End(-1)}}, -1, true},
		{"", args{"abcdefghi", "ghi", []Option{Start(0), End(8)}}, -1, true},
		{"", args{"abcdefghi", "ghi", []Option{Start(0), End(-1)}}, -1, true},
		{"", args{"rrarrrrrrrrra", "a", []Option{}}, 12, false},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4)}}, 12, false},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4), End(6)}}, -1, true},
		{"", args{"rrarrrrrrrrra", "a", []Option{Start(4)}}, 12, false},
		{"", args{"rrarrrrrrrrra", "a", []Option{End(6)}}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RIndex(tt.args.s, tt.args.sub, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
