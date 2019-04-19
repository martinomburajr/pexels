package utils

import (
	"crypto/rand"
	"reflect"
	"testing"
)

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
			if (tt.args.max <= 1) && (got != 0) {
				t.Errorf("if args == 1 or 0 must return 0")
			}
			if (got > tt.args.max && tt.args.max > 0) || (got < 0) {
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

func TestWriteToFile(t *testing.T) {
	rand1 := generateRandomBytes(0)
	rand2 := generateRandomBytes(1)
	rand3 := generateRandomBytes(100)
	type args struct {
		filepath string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"writing 0 bytes", args{"testdata/1.jpg", rand1}, false},
		{"writing 10 bytes", args{"testdata/2.jpg", rand2}, false},
		{"writing 100 bytes", args{"testdata/3.jpg", rand3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteToFile(tt.args.filepath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func generateRandomBytes(size int) []byte {
	if size < 0 {
		token := make([]byte, 100)
		rand.Read(token)
		return  token
	}
	token := make([]byte, size)
	rand.Read(token)
	return token
}