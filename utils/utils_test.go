package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
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


func TestUtils_ParseRequest(t *testing.T) {

	router := mux.NewRouter()
	serverOK := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if get := r.Header.Get("Authorization"); get  == "" {
			//t.Fatalf("Authorization header not passed | got %s", get)
			http.NotFound(w, r)
		}
		u := Utils{}
		randBytes := u.RandBytes(10)
		w.Write(randBytes)
	});
	serverErr := func(w http.ResponseWriter, r *http.Request) {
		if get := r.Header.Get("Authorization"); get == "" {
			//t.Fatalf("Authorization header not passed | got %s", get)
			http.NotFound(w, r)
		}

		w.Header().Set("Location", "")
		w.WriteHeader(http.StatusOK) // error header > 399
		w.Write([]byte{1,2})
		return
	}


	router.Methods(http.MethodGet).Path("/badapi").HandlerFunc(serverErr)
	router.Methods(http.MethodGet).Path("/goodapi").HandlerFunc(serverOK)

	srv := httptest.NewServer(router)
	defer srv.Close()

	u := Utils{}

	type args struct {
		url     string
		API_KEY string
	}
	tests := []struct {
		name    string
		w       *Utils
		args    args
		//want    []byte
		wantErr bool
	}{
		{"request-error | bad url", &u, args{"@:/fd:", "some-key"}, true },
		{"request-error | bad request - unsupported protocol", &u, args{"/badapi", "some-key"}, true },
		{"request-error | bad request - bad scheme", &u, args{srv.URL[3:len(srv.URL)]+"/badapi", "some-key"}, true },
		{"request - unreadable content", &u, args{srv.URL+"/badapi", "some-key"}, true },
		{"request-error | bad request", &u, args{srv.URL+"/goodapi", "some-key"}, false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {


			w := &Utils{}
			_, err := w.ParseRequest(tt.args.url, tt.args.API_KEY)
			if (err != nil) != tt.wantErr {
				t.Errorf("Utils.ParseRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
