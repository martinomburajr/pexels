package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/handlers"
	"github.com/martinomburajr/pexels/pexels"
	"github.com/martinomburajr/pexels/utils"
	"net/http"
)

type Server struct {
	PexelsDB pexels.Pexeler
	Router   mux.Router
	Utilizer utils.Utilizer
}

// routes returns a gorilla/mux Router which is a valid Router/Handler that can be served. Refactoring this into a function makes it testable.
func (s *Server) Routes() *mux.Router {
	s.Router.Methods(http.MethodGet).Path("/").HandlerFunc(s.HealthCheckHandler)
	s.Router.Methods(http.MethodGet).Path("/hc").HandlerFunc(s.HealthCheckHandler)
	s.Router.Methods(http.MethodGet).Path("/new/{id}").HandlerFunc(handlers.GetPexelHandler)
	s.Router.Methods(http.MethodGet).Path("/rand").HandlerFunc(s.GetRandomHandler)
	s.Router.Methods(http.MethodGet).Path("/sizes").HandlerFunc(s.GetSizesHandler)

	return &s.Router
}

// GetRandomHandler will download a random image from the curated list in pexels.
func (s *Server) GetRandomHandler(w http.ResponseWriter, r *http.Request) {
	//pexel := pexels.PexelPhoto{}
	imgSize := r.URL.Query().Get("size")

	id, data, err := s.PexelsDB.GetRandomImage(imgSize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(data) < 1 {
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

	data, err = s.Utilizer.ChangeUbuntuBackground(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
