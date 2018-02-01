package main

import (
	"encoding/json"
	"os"
	"time"
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
	time.Sleep(5 * time.Second)

	c := new(car)
	err := json.NewDecoder(os.Stdin).Decode(c)
	if err != nil {
		panic(err)
	}
	str := "Pulling CarFax Report for " + c.Manufacturer + " " + c.Model + " - Plate: " + c.Plate
	os.Stderr.WriteString(str)
}
