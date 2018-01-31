/*
Get an array of cars from a mock API and marshal into JSON.

*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type car struct {
	Plate        string      `json:"plate"`
	Color        string      `json:"color"`
	Model        string      `json:"model"`
	Manufacturer string      `json:"manufacturer"`
	Year         json.Number `json:"year, Number"`
	Description  string      `json:"description"`
	Condition    string      `json:"condition"`
}

// type cars struct {
// 	Cars []car `json:"cars"`
// }

func main() {
	getURL := "http://129.213.52.65:4000/cars"
	res, _ := http.Get(getURL)
	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	//cars := new(cars)

	var cars []car

	err := json.Unmarshal(bodyBytes, &cars)
	if err != nil {
		fmt.Println("error:", err)
	}

	b, _ := json.Marshal(cars)

	fmt.Println(string(b))
}
