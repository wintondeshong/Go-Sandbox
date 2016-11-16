package main

import (
	"encoding/json"
	"github.com/wintondeshong/Go-Sandbox/webapp-stdlib/controllers"
	"log"
	"net/http"
	"os"
)

var (
	config Configuration
	logger *log.Logger
)

// Startup
// -------

func main() {
	// Configuration
	initLogging()
	initConfig()

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))

	// Routes
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//mux.Handle("/404", error404)
	mux.HandleFunc("/", controllers.Index)

	// Startup
	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}

// Private Methods
// ---------------

type Configuration struct {
	Address string
	Static  string
}

func initConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot load config file", err)
	}

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatalln("Cannot parse configuration from file", err)
	}
}

func initLogging() {
	file, err := os.OpenFile("log/development.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}
