package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
