package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/martinomburajr/gopexels/pexels"
	"log"
	"net/http"
	"os"
)

func init() {
	//Retrieve APIKEY
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", nil)
	r.HandleFunc("/new", GetPexelHandler)

	apikey := flag.String("key", "", "sets the api to be used by the individual")
	if *apikey == "" {
		log.Fatal("you must give in your api-key - follow the following to register for one.\n You receive" +
			"a maximum of 20,000 calls per month. Thats a good amount of calls")
	}
	port := flag.Int("port", 9191, "sets the port the underlying server shall run from")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}


func initialize(path string) {
	//SetupFolder
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating base directory | %s", err.Error())
	}

	err = os.MkdirAll(fmt.Sprintf("%s/pexels", path), os.ModePerm)
	if err != nil {
		log.Fatalf("error creating base directory for pexels folder | %s", err.Error())
	}
}

func GetPexelHandler(w http.ResponseWriter, r *http.Request) {
	pexel := pexels.PexelPhoto{}
	bytes, err := pexel.Get(string(2027695))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print("written: ", len(bytes))
}

