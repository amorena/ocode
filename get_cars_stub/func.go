/*
Mock the return of a couple cars from a fake getcars API for testing
the getcars func.
*/
package main

import (
	"encoding/json"
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
	c1 := newCar("Tesla", "Model 3", "abc-123", "New fancy sudan")
	c2 := newCar("Tesla", "Model X", "xyz-321", "Big X-Wing SUV")

	var cars []car
	cars = append(cars, *c1, *c2)

	b, _ := json.Marshal(cars)
	os.Stdout.Write(b)
}

func newCar(manu string, model string, plate string, desc string) *car {
	c := new(car)
	c.Manufacturer = manu
	c.Model = model
	c.Plate = plate
	c.Description = desc
	return c
}
