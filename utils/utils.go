package utils

import (
	"github.com/martinomburajr/pexels/auth"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
)

//WriteImageToFile requires the full path to file as well as file extension e.g. ~./pexels/pictures/snow.jpg, as well as a byte array for the data
func WriteImageToFile(filepath string, data []byte) error {
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
func ChangeUbuntuBackground(filepath string) error {
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

//Returns an image
//func GetRandomImageWQuery() {
//
//}
//
//
//
//
func RandIntBetween(max int) int {
	return rand.Intn(max-1) + 1
}

//ParseRequest parses the request for a picture
func ParseRequest(urlWSize string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, urlWSize, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add(http.CanonicalHeaderKey("Authorization"), auth.PexelSession.API_KEY)

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



