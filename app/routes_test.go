package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/mocks"
	"github.com/martinomburajr/pexels/utils"
)

func TestServer_Routes(t *testing.T) {
	tests := []struct {
		name     string
		req      *http.Request
		wantCode int
		wantErr  bool
	}{
		{"test / handle", httptest.NewRequest(http.MethodGet, "/", nil), http.StatusOK, false},
		{"test // {non existent}", httptest.NewRequest(http.MethodGet, "//", nil), http.StatusMovedPermanently, false},
		{"test /a handle {non existent}", httptest.NewRequest(http.MethodGet, "/a", nil), http.StatusNotFound, true},
		{"test /hc handle", httptest.NewRequest(http.MethodGet, "/hc", nil), http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			s := Server{}
			s.Router = *s.Routes()
			s.Router.ServeHTTP(w, tt.req)

			if tt.wantErr && (w.Code < 400) {
				t.Fatalf("wanted error i.e  x < 400 | got %d, want %d", w.Code, tt.wantCode)
			}
			if w.Code != tt.wantCode {
				t.Fatalf("path is NOT ok | got %d, want %d", w.Code, tt.wantCode)
			}
		})
	}
}

func TestServer_GetRandomHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//Setup all the mock interfaces
	mockPexeler := mocks.NewMockPexeler(mockCtrl)
	mockUtils := mocks.NewMockUtilizer(mockCtrl)

	//Setup a request that will be used for all tests. It is not a singleton.
	request := httptest.NewRequest(http.MethodGet, "/rand", nil)
	u := utils.Utils{}

	randomBytes := u.RandBytes(10)
	zeroBytes := u.RandBytes(0)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		pexelerMock    func() *gomock.Call
		filerMock      func() *gomock.Call
		backgroundMock func() *gomock.Call
		wantErr        bool
	}{
		{"zero values", args{request},
			func() *gomock.Call { return mockPexeler.EXPECT().GetRandomImage("").Return(0, zeroBytes, nil).Times(1) },
			func() *gomock.Call {
				return mockUtils.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(""), 0), zeroBytes).Return(nil).Times(0)
			},
			func() *gomock.Call { return mockUtils.EXPECT().ChangeBackground("").Times(0).Return(nil) },
			false},
		{"error retrieving image", args{request},
			func() *gomock.Call {
				return mockPexeler.EXPECT().GetRandomImage("").Return(0, zeroBytes, errors.New("error")).Times(1)
			},
			func() *gomock.Call {
				return mockUtils.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(""), 0), zeroBytes).Return(nil).Times(0)
			},
			func() *gomock.Call { return mockUtils.EXPECT().ChangeBackground("").Times(0).Return(nil) },
			true},
		{"retrieved an image - file error ", args{request},
			func() *gomock.Call {
				return mockPexeler.EXPECT().GetRandomImage("").Return(0, randomBytes, nil).Times(1)
			},
			func() *gomock.Call {
				return mockUtils.EXPECT().WriteToFile(gomock.Eq(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(""), 0)), randomBytes).Return(errors.New("")).Times(1)
			},
			func() *gomock.Call { return mockUtils.EXPECT().ChangeBackground("").Return(nil).Times(0) },
			true},
		{"retrieved an image - no file error - background change error", args{request},
			func() *gomock.Call {
				return mockPexeler.EXPECT().GetRandomImage("").Return(1, randomBytes, nil).Times(1)
			},
			func() *gomock.Call {
				return mockUtils.EXPECT().WriteToFile(gomock.Eq(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(""), 1)), randomBytes).Return(nil).Times(1)
			},
			func() *gomock.Call {
				return mockUtils.EXPECT().ChangeBackground("").Return(errors.New("")).Times(1)
			},
			true},
		//{"retrieved an image - no file error - no background change error", args{request},
		//	func() *gomock.Call { return mockPexeler.EXPECT().GetRandomImage("").Return(u.RandBytes(10000000), nil).Times(1) },
		//	func() *gomock.Call { return mockUtils.EXPECT().WriteToFile("", "").Return(nil).Times(1) },
		//	func() *gomock.Call { return mockUtils.EXPECT().ChangeBackground("").Times(1).Return(nil) },
		//	false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Perform the calls
			tt.pexelerMock()
			tt.filerMock()
			//tt.backgroundMock()

			s := &Server{
				Router:   mux.Router{},
				PexelsDB: mockPexeler,
				Utilizer: mockUtils,
			}

			w := httptest.NewRecorder()

			http.HandlerFunc(s.GetRandomHandler).ServeHTTP(w, tt.args.r)

			if tt.wantErr {
				if status := w.Code; status < 400 {
					t.Fatalf("expected code error: status got %d", status)
				}

			} else {
				if status := w.Code; status != http.StatusOK {
					t.Fatalf("status not OK - got %d", status)
				}
			}
		})
	}
}

func TestServer_GetSizesHandler(t *testing.T) {
	sizes := testReadSizesFile(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/sizes", nil)

	srv := Server{}
	srv.Routes()

	handler := http.HandlerFunc(srv.GetSizesHandler)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if strings.Contains(string(w.Body.String()), sizes) {
		t.Errorf("handler body is not identical to file\n\nbody:\n%s\nfile:\n%s", w.Body.String(), sizes)
	}
}

func testReadSizesFile(t *testing.T) string {
	t.Helper()
	bytes, err := ioutil.ReadFile("testdata/sizes")
	if err != nil {
		t.Fail()
	}
	return string(bytes)
}

func TestServer_GetPexelHandler(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUtilizer := mocks.NewMockUtilizer(controller)
	mockPexeler := mocks.NewMockPexeler(controller)
	srv := &Server{
		Utilizer: mockUtilizer,
		PexelsDB: mockPexeler,
	}

	minPhotoSize := 1024

	smallBytes := testingGenerateBytes(8, t)
	largeBytes := testingGenerateBytes(minPhotoSize + 1, t)

	type args struct {
		r *http.Request
	}
	tests := []struct {
		name                 string
		args                 args
		pexelerMock          func() *gomock.Call
		filerMock            func() *gomock.Call
		changeBackgroundMock func() *gomock.Call
		fields               *Server
		wantErr              bool
	}{
		// Because we are using a gorilla/mux and registered the /new route with an id. Any path without an id is automatically forfeited.
		{"error no request id",
			args{httptest.NewRequest(http.MethodGet, "/new", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile("", nil).Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground("").Times(0) },
			srv, true},
			{"request id - invalid characters",
			args{httptest.NewRequest(http.MethodGet, "/new/gfgd", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile("", nil).Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground("").Times(0) },
			srv, true},
			{"request id - valid characters - server get error",
			args{httptest.NewRequest(http.MethodGet, "/new/0", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Return(nil, errors.New("error")).Times(1) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile("", nil).Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground("").Times(0) },
			srv, true},
		{"request id - valid characters - server no error no bytes",
			args{httptest.NewRequest(http.MethodGet, "/new/0", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Return(nil, nil).Times(1) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile("", nil).Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground("").Times(0) },
			srv, true},
			{"request id - valid characters - server no error small bytes",
			args{httptest.NewRequest(http.MethodGet, "/new/0", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Return(smallBytes, nil).Times(1) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile("", nil).Times(0) },
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground("").Times(0) },
			srv, true},
		{"request id - valid characters - server get ok bytes - err write-file ",
			args{httptest.NewRequest(http.MethodGet, "/new/0", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Return(largeBytes, nil).Times(1) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(config.GetHomeDir()), 0), largeBytes).Return(errors.New("error")).Times(1)},
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground("").Times(0) },
			srv, true},
		{"request id - valid characters - server get ok bytes - write-file - err change background ",
			args{httptest.NewRequest(http.MethodGet, "/new/0", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Return(largeBytes, nil).Times(1) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(config.GetHomeDir()), 0), largeBytes).Return(nil).Times(1)},
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(config.GetHomeDir()), 0)).Return(nil, errors.New("error")).Times(1) },
			srv, true},
		{"request id - valid characters - server get ok bytes - write-file -  change background ",
			args{httptest.NewRequest(http.MethodGet, "/new/0", nil)},
			func() *gomock.Call { return mockPexeler.EXPECT().Get(0, "").Return(largeBytes, nil).Times(1) },
			func() *gomock.Call { return mockUtilizer.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(config.GetHomeDir()), 0), largeBytes).Return(nil).Times(1)},
			func() *gomock.Call { return mockUtilizer.EXPECT().ChangeBackground(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(config.GetHomeDir()), 0)).Return([]byte("ok"), nil).Times(1) },
			srv, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pexelerMock()
			tt.filerMock()
			tt.changeBackgroundMock()

			s := &Server{
				PexelsDB: tt.fields.PexelsDB,
				Router:   tt.fields.Router,
				Utilizer: tt.fields.Utilizer,
			}

			w := httptest.NewRecorder()
			s.Routes().ServeHTTP(w, tt.args.r)

			if tt.wantErr {
				if w.Code < 400 {
					t.Fatalf(" wanted an error  want: <400 | got %d", w.Code)
				}
			} else {
				if status := w.Code; status != http.StatusOK {
					t.Fatalf("status not OK - got %d", status)
				}
			}
		})
	}
}

func testingGenerateBytes(size int, t *testing.T) []byte{
	t.Helper()
	arr := make([]byte, size)
	rand.Read(arr)
	return arr
}