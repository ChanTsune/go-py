package errors

import "testing"

func TestValueError_Error(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"error"}, "ValueError: error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewValueError(Message(tt.fields.message))
			if got := e.Error(); got != tt.want {
				t.Errorf("ValueError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
