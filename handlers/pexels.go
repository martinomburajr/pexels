package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/pexels"
	"github.com/martinomburajr/pexels/utils"
	"net/http"
)

//GetPexelHandler returns a single photo
func GetPexelHandler(w http.ResponseWriter, r *http.Request) {
	pexel := pexels.PexelPhoto{}

	vars := mux.Vars(r)
	id := vars["id"]

	data, err := pexel.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	msg := fmt.Sprintf("%s/%s.jpg", config.CanonicalPicturePath(), id)
	err = utils.WriteImageToFile(msg, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := fmt.Sprintf("Filename: %s\n" +
		"Directory: %s\n" +
		"Size: %d bytes", id+".jpg", msg, len(data))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
