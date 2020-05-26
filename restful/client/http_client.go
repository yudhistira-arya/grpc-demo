package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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

func GetMeteorite() []Meteorite {
	resp, err := http.Get("http://localhost:8090/meteorites")
	handleError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	var meteorites []Meteorite;
	handleError(json.Unmarshal(body, &meteorites))
	return meteorites
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("%v\n", GetMeteorite());
}
