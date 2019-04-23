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
