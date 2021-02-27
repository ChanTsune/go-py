package strings

import "testing"

func TestRemoevPrefix(t *testing.T) {
	type args struct {
		s      string
		prefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"spam", "sp"}, "am"},
		{"", args{"spamspamspam", "spam"}, "spamspam"},
		{"", args{"spam", "python"}, "spam"},
		{"", args{"spam", "spider"}, "spam"},
		{"", args{"spam", "spam and eggs"}, "spam"},
		{"", args{"", ""}, ""},
		{"", args{"", "abcde"}, ""},
		{"", args{"abcde", ""}, "abcde"},
		{"", args{"abcde", "abcde"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoevPrefix(tt.args.s, tt.args.prefix); got != tt.want {
				t.Errorf("RemoevPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
