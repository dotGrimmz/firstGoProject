package Validation

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"io/ioutil"
)

// Write a function that takes in an ice cream then reads the file toppings.json.
// Iterate through the toppings returned from toppings.json and iterate throught the toppings in the IceCream
// to verify that the IceCream ony contains toppings listed in toppings.json.
// If the ice cream fails validation return false, otherwise return true.

type toppings struct {
	Toppings []string
}

// Validate to get away the nasty message
func Validate(data ice_cream.IceCream) ( bool, error) {
	toppingsFile, err := ioutil.ReadFile("input/toppings.json")
	if err != nil {
		panic(err)
	}
	ToppingsStruct := toppings{}
	var valid = true
	err = json.Unmarshal([]byte(toppingsFile), &ToppingsStruct)
	if err != nil {
		panic(err)
	}
	elementsMap := make(map[string]bool)

	for _, topping := range ToppingsStruct.Toppings {
		elementsMap[topping] = true
	}

	for _, value := range data.Toppings {
	if(!elementsMap[value]) {
		valid = elementsMap[value]
		}
	}

	return valid, err
}
