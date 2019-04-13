package pexels

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPexelPhoto_Get(t *testing.T) {
	type fields struct {
		Width        int
		Height       int
		URL          string
		Photographer string
		Source       PexelPhotoSource
	}
	type args struct {
		id string
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
				Width:        tt.fields.Width,
				Height:       tt.fields.Height,
				URL:          tt.fields.URL,
				Photographer: tt.fields.Photographer,
				Source:       tt.fields.Source,
			}
			got, err := pi.Get(tt.args.id)
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

func HandlerTest(writer http.ResponseWriter, request *http.Request) {

}

func Test_parseRequest(t *testing.T) {

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/gist", nil)
	handler := http.HandlerFunc(HandlerTest)

	handler.ServeHTTP(recorder, request)

	//type args struct {
	//	urlWSize string
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    []byte
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := parseRequest(tt.args.urlWSize)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("parseRequest() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("parseRequest() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
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
		{name: "should return large", args: args{size: "hello"}, want: ImageSizeLarge},
		{name: "should return medium", args: args{size: ImageSizeMedium}, want: ImageSizeMedium},
		{name: "should return portrait", args: args{size: "PoRtRaiT"}, want: ImageSizePortrait},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSize(tt.args.size); got != tt.want {
				t.Errorf("parseSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
