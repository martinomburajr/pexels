package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func PexelsServer(t *testing.T) *http.Server {
//	t.Helper()
//
//	mux := mux.Router{}
//	mux.Methods(http.Get(""))
//}

func testReadSizesFile(t *testing.T) string {
	t.Helper()
	bytes, err := ioutil.ReadFile("testdata/sizes")
	if err != nil {
		t.Fail()
	}
	return string(bytes)
}

func TestGetSizesHandler(t *testing.T) {
	sizes := testReadSizesFile(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/sizes", nil)

	handler := http.HandlerFunc(GetSizesHandler)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if w.Body.String() != sizes {
		t.Errorf("handler body is not identical to file\n\nbody:\n%s\nfile:\n%s", w.Body.String(), sizes)
	}
}

func TestGetRandomHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/rand", nil)
	if err != nil {
		t.Error(err)
	}
	recorder := httptest.NewRecorder()

	body := recorder.Result().Body
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		t.Error(err)
	}

	if len(data) < 1 {
		t.Errorf("returned no data")
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"basic call", args{recorder, request}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//getRandomImager := pexels.GetRandomPexeler.GetRandomImage
			//getRandomImager(nil, "large")
			//handler := http.HandlerFunc(GetRandomHandler(getRandomImager(nil, "large")))
			//handler.ServeHTTP(recorder, request)
			//
			//status := recorder.Code
			//if !tt.wantErr {
			//	if status != http.StatusOK {
			//		t.Errorf("Status not OK")
			//	}
			//}
		})
	}
}

func TestHealthCheckHandler(t *testing.T) {
	type args struct {
		r *http.Request
		bodyWant string
	}
	tests := []struct {
		name string
		args args
	}{
		{ "simple request", args{r: httptest.NewRequest(http.MethodGet, "/hc", nil), bodyWant: `{"status":"ok"}`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			http.HandlerFunc(HealthCheckHandler).ServeHTTP(w, tt.args.r)

			if status := w.Code; status != http.StatusOK {
				t.Errorf("status is not OK | got %d", status)
			}

			defer w.Result().Body.Close()
			data, err := ioutil.ReadAll(w.Result().Body)

			if err != nil {
				t.Errorf("error reading body %v", err)
			}

			strResponse := string(data)
			if strResponse != tt.args.bodyWant {
				t.Errorf("invalid body, got %s | want %s", strResponse, tt.args.bodyWant)
			}
		})
	}
}
