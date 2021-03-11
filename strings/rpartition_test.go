package strings

import "testing"

func TestRPartition(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	S := "http://www.python.org"
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{"", args{"go-py-go-py", "-"}, "go-py-go", "-", "py"},
		{"", args{"this is the rpartition method", "ti"}, "this is the rparti", "ti", "on method"},
		{"", args{S, "://"}, "http", "://", "www.python.org"},
		{"", args{S, "?"}, "", "", "http://www.python.org"},
		{"", args{S, "http://"}, "", "http://", "www.python.org"},
		{"", args{S, "org"}, "http://www.python.", "org", ""},
		{"", args{S, ""}, "http://www.python.or", "", "g"}, // This is not compat python. Python will raise error.
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
