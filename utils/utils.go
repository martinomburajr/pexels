package utils

//go:generate mockgen --destination=../mocks/mock_utils.go --package mocks github.com/martinomburajr/pexels/utils BackgroundChanger,Filer,Rander,Utilizer

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
)
// Utilizer embodies all the different interfaces
type Utilizer interface {
	Filer
	BackgroundChanger
	Rander
}

// Filer is responsible for behaviour that performs I/O to files
type Filer interface {
	// WriteToFile writes data to a given filepath, error handling is left to the implementor
	WriteToFile(filepath string, data []byte) error
}

// BackgroundChanger is involved in calling the OS to perform writes
type BackgroundChanger interface {
	// ChangeUbuntuBackground performs desktop background change. This method can be reimplemented to suite changes in OS behaviors. error handling is left to the implementor
	ChangeUbuntuBackground(filepath string) error
}

// Rander contains functions that perform the creation of random data
type Rander interface {
	// RandInt returns a random integer. You can specify a non-negative max.
	RandInt(max int) int
	// RandBytes returns a byte array of random values.
	RandBytes(size int) []byte
}

// Utils acts as a container type for the methods within this file
type Utils struct {}

//WriteToFile requires the full path to file as well as file extension e.g. ~./pexels/pictures/snow.jpg, as well as a byte array for the data
func (w *Utils) WriteToFile(filepath string, data []byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	err = ioutil.WriteFile(filepath, data, 0755)
	if err != nil {
		return err
	}
	return nil
}

//ChangeUbuntuBackground works on Ubuntu. Must be the full filepath
func (w *Utils) ChangeUbuntuBackground(filepath string) error {
	//gsettings set org.gnome.desktop.background picture-uri file:///path/to/your/image.png from https://askubuntu.com/a/156722
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (w *Utils) RandInt(max int) int {
	if max < 2{
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
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}



