package strings

import "testing"

func TestExpandTabs(t *testing.T) {
	type args struct {
		s       string
		tabSize int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abc\rab\tdef\ng\thi", 8}, "abc\rab      def\ng       hi"},
		{"", args{"abc\rab\tdef\ng\thi", 4}, "abc\rab  def\ng   hi"},
		{"", args{"abc\r\nab\tdef\ng\thi", 8}, "abc\r\nab      def\ng       hi"},
		{"", args{"abc\r\nab\tdef\ng\thi", 4}, "abc\r\nab  def\ng   hi"},
		{"", args{"abc\r\nab\r\ndef\ng\r\nhi", 4}, "abc\r\nab\r\ndef\ng\r\nhi"},
		{"", args{"abc\rab\tdef\ng\thi", 8}, "abc\rab      def\ng       hi"},
		{"", args{"abc\rab\tdef\ng\thi", 4}, "abc\rab  def\ng   hi"},
		{"", args{" \ta\n\tb", 1}, "  a\n b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandTabs(tt.args.s, tt.args.tabSize); got != tt.want {
				t.Errorf("ExpandTabs() = %v, want %v", got, tt.want)
			}
		})
	}
}
