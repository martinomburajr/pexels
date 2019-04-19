package utils

import (
	"reflect"
	"testing"
)

func TestWriteImageToFile(t *testing.T) {
	type args struct {
		filepath string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteImageToFile(tt.args.filepath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteImageToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChangeUbuntuBackground(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ChangeUbuntuBackground(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("ChangeUbuntuBackground() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRandIntBetween(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"max positive integer > 1", args{max: 10}, 3},
		{"integer 1 should return 0", args{max: 1}, 0},
		{"integer 0 should return 0", args{max: 0}, 0},
		{"integer less than 0", args{max: -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandIntBetween(tt.args.max)
			if  (tt.args.max <= 1) && (got != 0) {
				t.Errorf("if args == 1 or 0 must return 0")
			}
			if  (got > tt.args.max && tt.args.max > 0) || (got < 0) {
				t.Errorf("got %d is out of range max %d", got, tt.args.max)
			}
		})
	}
}

func TestParseRequest(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRequest(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
