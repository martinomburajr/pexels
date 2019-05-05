package utils

//go:generate mockgen --destination=../mocks/mock_utils.go --package mocks github.com/martinomburajr/pexels/utils BackgroundChanger,Filer,Rander,Utilizer,HTTPRequester

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// Utilizer embodies all the different interfaces
type Utilizer interface {
	Filer
	BackgroundChanger
	Rander
	HTTPRequester
}

// Filer is responsible for behaviour that performs I/O to files
type Filer interface {
	// WriteToFile writes data to a given filepath, error handling is left to the implementor
	WriteToFile(filepath string, data []byte) error
}

// BackgroundChanger is involved in calling the OS to perform writes
type BackgroundChanger interface {
	// ChangeBackground performs desktop background change. This method can be reimplemented to suite changes in OS behaviors. error handling is left to the implementor
	ChangeBackground(filepath string) ([]byte, error)
}

type HTTPRequester interface {
	ParseRequest(url string, API_KEY string) ([]byte, error)
}

// Rander contains functions that perform the creation of random data
type Rander interface {
	// RandInt returns a random integer. You can specify a non-negative max.
	RandInt(max int) int
	// RandBytes returns a byte array of random values.
	RandBytes(size int) []byte
}

// Utils acts as a container type for the methods within this file
type Utils struct{}

//WriteToFile requires the full path to file as well as file extension e.g. ~./pexels/pictures/snow.jpg, as well as a byte array for the data
func (w *Utils) WriteToFile(filepath string, data []byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	return ioutil.WriteFile(filepath, data, 0755)
}

var execCommand = exec.Command

// ChangeBackground automatically changes the background on Ubuntu. Must be the full filepath
func (w *Utils) ChangeBackground(filepath string) ([]byte, error) {
	//gsettings set org.gnome.desktop.background picture-uri file:///path/to/your/image.png from https://askubuntu.com/a/156722

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	hasSuffix := false
	for _, v := range supportedimageformats {
		if strings.HasSuffix(strings.ToLower(file.Name()), strings.ToLower(v)) {
			hasSuffix = true
			break
		}
	}
	if !hasSuffix {
		return nil, fmt.Errorf("not a supported image format. Pexels supports the following: %v", supportedimageformats)
	}

	cmd := execCommand("gsettings", "set", "org.gnome.desktop.background", "picture-uri", file.Name())
	return []byte("ok"), cmd.Run()
}

// RandInt returns a number between [0, max)
func (w *Utils) RandInt(max int) int {
	if max < 2 {
		return 0
	}
	return rand.Intn(max-1) + 1
}

// RandBytes returns a byte array with random bytes in it. It will not accept sizes less than 0, and will default the value to 0.
func (w *Utils) RandBytes(size int) []byte {
	if size < 0 {
		size = 0
	}
	x := make([]byte, size)
	rand.Read(x)
	return x
}

//ParseRequest parses the request for a picture
func (w *Utils) ParseRequest(url string, API_KEY string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add(http.CanonicalHeaderKey("Authorization"), API_KEY)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

//common image formats that are allowed to be set by this applications. Contributors may need to add more.
var supportedimageformats = []string{".jpg", ".png", ".jpeg", ".bmp", ".gif"}
