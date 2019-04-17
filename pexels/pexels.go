package pexels

//go:generate mockgen --destination=../mocks/mock_getter.go --package mocks github.com/martinomburajr/gopexels/pexels Getter

import (
	"encoding/json"
	"fmt"
	"github.com/martinomburajr/pexels/utils"
	"strings"
)

const (
	//@todo get exact description of sizes
	//ImageSizeOriginal represents the original size. Typically the largest with the best quality
	ImageSizeOriginal  = "original"
	//ImageSizeLarge is a large photo
	ImageSizeLarge     = "large"
	//ImageSizeLarge2x double the resolution of the largest
	ImageSizeLarge2x   = "large2x"
	//ImageSizeMedium medium photo
	ImageSizeMedium    = "medium"
	//ImageSizeSmall small photo (lacks in resolution)
	ImageSizeSmall     = "small"
	//ImageSizePortrait portrait mode. This image is usually cropped to fit that size
	ImageSizePortrait  = "portrait"
	//ImageSizeLandscape landscape sized photo
	ImageSizeLandscape = "landscape"
	//ImageSizeLandscape tiny photo
	ImageSizeTiny      = "tiny"
	//BaseURL is the base URL to the API
	BaseURL            = "https://api.pexels.com/v1/"
	//URLCurated is a path to the curated section within pexels. According to pexels ... We add at least one new photo per hour to our curated list so that you get a changing selection of trending photos. For more information about the request parameters and response structure have a look at the search method above.
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
	ID           int              `json:"id,omitempty"`
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
func (pi *PexelPhoto) Get(id, size string) ([]byte, error) {
	urll := BaseURL + "photos/" + id
	data, err := utils.ParseRequest(urll)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, pi)
	if err != nil {
		return nil, err
	}

	s := parseSize(size)
	bySize := pi.GetBySize(s)
	data2, err := utils.ParseRequest(bySize)
	return data2, nil
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
func (pi *PexelPhoto) GetRandomImage(size string) ([]byte, error) {
	randomInt := utils.RandIntBetween(1000)
	urll := fmt.Sprintf("%s%s?per_page=%d&page=%d", BaseURL, URLCurated, 1, randomInt)

	data, err := utils.ParseRequest(urll)
	if err != nil {
		return nil, err
	}
	s := parseSize(size)

	var pr PexelImageResponse
	err = json.Unmarshal(data, &pr)
	if err != nil {
		return nil, err
	}

	*pi = pr.Photos[0]

	bySize := pi.GetBySize(s)

	data2, err := utils.ParseRequest(bySize)
	return data2, nil
}

//GetBySize returns the exact size based url based on the size parameter
func (pi *PexelPhoto) GetBySize(size string) string {
	switch size {
	case ImageSizeLarge2x:
		return pi.Source.Large2x
	case ImageSizeLarge:
		return pi.Source.Large
	case ImageSizeLandscape:
		return pi.Source.Landscape
	case ImageSizeMedium:
		return pi.Source.Medium
	case ImageSizeOriginal:
		return pi.Source.Original
	case ImageSizeSmall:
		return pi.Source.Small
	case ImageSizeTiny:
		return pi.Source.Tiny
	default:
		return pi.Source.Large
	}
	return pi.Source.Large
}
