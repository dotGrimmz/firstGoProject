package Repo

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"io/ioutil"
	//   "fmt"
)


type IceCreamList struct {
	SelectedIceCream []ice_cream.IceCream
}

// Create a function that accepts an IceCream, converts it to JSON and stores it on the next available line in iceCreamDb.txt
// Return an error if any of the functions you call return an error


func (list IceCreamList) CreateIceCream(ice ice_cream.IceCream) []ice_cream.IceCream {
list.SelectedIceCream = append(list.SelectedIceCream, ice)
return list.SelectedIceCream
}

func (list IceCreamList) FindIceCreamByName(name string) ice_cream.IceCream {
	// fmt.Printf("%v", list.SelectedIceCream)
	// best way to handle this situation
	for _, ice := range list.SelectedIceCream {
		if(name == ice.Name) {
			return ice
		} 
	}
	// type check? 
	// im returning just the values..right? 

	return ice_cream.IceCream{}
}



func (list IceCreamList) UpdateIceCreamName(name string, ice ice_cream.IceCream) ice_cream.IceCream  {
	p := &ice
	 p.Name = name

	 return ice
}




// Create function that can retrieve all the ice creams from iceCreamDb.txt
// The function should unmarshal each one into an ice cream struct and return a slice of IceCream --> []IceCream
// Return an error if any of the functions you call return an error
// func (list *IceCreamList) GetAllIceCreams() []ice_cream.IceCream{

func (list IceCreamList) GetAllIceCreams() ([]ice_cream.IceCream, error){

		file, _ :=ioutil.ReadFile("iceCreamDb.json")
		data := IceCreamList{}
		err := json.Unmarshal([]byte(file), &data)
		if (err != nil) {
			panic(err)
		}
	return data.SelectedIceCream,  err
		
}

