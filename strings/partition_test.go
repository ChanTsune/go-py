package strings

import "testing"

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
		{"", args{"go-py-go-py", "-"}, "go", "-", "py-go-py"},
		{"", args{"this is the partition method", "ti"}, "this is the par", "ti", "tion method"},
		{"", args{"http://www.python.org", "://"}, "http", "://", "www.python.org"},
		{"", args{"http://www.python.org", "?"}, "http://www.python.org", "", ""},
		{"", args{"http://www.python.org", "http://"}, "", "http://", "www.python.org"},
		{"", args{"http://www.python.org", "org"}, "http://www.python.", "org", ""},
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
