/*
Grab a car from a json payload and post as a sell (POST /cars) to the Blockchain
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
	postURL := "http://129.213.52.65:4000/cars"

	client := &http.Client{}
	req, _ := http.NewRequest("POST", postURL, b)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Owner", "Chad Arimura")

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
