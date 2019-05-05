package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/martinomburajr/pexels/app"
	"github.com/martinomburajr/pexels/auth"
	"github.com/martinomburajr/pexels/config"
	"github.com/martinomburajr/pexels/pexels"
	"github.com/martinomburajr/pexels/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var port int

func init() {
	//Retrieve APIKEY
	var ap string
	flag.StringVar(&ap, "key", "", "sets the api to be used by the individual")
	port = *(flag.Int("port", 9191, "sets the port the underlying server shall run from"))

	flag.Parse()

	if ap == "" {
		log.Print("Locating API KEY ...")

		dir := config.GetHomeDir()
		pexelsConfig := config.PexelsConfig{}

		err := pexelsConfig.Load(config.ConfigPath(dir))
		if err != nil {
			msg := "You need to supply your api-key using the -key flag.\n " +
				"If you DO NOT have a key, follow the following link to register for one.\n " +
				"https://www.pexels.com/api/new/\n " +
				"You receive a maximum of 20,000 calls per month. That's a good amount of calls ;-)"
			log.Fatalf("\n\n%s\n\n%s", msg, err.Error())
		}
		log.Print("API KEY Found! :D")
		return
	}

	auth.PexelSession.API_KEY = ap
	initialize()
}

func main() {
	p := pexels.PexelPhoto{}
	u := utils.Utils{}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: nil,
	}

	server := &app.PrimeServer{
		PexelsDB: &p,
		Utilizer: &u,
		HTTPDefaultClient: client,
		Session: auth.PexelSession,
	}

	log.Print(fmt.Sprintf("pexels server started on port %d", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server.Routes()))
}

//initialize performs the initialization steps to ensure there is a pexels folder and config file
func initialize() {
	//SetupFolder
	err := createPexelsFolder()
	if err != nil {
		log.Fatal(err)
	}

	err = createConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = createPexelsPictureFolder()
	if err != nil {
		log.Fatal(err)
	}
}

//createConfig creates a pexels.config.json file in the .pexels directory that holds information such as the API_KEY
func createConfig() error {
	pexelsConfig := config.PexelsConfig{
		APIKEY: auth.PexelSession.API_KEY,
	}

	data, err := json.Marshal(pexelsConfig)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(config.ConfigPath(""), data, 0755)
}

//createPexelsFolder creates the canonical base path for this application. Without it, the application will try and recreate it.
func createPexelsFolder() error {
	err := os.MkdirAll(config.CanonicalBasePath(""), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating base directory | %s", err.Error())
	}
	return nil
}

//createPexelsPictureFolder creates a picture folder within the canonical base path
func createPexelsPictureFolder() error {
	err := os.MkdirAll(config.CanonicalBasePath("")+"/pictures", os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating pictures folder in config.CanonicalBasePath %s | %s", config.CanonicalBasePath(""), err.Error())
	}
	return nil
}
