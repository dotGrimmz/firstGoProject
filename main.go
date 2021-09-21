package main

// Import the necessary packages and call the appropriate methods to read a files input and convert it to an IceCream
// You should then validate IceCream if it fails validation print that out.
// If the icecream passes validation, store the IceCream in the db
import (
	 "fmt"
	"encoding/json"
	"io/ioutil"
	"ReadFileProj/FileReader"
	"ReadFileProj/Validation"
	"ReadFileProj/Repo"
	"os"
)


func main() {
	// Read in the input and unmarshal to an IceCream struct


	IceCreamRepo := Repo.IceCreamList{};
	

	// slice of the database
	store, _ := IceCreamRepo.GetAllIceCreams();
	IceCreamRepo.SelectedIceCream = store
	// fmt.Println("selected icecream list", IceCreamRepo.SelectedIceCream)
	// Compare the toppings to the toppings in toppings.json
	// utils.ReadFile("/input/ic1.json")
	// If validation fails returns false from the Validate function and return from the main function

	// If validation passes, marshall the icecream into JSON and store it in iceCreamDb.txt


	data, _ := FileReader.ReadFile("input/ic1.json")
	valid, _ := Validation.Validate(data)
	if(valid) {
		IceCreamRepo.SelectedIceCream = IceCreamRepo.CreateIceCream(data)
		iceCreamDatabase, err := os.OpenFile("iceCreamDb.txt", os.O_APPEND| os.O_CREATE|os.O_WRONLY, 0644)
		defer iceCreamDatabase.Close();
		if err != nil {
			panic(err)
		}
		for _,  iceCreamStruct := range IceCreamRepo.SelectedIceCream {
			ice, _ := json.Marshal(iceCreamStruct)			
		if _, err := iceCreamDatabase.WriteString(string(ice) +"\n"); err != nil {
			panic(err)
		}
	}

	jsonFile, err := json.MarshalIndent(IceCreamRepo, "", " ")
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("iceCreamDb.json", jsonFile, 0644)

	fmt.Println("Successfully uploaded new icecream")
	
	} else{
		p := IceCreamRepo.SelectedIceCream[0]
		fmt.Printf("%T", p)
		fmt.Println(IceCreamRepo.UpdateIceCreamName("Rakke",IceCreamRepo.SelectedIceCream[0]))
		// fmt.Println(IceCreamRepo.SelectedIceCream[0])
		// fmt.Println("One or more of these toppings are invalid")
	}

	c := Gen(2, 3)
	out := Sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)

}






func  Gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Println("this is n",n)
			out <- n
		}
		close(out)
	}()
	return out
}

func  Sq(in <- chan int) <- chan int {
	out:= make(chan int)
	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()
	return out
}