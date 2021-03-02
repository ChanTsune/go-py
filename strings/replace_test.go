package strings

import (
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	type args struct {
		s   string
		old string
		new string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"", "", ""}, ""},
		{"", args{"", "", "A"}, "A"},
		{"", args{"", "A", ""}, ""},
		{"", args{"", "A", "A"}, ""},
		{"", args{"A", "", ""}, "A"},
		{"", args{"A", "", "*"}, "*A*"},
		{"", args{"A", "", "*1"}, "*1A*1"},
		{"", args{"A", "", "*-#"}, "*-#A*-#"},
		{"", args{"AA", "", "*-"}, "*-A*-A*-"},
		{"", args{"A", "A", ""}, ""},
		{"", args{"AAA", "A", ""}, ""},
		{"", args{"AAAAAAAAAA", "A", ""}, ""},
		{"", args{"ABACADA", "A", ""}, "BCD"},

		{"", args{"ABCAD", "A", ""}, "BCD"},
		{"", args{"ABCADAA", "A", ""}, "BCD"},
		{"", args{"BCD", "A", ""}, "BCD"},
		{"", args{"*************", "A", ""}, "*************"},
		{"", args{"the", "the", ""}, ""},
		{"", args{"theater", "the", ""}, "ater"},
		{"", args{"thethe", "the", ""}, ""},
		{"", args{"thethethethe", "the", ""}, ""},
		{"", args{"theatheatheathea", "the", ""}, "aaaa"},
		{"", args{"that", "the", ""}, "that"},
		{"", args{"thaet", "the", ""}, "thaet"},
		{"", args{"here and there", "the", ""}, "here and re"},
		{"", args{"here and there and there", "the", ""}, "here and re and re"},
		{"", args{"abc", "the", ""}, "abc"},
		{"", args{"abcdefg", "the", ""}, "abcdefg"},
		{"", args{"bbobob", "bob", ""}, "bob"},
		{"", args{"bbobobXbbobob", "bob", ""}, "bobXbob"},
		{"", args{"aaaaaaabob", "bob", ""}, "aaaaaaa"},
		{"", args{"aaaaaaa", "bob", ""}, "aaaaaaa"},
		{"", args{"Who goes there?", "o", "o"}, "Who goes there?"},
		{"", args{"Who goes there?", "o", "O"}, "WhO gOes there?"},
		{"", args{"Who goes there?", "a", "q"}, "Who goes there?"},
		{"", args{"Who goes there?", "W", "w"}, "who goes there?"},
		{"", args{"WWho goes there?WW", "W", "w"}, "wwho goes there?ww"},
		{"", args{"Who goes there?", "?", "!"}, "Who goes there!"},
		{"", args{"Who goes there??", "?", "!"}, "Who goes there!!"},
		{"", args{"Who goes there?", ".", "!"}, "Who goes there?"},
		{"", args{"This is a tissue", "is", "**"}, "Th** ** a t**sue"},
		{"", args{"bobob", "bob", "cob"}, "cobob"},
		{"", args{"bobobXbobobob", "bob", "cob"}, "cobobXcobocob"},
		{"", args{"bobob", "bot", "bot"}, "bobob"},
		{"", args{"Reykjavik", "k", "KK"}, "ReyKKjaviKK"},
		{"", args{"A.B.C.", ".", "----"}, "A----B----C----"},
		{"", args{"...\u043C......<", "<", "&lt;"}, "...\u043C......&lt;"},
		{"", args{"Reykjavik", "q", "KK"}, "Reykjavik"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham"}, "ham, ham, eggs and ham"},
		{"", args{"bobobob", "bobob", "bob"}, "bobob"},
		{"", args{"bobobobXbobobob", "bobob", "bob"}, "bobobXbobob"},
		{"", args{"BOBOBOB", "bob", "bobby"}, "BOBOBOB"},
		{"", args{"one!two!three!", "!", ""}, "onetwothree"},
		{"", args{"one!two!three!", "!", "@"}, "one@two@three@"},
		{"", args{"one!two!three!", "x", "@"}, "one!two!three!"},
		{"", args{"abc", "", "-"}, "-a-b-c-"},
		{"", args{"", "", ""}, ""},
		{"", args{"abc", "xy", "--"}, "abc"},
		{"", args{"123", "123", ""}, ""},
		{"", args{"123123", "123", ""}, ""},
		{"", args{"123x123", "123", ""}, "x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.args.s, tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceN(t *testing.T) {
	type args struct {
		s     string
		old   string
		new   string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"", "", "", 100}, ""},
		{"", args{"", "", "A", 100}, "A"},
		{"", args{"", "", "", MaxInt}, ""},
		{"", args{"AA", "", "*-", -1}, "*-A*-A*-"},
		{"", args{"AA", "", "*-", MaxInt}, "*-A*-A*-"},
		{"", args{"AA", "", "*-", 4}, "*-A*-A*-"},
		{"", args{"AA", "", "*-", 3}, "*-A*-A*-"},
		{"", args{"AA", "", "*-", 2}, "*-A*-A"},
		{"", args{"AA", "", "*-", 1}, "*-AA"},
		{"", args{"AA", "", "*-", 0}, "AA"},
		{"", args{"AAA", "A", "", -1}, ""},
		{"", args{"AAA", "A", "", MaxInt}, ""},
		{"", args{"AAA", "A", "", 4}, ""},
		{"", args{"AAA", "A", "", 3}, ""},
		{"", args{"AAA", "A", "", 2}, "A"},
		{"", args{"AAA", "A", "", 1}, "AA"},
		{"", args{"AAA", "A", "", 0}, "AAA"},
		{"", args{"ABACADA", "A", "", -1}, "BCD"},
		{"", args{"ABACADA", "A", "", MaxInt}, "BCD"},
		{"", args{"ABACADA", "A", "", 5}, "BCD"},
		{"", args{"ABACADA", "A", "", 4}, "BCD"},
		{"", args{"ABACADA", "A", "", 3}, "BCDA"},
		{"", args{"ABACADA", "A", "", 2}, "BCADA"},
		{"", args{"ABACADA", "A", "", 1}, "BACADA"},
		{"", args{"ABACADA", "A", "", 0}, "ABACADA"},
		{"", args{("^" + strings.Repeat("A", 1000) + "^"), "A", "", 999}, "^A^"},
		{"", args{"here and there and there", "the", "", MaxInt}, "here and re and re"},
		{"", args{"here and there and there", "the", "", -1}, "here and re and re"},
		{"", args{"here and there and there", "the", "", 3}, "here and re and re"},
		{"", args{"here and there and there", "the", "", 2}, "here and re and re"},
		{"", args{"here and there and there", "the", "", 1}, "here and re and there"},
		{"", args{"here and there and there", "the", "", 0}, "here and there and there"},
		{"", args{"Who goes there?", "o", "O", MaxInt}, "WhO gOes there?"},
		{"", args{"Who goes there?", "o", "O", -1}, "WhO gOes there?"},
		{"", args{"Who goes there?", "o", "O", 3}, "WhO gOes there?"},
		{"", args{"Who goes there?", "o", "O", 2}, "WhO gOes there?"},
		{"", args{"Who goes there?", "o", "O", 1}, "WhO goes there?"},
		{"", args{"Who goes there?", "o", "O", 0}, "Who goes there?"},
		{"", args{"This is a tissue", "is", "**", MaxInt}, "Th** ** a t**sue"},
		{"", args{"This is a tissue", "is", "**", -1}, "Th** ** a t**sue"},
		{"", args{"This is a tissue", "is", "**", 4}, "Th** ** a t**sue"},
		{"", args{"This is a tissue", "is", "**", 3}, "Th** ** a t**sue"},
		{"", args{"This is a tissue", "is", "**", 2}, "Th** ** a tissue"},
		{"", args{"This is a tissue", "is", "**", 1}, "Th** is a tissue"},
		{"", args{"This is a tissue", "is", "**", 0}, "This is a tissue"},
		{"", args{"Reykjavik", "k", "KK", -1}, "ReyKKjaviKK"},
		{"", args{"Reykjavik", "k", "KK", MaxInt}, "ReyKKjaviKK"},
		{"", args{"Reykjavik", "k", "KK", 2}, "ReyKKjaviKK"},
		{"", args{"Reykjavik", "k", "KK", 1}, "ReyKKjavik"},
		{"", args{"Reykjavik", "k", "KK", 0}, "Reykjavik"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", MaxInt}, "ham, ham, eggs and ham"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", -1}, "ham, ham, eggs and ham"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", 4}, "ham, ham, eggs and ham"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", 3}, "ham, ham, eggs and ham"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", 2}, "ham, ham, eggs and spam"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", 1}, "ham, spam, eggs and spam"},
		{"", args{"spam, spam, eggs and spam", "spam", "ham", 0}, "spam, spam, eggs and spam"},
		{"", args{"one!two!three!", "!", "@", 1}, "one@two!three!"},
		{"", args{"one!two!three!", "!", "@", 2}, "one@two@three!"},
		{"", args{"one!two!three!", "!", "@", 3}, "one@two@three@"},
		{"", args{"one!two!three!", "!", "@", 4}, "one@two@three@"},
		{"", args{"one!two!three!", "!", "@", 0}, "one!two!three!"},
		{"", args{"one!two!three!", "x", "@", 2}, "one!two!three!"},
		{"", args{"abc", "", "-", 3}, "-a-b-c"},
		{"", args{"abc", "", "-", 0}, "abc"},
		{"", args{"abc", "ab", "--", 0}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceN(tt.args.s, tt.args.old, tt.args.new, tt.args.count); got != tt.want {
				t.Errorf("ReplaceN() = %v, want %v", got, tt.want)
			}
		})
	}
}
