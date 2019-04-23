package pexels

import (
	"reflect"
	"testing"
)

func TestPexelPhoto_Get(t *testing.T) {
	type fields struct {
		ID           int
		Width        int
		Height       int
		URL          string
		Photographer string
		Source       PexelPhotoSource
	}
	type args struct {
		id   string
		size string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &PexelPhoto{
				ID:           tt.fields.ID,
				Width:        tt.fields.Width,
				Height:       tt.fields.Height,
				URL:          tt.fields.URL,
				Photographer: tt.fields.Photographer,
				Source:       tt.fields.Source,
			}
			got, err := pi.Get(tt.args.id, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("PexelPhoto.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PexelPhoto.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSize(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSize(tt.args.size); got != tt.want {
				t.Errorf("parseSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPexelPhoto_GetRandomImage(t *testing.T) {
	type fields struct {
		ID           int
		Width        int
		Height       int
		URL          string
		Photographer string
		Source       PexelPhotoSource
	}
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &PexelPhoto{
				ID:           tt.fields.ID,
				Width:        tt.fields.Width,
				Height:       tt.fields.Height,
				URL:          tt.fields.URL,
				Photographer: tt.fields.Photographer,
				Source:       tt.fields.Source,
			}
			_, got, err := pi.GetRandomImage(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("PexelPhoto.GetRandomImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PexelPhoto.GetRandomImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPexelPhoto_GetBySize(t *testing.T) {
	type fields struct {
		ID           int
		Width        int
		Height       int
		URL          string
		Photographer string
		Source       PexelPhotoSource
	}
	type args struct {
		size string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := &PexelPhoto{
				ID:           tt.fields.ID,
				Width:        tt.fields.Width,
				Height:       tt.fields.Height,
				URL:          tt.fields.URL,
				Photographer: tt.fields.Photographer,
				Source:       tt.fields.Source,
			}
			if got := pi.GetBySize(tt.args.size); got != tt.want {
				t.Errorf("PexelPhoto.GetBySize() = %v, want %v", got, tt.want)
			}
		})
	}
}
