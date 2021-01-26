package strings

import (
	"reflect"
	"testing"
)

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
