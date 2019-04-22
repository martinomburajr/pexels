package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/pexels"
	"github.com/martinomburajr/pexels/utils"
	"log"
	"net/http"
)




//GetPexelHandler returns a single photo
func GetPexelHandler(w http.ResponseWriter, r *http.Request) {
	pexel := pexels.PexelPhoto{}
	u := utils.Utils{}

	vars := mux.Vars(r)
	id := vars["id"]

	imgSize := r.URL.Query().Get("size")

	data, err := pexel.Get(id, imgSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	filepath := fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(), pexel.ID)
	err = u.WriteToFile(filepath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := fmt.Sprintf(
		"Filename: %s\n" +
			"Directory: %s\n" +
			"Size: %d bytes", fmt.Sprintf("%d.jpg\n", pexel.ID), filepath, len(data))

	err = u.ChangeUbuntuBackground(filepath)
	if err != nil {
		log.Print("failed to change ubuntu background")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}



//GetSizesHandler returns information about all supported sizes
func GetSizesHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprint("\nOriginal - The size of the original image is given with the attributes width and height.\n" +
		"Large - This image has a maximum width of 940px and a maximum height of 650px. It has the aspect ratio of the original image.\n" +
		"Large2x - This image has a maximum width of 1880px and a maximum height of 1300px. It has the aspect ratio of the original image.\n" +
		"Medium - This image has a height of 350px and a flexible width. It has the aspect ratio of the original image.\n" +
		"Small - This image has a height of 130px and a flexible width. It has the aspect ratio of the original image.\n" +
		"Portrait - This image has a width of 800px and a height of 1200px.\n" +
		"Landscape - This image has a width of 1200px and height of 627px.\n" +
		"Tiny - This image has a width of 280px and height of 200px.\n")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}