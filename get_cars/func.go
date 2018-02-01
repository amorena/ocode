/*
This is a stub that mocks the return of a couple cars from a fake getcars API for testing
the getcars func.
*/
package main

import (
	"bytes"
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
type cars struct {
	Cars []car `json:"cars"`
}

func main() {
	time.Sleep(3 * time.Second)
	c1 := newCar("Tesla", "Model 3", "2NOGA55", "New fancy sedan", "black")
	c2 := newCar("Tesla", "Model X", "4PWR123", "Big X-Wing SUV", "red")
	cars := &cars{
		Cars: []car{*c1, *c2},
	}
	b, _ := json.Marshal(cars)

	var buffer bytes.Buffer

	buffer.WriteString("Found " + string(len(cars.Cars)) + "Cars:\n")

	for _, c := range cars.Cars {
		buffer.WriteString(c.Manufacturer + " " + c.Model + " - Plate: " + c.Plate + "\n")
	}

	os.Stderr.WriteString(buffer.String())
	os.Stdout.Write(b)
}

func newCar(manu string, model string, plate string, desc string, color string) *car {
	c := new(car)
	c.Manufacturer = manu
	c.Model = model
	c.Plate = plate
	c.Description = desc
	c.Color = color
	return c
}
