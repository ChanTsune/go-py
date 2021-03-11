package strings

import (
	"reflect"
	"testing"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func TestCaseFold(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"hello"}, "hello"},
		{"", args{"hELlo"}, "hello"},
		{"", args{"ß"}, "ss"},
		{"", args{"ﬁ"}, "fi"},
		{"", args{"\u03a3"}, "\u03c3"},
		{"", args{"A\u0345\u03a3"}, "a\u03b9\u03c3"},
		{"", args{"\u00b5"}, "\u03bc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaseFold(tt.args.s); got != tt.want {
				t.Errorf("CaseFold() = %v, want %v", got, tt.want)
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
