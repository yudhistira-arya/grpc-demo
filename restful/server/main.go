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

var meteorites []Meteorite

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(currentDir, "data", meteoriteJsonFile)
	log.Println(path)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(bytes, &meteorites); err != nil {
		log.Fatal(err)
	}
	log.Printf("JSON HTTP server initialization complete")
}

func main() {
	http.HandleFunc("/meteorites", getMeteors)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

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

func getMeteors(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	responseByte, err := json.Marshal(&meteorites)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = writer.Write(responseByte)
}
