package utils

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"

	"github.com/gorilla/mux"
)

func TestRandBytes(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
	}{
		{"size -1", args{-1}},
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
		if get := r.Header.Get("Authorization"); get == "" {
			//t.Fatalf("Authorization header not passed | got %s", get)
			http.NotFound(w, r)
		}
		u := Utils{}
		randBytes := u.RandBytes(10)
		w.Write(randBytes)
	})

	//Returns a non-existent body that will cause the read to the body to fail.
	//Also states there is an HTTP 500 error
	serverErr := func(w http.ResponseWriter, r *http.Request) {
		if get := r.Header.Get("Authorization"); get == "" {
			//t.Fatalf("Authorization header not passed | got %s", get)
			http.NotFound(w, r)
		}

		w.Header().Set("Content-Length", "1")
		w.WriteHeader(http.StatusInternalServerError) // error header > 399
		w.Write(nil)
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
		name string
		w    *Utils
		args args
		//want    []byte
		wantErr bool
	}{
		{"request-error | bad url", &u, args{"@:/fd:", "some-key"}, true},
		{"request-error | bad request - unsupported protocol", &u, args{"/badapi", "some-key"}, true},
		{"request-error | bad request - bad scheme", &u, args{srv.URL[3:len(srv.URL)] + "/goodapi", "some-key"}, true},
		{"request - unreadable content - unexpected EOF", &u, args{srv.URL + "/badapi", "some-key"}, true},
		{"request-error | bad request", &u, args{srv.URL + "/goodapi", "some-key"}, false},
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

func TestUtilsRandInt(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{"max positive integer > 1", args{max: 10}},
		{"integer 1 should return 0", args{max: 1}},
		{"integer 0 should return 0", args{max: 0}},
		{"integer less than 0", args{max: -1}},
	}
	u := Utils{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := u.RandInt(tt.args.max)
			if (tt.args.max <= 1) && (got != 0) {
				t.Errorf("if args == 1 or 0 must return 0")
			}
			if (got > tt.args.max && tt.args.max > 0) || (got < 0) {
				t.Errorf("got %d is out of range max %d", got, tt.args.max)
			}
		})
	}
}

func TestUtilsWriteToFile(t *testing.T) {
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
		{"writing to bad filepath", args{":", rand1}, true},
		{"writing 0 bytes", args{"testdata/1.jpg", rand1}, false},
		{"writing 10 bytes", args{"testdata/2.jpg", rand2}, false},
		{"writing 100 bytes", args{"testdata/3.jpg", rand3}, false},
	}
	u := Utils{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := u.WriteToFile(tt.args.filepath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func generateRandomBytes(size int) []byte {
	if size < 0 {
		token := make([]byte, 100)
		rand.Read(token)
		return token
	}
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func fakeExecCommand(command string, s ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, s...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

const ubuntuChangeBackgroundResult = "ok"

func testUtilsChangeBackground(t *testing.T, filepath string, wantErr bool) {
	t.Helper()
	u := Utils{}
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := u.ChangeUbuntuBackground(filepath)
	if wantErr {
		if err != nil {

		} else {
			t.Fatalf("Expected error, got %#v", err)
		}
	} else {
		if string(out) != ubuntuChangeBackgroundResult {
			t.Fatalf("Expected %q, got %q", ubuntuChangeBackgroundResult, out)
		}
	}
}

func TestUtils_ChangeUbuntuBackground(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"bad filepath", args{"fakepath"}, true},
		{"bad file extension", args{"testdata/notimage"}, true},
		{"good file extension no bytes", args{"testdata/0.jpg"}, false},
		{"good file extension bytes", args{"testdata/3.jpg"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testUtilsChangeBackground(t, tt.args.filepath, tt.wantErr)
		})
	}
}
