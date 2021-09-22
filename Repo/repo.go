package Repo

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	"io/ioutil"
	"ReadFileProj/Validation"
	"time"
)


type IceCreamList struct {
	SelectedIceCream []ice_cream.IceCream
}


func CreateIceCream(ice ice_cream.IceCream) error{

	database, _ := GetAllIceCreams()

	valid, _ := Validation.Validate(ice)
	if (valid) {		
		id := time.Now().Unix()
		p := &ice
		p.Id = id
		database.SelectedIceCream = append(database.SelectedIceCream, ice)
		saveAndFlush(database)
	}
		return nil
}

func saveAndFlush (data IceCreamList) error {
	jsonFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("iceCreamDb.json", jsonFile, 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

func  FindIceCreamByID(id int64) ice_cream.IceCream {
	database, _ := GetAllIceCreams()

	for _, ice := range database.SelectedIceCream {
		if(id == ice.Id) {
			return ice
		} 
	}
	return ice_cream.IceCream{}
}



func UpdateIceCreamName(name string, id int64) ice_cream.IceCream  {
	database, _ := GetAllIceCreams()
	for i, ice := range database.SelectedIceCream {
		if(id == ice.Id) {
			itemsBefore := database.SelectedIceCream[:i]
			itemsAfter := database.SelectedIceCream[i +1:]
			updatedIceCream := database.SelectedIceCream[i]
			updatedIceCream.Name = name
			addingBack := append(itemsBefore, updatedIceCream)
			addingBack = append(addingBack, itemsAfter...)
			database.SelectedIceCream = addingBack
			saveAndFlush(database)
			return updatedIceCream
		}
	}
	 return 	ice_cream.IceCream{} 
}

func DeleteIceCreamByID(id int64) bool{
	database, _ := GetAllIceCreams()
	for i, ice := range database.SelectedIceCream {
		if id == ice.Id {
			itemsBefore := database.SelectedIceCream[:i]
			itemsAfter := database.SelectedIceCream[i +1:]
			addingBack := append(itemsBefore, itemsAfter...)
			database.SelectedIceCream = addingBack
			saveAndFlush(database)
			return true
		}
	}
	return 	false 

}



func  GetAllIceCreams() (IceCreamList, error){

		file, _ :=ioutil.ReadFile("iceCreamDb.json")
		data := IceCreamList{}
		err := json.Unmarshal([]byte(file), &data)
		if (err != nil) {
			panic(err)
		}
	return data,  err
		
}

