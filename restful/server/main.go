package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const meteoriteJsonFile = "y77d-th95.json"

type Meteorite struct {
	Name        string      `json:"name"`
	Id          string      `json:"id"`
	NameType    string      `json:"nametype"`
	RecClass    string      `json:"recclass"`
	Mass        string      `json:"mass"`
	Fall        string      `json:"fall"`
	Year        string      `json:"year"`
	Reclat      string      `json:"reclat"`
	Reclong     string      `json:"reclong"`
	Geolocation Geolocation `json:"geolocation"`
}

type Geolocation struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

var meteorites []Meteorite

func init() {
	currentDir, err := os.Getwd()
	handleError(err)

	path := filepath.Join(currentDir, "data", meteoriteJsonFile)
	log.Println(path)

	bytes, err := ioutil.ReadFile(path)
	handleError(err)

	handleError(json.Unmarshal(bytes, &meteorites))
	log.Printf("JSON HTTP server initialization complete")
}

func main() {
	http.HandleFunc("/meteorites", getMeteors)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func getMeteors(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	responseByte, err := json.Marshal(&meteorites)
	handleError(err)
	_, _ = writer.Write(responseByte)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
