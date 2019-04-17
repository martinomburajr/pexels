package pexels

import (
	"testing"
)

var pps1 PexelPhotoSource = PexelPhotoSource{"orig1", "large1", "large2x1", "medium1", "small1", "portrait1", "landscape1", "tiny1"}
var pps2 PexelPhotoSource = PexelPhotoSource{"orig2", "large2", "large2x2", "medium2", "small2", "portrait2", "landscape2", "tiny2"}

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
		{"normal string: small", pp1, "small", "small1"},
		{"normal string: landscape", pp1, "landscape", "landscape1"},
		{"normal string: large2x", pp2, "laRge2x", "large2x2"},
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
			if got := pi.GetBySize(tt.args); got != tt.want {
				t.Errorf("PexelPhoto.GetBySize() = %v, want %v", got, tt.want)
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
