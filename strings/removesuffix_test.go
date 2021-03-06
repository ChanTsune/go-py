package strings

import "testing"

func TestRemoveSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"spam", "am"}, "sp"},
		{"", args{"spamspamspam", "spam"}, "spamspam"},
		{"", args{"spam", "python"}, "spam"},
		{"", args{"spam", "blam"}, "spam"},
		{"", args{"spam", "eggs and spam"}, "spam"},
		{"", args{"", ""}, ""},
		{"", args{"", "abcde"}, ""},
		{"", args{"abcde", ""}, "abcde"},
		{"", args{"abcde", "abcde"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("RemoveSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
