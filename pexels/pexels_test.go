package pexels

import (
	"github.com/martinomburajr/pexels/auth"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var pps1 PexelPhotoSource = PexelPhotoSource{"original1", "large1", "large2x1", "medium1", "small1", "portrait1", "landscape1", "tiny1"}
var pps2 PexelPhotoSource = PexelPhotoSource{"original2", "large2", "large2x2", "medium2", "small2", "portrait2", "landscape2", "tiny2"}

var pp1 PexelPhoto = PexelPhoto{14343, 1080, 1920, "https://pexels/some/fake/url", "martin", pps1}
var pp2 PexelPhoto = PexelPhoto{4433, 1080, 1920, "https://pexels/some/fake/url2", "carla", pps2}

//func TestPexelPhoto_GetBySize(t *testing.T) {
//	tests := []struct {
//		name   string
//		fields PexelPhoto
//		args   string
//		want   string
//	}{
//		{"empty string as size value", pp1, "", "large1"},
//		{"random unrelated text", pp1, "asbngaojubgao;ga", "large1"},
//		{"normal string: small", pp1, "small", "small1"},
//		{"normal string: landscape", pp1, "landscape", "landscape1"},
//		{"normal string: large2x", pp2, "laRge2x", "large2x2"},
//		{"normal string: medium", pp2, "medium", "medium2"},
//		{"normal string: tiny", pp2, "tiny", "tiny2"},
//		{"normal string: large", pp2, "large", "large2"},
//		{"mixed case string: original", pp2, "origINAL", "original2"},
//		{"misspelled string: large2x", pp2, "laRg2x", "large2"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pi := &PexelPhoto{
//				ID:           tt.fields.ID,
//				Width:        tt.fields.Width,
//				Height:       tt.fields.Height,
//				URL:          tt.fields.URL,
//				Photographer: tt.fields.Photographer,
//				Source:       tt.fields.Source,
//			}
//			got, err := pi.Get(tt.args.id, tt.args.size)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("PexelPhoto.Get() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("PexelPhoto.Get() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
		want    int
		want1   []byte
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
			got, got1, err := pi.GetRandomImage(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("PexelPhoto.GetRandomImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PexelPhoto.GetRandomImage() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PexelPhoto.GetRandomImage() got1 = %v, want %v", got1, tt.want1)
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
//
//func TestPexelPhoto_Get(t *testing.T) {
//	type fields struct {
//		ID           int
//		Width        int
//		Height       int
//		URL          string
//		Photographer string
//		Source       PexelPhotoSource
//	}
//	type args struct {
//		client  *http.Client
//		session *auth.PexelSessionObj
//		id      int
//		size    string
//	}
//	tests := []struct {
//		name         string
//		fields       fields
//		args         args
//		testHandler1 http.HandlerFunc
//		testHandler2 http.HandlerFunc
//		want         []byte
//		wantErr      bool
//	}{
//		{"apikey empty",
//			fields{},
//			args{http.DefaultClient, &auth.PexelSessionObj{API_KEY: ""}, 0, ""},
//			func(w http.ResponseWriter, r *http.Request) {
//				if r.URL.String() != fmt.Sprintf("%s%s/%d", BaseURL, "photos", 0) {
//					w.WriteHeader(http.StatusBadRequest)
//					w.Write(nil)
//					return
//				}
//				w.WriteHeader(http.StatusOK)
//				w.Write([]byte("ok"))
//			},
//			func(w http.ResponseWriter, r *http.Request) {
//				if r.URL.String() != fmt.Sprintf("%s%s/%d", BaseURL, "photos", 0) {
//					w.WriteHeader(http.StatusBadRequest)
//					w.Write(nil)
//					return
//				}
//				w.WriteHeader(http.StatusOK)
//				w.Write([]byte("ok"))
//			},
//			nil,
//			true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pi := &PexelPhoto{
//				ID:           tt.fields.ID,
//				Width:        tt.fields.Width,
//				Height:       tt.fields.Height,
//				URL:          tt.fields.URL,
//				Photographer: tt.fields.Photographer,
//				Source:       tt.fields.Source,
//			}
//
//			srv := httptest.NewServer(tt.testHandler1)
//			defer srv.Close()
//
//			srv2 := httptest.NewServer(tt.testHandler2)
//			defer srv2.Close()
//
//			got, err := pi.Get(tt.args.session, tt.args.id, tt.args.size)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("PexelPhoto.Get() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("PexelPhoto.Get() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestGetImage(t *testing.T) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: http.DefaultTransport,
	}

	session := &auth.PexelSessionObj{
		API_KEY: "some-api-key",
	}

	//goodHandler := func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//	w.Write([]byte("some image in here"))
	//}
	//errHandler := func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write([]byte("some error"))
	//}
	//nilHandler := func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write(nil)
	//}

	type args struct {
		client   *http.Client
		session  *auth.PexelSessionObj
		imageURL string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"empty client",  args{nil, session, ""}, nil, true},
		{"client - empty image url",  args{client, session, ""}, nil, true},
		{"client - error image url",  args{client, session, ":"}, nil, true},
		{"client - error image url",  args{client, session, "http://someapi.com/somevalue"}, nil, false},

		//{"client - request - bad do",  args{client, session, "bad"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//mux := http.DefaultServeMux
			//mux.HandleFunc("/err" , errHandler)
			//mux.HandleFunc("/nil" , nilHandler)
			//mux.HandleFunc("/good", goodHandler)
			//
			//srv := httptest.NewServer(mux)
			//srvURL := srv.URL

			got, err := GetImage(tt.args.client, tt.args.session, tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
