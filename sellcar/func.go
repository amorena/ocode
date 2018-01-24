/*
Grab a car from a json payload and post to requestbin

TODO: Post the car to the Blockchain once I have an endpoint
*/
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type car struct {
	Plate        string `json:"plate"`
	Color        string `json:"color"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Year         string `json:"year"`
	Description  string `json:"description"`
	Condition    string `json:"condition"`
}

func main() {
	c := new(car)
	json.NewDecoder(os.Stdin).Decode(c)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	postURL := "https://requestb.in/1icjzd11"
	log.Printf("Sending %s to %s", string(b.Bytes()), postURL)
	res, _ := http.Post(postURL, "application/json", b)

	defer res.Body.Close()
}
