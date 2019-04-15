package pexels

//go:generate mockgen --destination=../mocks/mock_getter.go --package mocks github.com/martinomburajr/gopexels/pexels Getter

import (
	"encoding/json"
	"fmt"
	"github.com/martinomburajr/pexels/utils"
	"strings"
)

const (
	ImageSizeOriginal  = "original"
	ImageSizeLarge     = "large"
	ImageSizeLarge2x   = "large2x"
	ImageSizeMedium    = "medium"
	ImageSizeSmall     = "small"
	ImageSizePortrait  = "portrait"
	ImageSizeLandscape = "landscape"
	ImageSizeTiny      = "tiny"
	BaseURL            = "https://api.pexels.com/v1/"
	URLCurated         = "curated"
)

//ImageSizes represents a set of image sizes that pexels uses
var ImageSizes = []string{ImageSizeOriginal, ImageSizeLarge, ImageSizeLarge2x, ImageSizeMedium, ImageSizeSmall, ImageSizePortrait, ImageSizeLandscape, ImageSizeTiny}

// Getter belongs to any types that must retrieve an item based on an id.
type Getter interface {
	// Given an id or resource locator, Get implements the functionality of retrieving an item id doesnt necessarily need to be a standardized id e.g. for a database record,
	// It can be any string value that identifies an item as unique.
	// Returns an error and nil bytes if there is an error. A nil error and a non-nil bytes array could represent an error returned in the form of bytes or a successful retrieval
	Get(id string) ([]byte, error)

	//GetR will retrieve a random element
	GetR() ([]byte, error)
}

//PexelImageRespoonse represents a response from the server regarding an image request
type PexelImageResponse struct {
	Page         int          `json:"page,omitempty"`
	PerPage      int          `json:"per_page,omitempty"`
	TotalResults int          `json:"per_page,omitempty"`
	URL          string       `json:"url"`
	NextPage     string       `json:"next_page"`
	Photos       []PexelPhoto `json:"photos"`
}

//PexelPhoto represents the information of photo
type PexelPhoto struct {
	ID           int           `json:"id,omitempty"`
	Width        int              `json:"width,omitempty"`
	Height       int              `json:"height,omitempty"`
	URL          string           `json:"url,omitempty"`
	Photographer string           `json:"photographer,omitempty"`
	Source       PexelPhotoSource `json:"src,omitempty"`
}

//PexelPhotoSource represents a photo source embedded within the PexelPhoto
type PexelPhotoSource struct {
	Original  string `json:"original,omitempty"`
	Large     string `json:"large,omitempty"`
	Large2x   string `json:"large2x,omitempty"`
	Medium    string `json:"medium,omitempty"`
	Small     string `json:"small,omitempty"`
	Portrait  string `json:"portrait,omitempty"`
	Landscape string `json:"landscape,omitempty"`
	Tiny      string `json:"tiny,omitempty"`
}

// PexelPhoto implementation of Getter that retrieves a random image based on its size.
func (pi *PexelPhoto) Get(id string) ([]byte, error) {
	urll := BaseURL + "photos/" + id
	data, err := utils.ParseRequest(urll)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, pi)
	if err != nil {
		return nil, err
	}

	data2, err := utils.ParseRequest(pi.Source.Original)
	return data2, nil
}

// PexelPhoto implementation of Getter that retrieves a random image based on its size.
func (pi *PexelPhoto) GetBySize(size string) ([]byte, error) {
	s := parseSize(size)
	return utils.ParseRequest(s)
}

//parseSize obtains the size arg and if it is empty returns the ImageSizeLarge
func parseSize(size string) string {
	lower := strings.ToLower(size)
	for _, v := range ImageSizes {
		if lower == strings.ToLower(v) {
			return lower
		}
	}
	size = ImageSizeLarge
	return size
}

//GetRandomImage returns a random image from the Pexel API
func(pi *PexelPhoto) GetRandomImage() ([]byte, error) {
	randomInt := utils.RandIntBetween(1000)
	urll := fmt.Sprintf("%s%s?per_page=%d&page=%d", BaseURL, URLCurated, 1, randomInt)

	data, err := utils.ParseRequest(urll)
	if err != nil {
		return nil, err
	}

	var pr PexelImageResponse
	err = json.Unmarshal(data, &pr)
	if err != nil {
		return nil, err
	}

	*pi = pr.Photos[0]

	data2, err := utils.ParseRequest(pi.Source.Original)
	return data2, nil
}
