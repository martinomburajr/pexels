package utils

import (
	"testing"
)

func TestRandBytes(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		{"size -1", args{0}},
		{"size 0", args{0}},
		{"size 1", args{1}},
		{"size 100000000", args{100000000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Utils{}
			got := u.RandBytes(tt.args.size)
			if len(got) < tt.args.size {
				t.Fatalf("length of byte array should be >= size | got %d", len(got))
			}
		})
	}
}
