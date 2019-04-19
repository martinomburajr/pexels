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
	err = utils.WriteToFile(filepath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := fmt.Sprintf(
		"Filename: %s\n" +
			"Directory: %s\n" +
			"Size: %d bytes", fmt.Sprintf("%d.jpg\n", pexel.ID), filepath, len(data))

	utils.ChangeUbuntuBackground(filepath)
	if err != nil {
		log.Print("failed to change ubuntu background")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func GetRandomHandler(w http.ResponseWriter, r *http.Request) {
	pexel := pexels.PexelPhoto{}

	imgSize := r.URL.Query().Get("size")

	data, err := pexel.GetRandomImage(imgSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(data) < 1 {
		return
	}

	filepath := fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(), pexel.ID)
	err = utils.WriteToFile(filepath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := fmt.Sprintf(
		"Filename: %s\n" +
		"Directory: %s\n" +
		"Size: %d bytes", fmt.Sprintf("%d.jpg\n", pexel.ID), filepath, len(data))

	err = utils.ChangeUbuntuBackground(filepath)
	if err != nil {
		log.Print("failed to change ubuntu background")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}