/*
Grab a car from a json payload and post as an owner transfer (POST /owner) on Blockchain
The path is /cars/<carId>/owner

header: Owner
body: new_owner
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type car struct {
	Plate        string `json:"plate"`
	CurrentOwner string `json:"current_owner"`
	NewOwner     string `json:"new_owner"`
}

func main() {
	c := new(car)
	json.NewDecoder(os.Stdin).Decode(c)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	postURL := "http://129.213.52.65:4000/cars/" + c.Plate + "/owner"

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", postURL, b)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Owner", c.CurrentOwner)

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println("error with dumprequest")
		log.Fatal(err)
	}
	log.Println(string(dump))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response --> " + res.Status)
	defer res.Body.Close()
}
