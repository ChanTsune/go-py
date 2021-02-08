package strings

import (
	"reflect"
	"testing"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

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

func TestCount(t *testing.T) {
	type args struct {
		s       string
		sub     string
		options []Option
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"aaa", "a", []Option{}}, 3},
		{"", args{"aaa", "b", []Option{}}, 0},
		{"", args{"aaa", "a", []Option{Start(1)}}, 2},
		{"", args{"aaa", "a", []Option{Start(10)}}, 0},
		{"", args{"aaa", "a", []Option{Start(-1)}}, 1},
		{"", args{"aaa", "a", []Option{Start(-10)}}, 3},
		{"", args{"aaa", "a", []Option{Start(0), End(1)}}, 1},
		{"", args{"aaa", "a", []Option{Start(0), End(10)}}, 3},
		{"", args{"aaa", "a", []Option{Start(0), End(-1)}}, 2},
		{"", args{"aaa", "a", []Option{Start(0), End(-10)}}, 0},
		{"", args{"aaa", "", []Option{Start(1)}}, 3},
		{"", args{"aaa", "", []Option{Start(3)}}, 1},
		{"", args{"aaa", "", []Option{Start(10)}}, 0},
		{"", args{"aaa", "", []Option{Start(-1)}}, 2},
		{"", args{"aaa", "", []Option{Start(-10)}}, 4},
		{"", args{"", "", []Option{}}, 1},
		{"", args{"", "", []Option{Start(1), End(1)}}, 0},
		{"", args{"", "", []Option{Start(MaxInt), End(0)}}, 0},
		{"", args{"", "xx", []Option{}}, 0},
		{"", args{"", "xx", []Option{Start(1), End(1)}}, 0},
		{"", args{"", "xx", []Option{Start(MaxInt), End(0)}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s, tt.args.sub, tt.args.options...); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestSwapCase",
			args: args{
				s: "hello_WORLD",
			},
			want: "HELLO_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapCase(tt.args.s); got != tt.want {
				t.Errorf("SwapCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{
			name: "TestRJust",
			args: args{
				s:        "0123",
				width:    5,
				fillChar: ' ',
			},
			want: " 0123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RJust(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("RJust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLJust(t *testing.T) {
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
		{
			name: "TestRJust",
			args: args{
				s:        "0123",
				width:    5,
				fillChar: ' ',
			},
			want: "0123 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LJust(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("LJust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCenter(t *testing.T) {
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
		{
			name: "TestRJust",
			args: args{
				s:        "0123",
				width:    7,
				fillChar: ' ',
			},
			want: "  0123 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Center(tt.args.s, tt.args.width, tt.args.fillChar); got != tt.want {
				t.Errorf("Center() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestLength",
			args: args{
				s: "01234",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Length(tt.args.s); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			name: "TestPartition",
			args: args{
				s:   "go-py-go-py",
				sep: "-",
			},
			want:  "go",
			want1: "-",
			want2: "py-go-py",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Partition(tt.args.s, tt.args.sep)
			if got != tt.want {
				t.Errorf("Partition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Partition() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Partition() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestRPartition(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{
			name: "TestRPartition",
			args: args{
				s:   "go-py-go-py",
				sep: "-",
			},
			want:  "go-py-go",
			want1: "-",
			want2: "py",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := RPartition(tt.args.s, tt.args.sep)
			if got != tt.want {
				t.Errorf("RPartition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RPartition() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("RPartition() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestLastSplitN(t *testing.T) {
	type args struct {
		s   string
		sep string
		n   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestLastSplitN",
			args: args{
				s:   "0,1,2,3",
				sep: ",",
				n:   3,
			},
			want: []string{"0,1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastSplitN(tt.args.s, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastSplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestLastSplit",
			args: args{
				s:   "0,1,2,3",
				sep: ",",
			},
			want: []string{"0", "1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastSplit(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastSplitAfterN(t *testing.T) {
	type args struct {
		s   string
		sep string
		n   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestLastSplitAfterN",
			args: args{
				s:   "0,1,2,3",
				sep: ",",
				n:   3,
			},
			want: []string{"0,1,", "2,", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastSplitAfterN(tt.args.s, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastSplitAfterN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastSplitAfter(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestLastSplitAfterN",
			args: args{
				s:   "0,1,2,3",
				sep: ",",
			},
			want: []string{"0,", "1,", "2,", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastSplitAfter(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastSplitAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
