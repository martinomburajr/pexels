package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/pexels"
	"github.com/martinomburajr/pexels/utils"
	"net/http"
	"strconv"
)

type Server struct {
	PexelsDB pexels.Pexeler
	Router   mux.Router
	Utilizer utils.Utilizer
}

// MinImageBytes states minimum size the response from retrieving an image can be. Responses smaller than this will not be accepted.
const MinImageBytes = 1024

// routes returns a gorilla/mux Router which is a valid Router/Handler that can be served. Refactoring this into a function makes it testable.
func (s *Server) Routes() *mux.Router {
	s.Router.Methods(http.MethodGet).Path("/").HandlerFunc(s.HealthCheckHandler)
	s.Router.Methods(http.MethodGet).Path("/hc").HandlerFunc(s.HealthCheckHandler)
	s.Router.Methods(http.MethodGet).Path("/new/{id}").HandlerFunc(s.GetPexelHandler)
	s.Router.Methods(http.MethodGet).Path("/rand").HandlerFunc(s.GetRandomHandler)
	s.Router.Methods(http.MethodGet).Path("/sizes").HandlerFunc(s.GetSizesHandler)

	return &s.Router
}

// GetPexelHandler returns a single photo based on an id.
func (s *Server) GetPexelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	// Not checking if idString is empty.
	// Because we are using a gorilla/mux and registered the /new route with an id. Any path without an id is automatically forfeited.

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(w, "invalid characters: id parameter in url should be an integer", http.StatusBadRequest)
		return
	}

	imgSize := r.URL.Query().Get("size")
	data, err := s.PexelsDB.Get(int(id), imgSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if data == nil {
		//@todo change return message in router/GetPexelHandler to no bytes
		http.Error(w, "pexels server returned small bytes", http.StatusInternalServerError)
		return
	}

	if len(data) < MinImageBytes  {
		http.Error(w, "pexels server returned small bytes", http.StatusInternalServerError)
		return
	}

	homeDir := config.GetHomeDir()
	filepath := fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(homeDir), id)
	err = s.Utilizer.WriteToFile(filepath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf(
		"Filename: %s\n"+
			"Directory: %s\n"+
			"Size: %d bytes", fmt.Sprintf("%d.jpg\n", id), filepath, len(data))

	data, err = s.Utilizer.ChangeBackground(filepath)
	if err != nil {
		http.Error(w, "failed to change background " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

// GetRandomHandler will download a random image from the curated list in pexels.
func (s *Server) GetRandomHandler(w http.ResponseWriter, r *http.Request) {
	imgSize := r.URL.Query().Get("size")

	id, data, err := s.PexelsDB.GetRandomImage(imgSize)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	if data == nil {
		http.Error(w, "pexels server returned no bytes", http.StatusInternalServerError)
		return
	}
	if len(data) < MinImageBytes  {
		http.Error(w, "pexels server returned small bytes", http.StatusInternalServerError)
		return
	}

	homeDir := config.GetHomeDir()
	filepath := fmt.Sprintf("%s/%d.jpg", config.CanonicalPicturePath(homeDir), id)
	err = s.Utilizer.WriteToFile(filepath, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf(
		"Filename: %s\n"+
			"Directory: %s\n"+
			"Size: %d bytes", fmt.Sprintf("%d.jpg\n", id), filepath, len(data))

	data, err = s.Utilizer.ChangeBackground(filepath)
	if err != nil {
		http.Error(w, "failed to change background " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

// HealthCheckHandler is a simple test to see that the router and server are able to pick up incoming requests and handle them appropriately.
func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

//GetSizesHandler returns information about all supported sizes
func (s *Server) GetSizesHandler(w http.ResponseWriter, r *http.Request) {
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


