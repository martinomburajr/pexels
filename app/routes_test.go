package app

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/mocks"
	"github.com/martinomburajr/pexels/utils"
	"net/http"
	"net/http/httptest"
	"testing"
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
			func() *gomock.Call { return mockUtils.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(), 0), zeroBytes).Return(nil).Times(0) },
			func() *gomock.Call { return mockUtils.EXPECT().ChangeUbuntuBackground("").Times(0).Return(nil) },
			false},
		{"error retrieving image", args{request},
			func() *gomock.Call { return mockPexeler.EXPECT().GetRandomImage("").Return(0, zeroBytes, errors.New("error")).Times(1) },
			func() *gomock.Call { return mockUtils.EXPECT().WriteToFile(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(), 0), zeroBytes).Return(nil).Times(0) },
			func() *gomock.Call { return mockUtils.EXPECT().ChangeUbuntuBackground("").Times(0).Return(nil) },
			true},
		{"retrieved an image - file error ", args{request},
			func() *gomock.Call { return mockPexeler.EXPECT().GetRandomImage("").Return(0, randomBytes, nil).Times(1) },
			func() *gomock.Call { return mockUtils.EXPECT().WriteToFile(gomock.Eq(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(), 0)), randomBytes).Return(errors.New("")).Times(1) },
			func() *gomock.Call { return mockUtils.EXPECT().ChangeUbuntuBackground("").Return(nil).Times(0) },
			true},
		{"retrieved an image - no file error - background change error", args{request},
			func() *gomock.Call { return mockPexeler.EXPECT().GetRandomImage("").Return(1, randomBytes, nil).Times(1) },
			func() *gomock.Call { return mockUtils.EXPECT().WriteToFile(gomock.Eq(fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(), 1)), randomBytes).Return(nil).Times(1) },
			func() *gomock.Call { return mockUtils.EXPECT().ChangeUbuntuBackground("").Return(errors.New("")).Times(1) },
			true},
		//{"retrieved an image - no file error - no background change error", args{request},
		//	func() *gomock.Call { return mockPexeler.EXPECT().GetRandomImage("").Return(u.RandBytes(10000000), nil).Times(1) },
		//	func() *gomock.Call { return mockUtils.EXPECT().WriteToFile("", "").Return(nil).Times(1) },
		//	func() *gomock.Call { return mockUtils.EXPECT().ChangeUbuntuBackground("").Times(1).Return(nil) },
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
