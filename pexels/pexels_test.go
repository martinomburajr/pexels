package pexels

import (
	"testing"
)

var pps1 PexelPhotoSource = PexelPhotoSource{"original1", "large1", "large2x1", "medium1", "small1", "portrait1", "landscape1", "tiny1"}
var pps2 PexelPhotoSource = PexelPhotoSource{"original2", "large2", "large2x2", "medium2", "small2", "portrait2", "landscape2", "tiny2"}

//
var pp1 PexelPhoto = PexelPhoto{14343, 1080, 1920, "https://pexels/some/fake/url", "martin", pps1}
var pp2 PexelPhoto = PexelPhoto{4433, 1080, 1920, "https://pexels/some/fake/url2", "carla", pps2}

func TestPexelPhoto_GetBySize(t *testing.T) {
	tests := []struct {
		name   string
		fields PexelPhoto
		args   string
		want   string
	}{
		{"empty string as size value", pp1, "", "large1"},
		{"random unrelated text", pp1, "asbngaojubgao;ga", "large1"},
		{"normal string: small", pp1, "small", "small1"},
		{"normal string: landscape", pp1, "landscape", "landscape1"},
		{"normal string: large2x", pp2, "laRge2x", "large2x2"},
		{"normal string: medium", pp2, "medium", "medium2"},
		{"normal string: tiny", pp2, "tiny", "tiny2"},
		{"normal string: large", pp2, "large", "large2"},
		{"mixed case string: original", pp2, "origINAL", "original2"},
		{"misspelled string: large2x", pp2, "laRg2x", "large2"},
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
	tests := []struct {
		name string
		args string
		want string
	}{
		{"small with mispelling", "smal", "large"},
		{"landscape with mixed case", "LanDScape", "landscape"},
		{"large2x", "large2x", "large2x"},
		{"empty string", "", "large"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSize(tt.args); got != tt.want {
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
			got, err := pi.GetRandomImage(tt.args.size)
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
