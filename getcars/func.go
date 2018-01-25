/*
Get an array of cars from a mock API and marshal into JSON.

TODO: not sure if we really need this. could do something interesting.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

type cars struct {
	Cars []car `json:"cars"`
}

func main() {
	getURL := "http://129.213.52.65:4000/cars"
	fmt.Println(getURL)
	res, _ := http.Get(getURL)
	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(bodyBytes))

	cars := new(cars)

	json.Unmarshal(bodyBytes, &cars)

	for k, v := range cars.Cars {
		fmt.Printf("%v -> %v\n", k, v)
	}
}
