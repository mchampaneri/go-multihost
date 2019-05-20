package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/fatih/color"
)

// ServerConfigData is server wide configuration
// carrier
var ServerConfigData ServerConfig
var dispatchMux http.ServeMux

func main() {

	// Loading configurations from json
	// file
	configError := configLoader()
	if configError != nil {
		color.Red("Failed to load configurations ")
		return
	}

	// Preparing dispatchMux For ReverseProxy
	dispatchMux := http.NewServeMux()
	for _, site := range ServerConfigData.Sites {
		parsedURL, _ := url.Parse(site.Forward)
		fmt.Println(site.Incomming)
		dispatchMux.Handle(site.Incomming, httputil.NewSingleHostReverseProxy(parsedURL))
		// redirector(dispatchMux, site.Incomming, site.Forward)
	}

	// Start the server
	color.Green("Server Spinned Up")
	http.ListenAndServe(":80", dispatchMux)
}

func configLoader() error {
	configFileHandle, fileError := os.Open("./server.json")
	if fileError != nil {
		color.Red(" Failed to load configuration file  ", fileError.Error())
		return fileError
	}

	readBytes, readError := ioutil.ReadAll(configFileHandle)
	if readError != nil {
		color.Red("Failed to read the file ", readError.Error())
		return readError
	}

	if unMarshalError := json.Unmarshal(readBytes, &ServerConfigData); unMarshalError != nil {
		color.Red("failed to unmarshal data ", unMarshalError.Error())
		return unMarshalError
	}

	color.Green("Configuration loaded successfully from file")
	return nil
}
