package utils

import (
	"io/ioutil"
	"os"
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
