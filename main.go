package main

// Import the necessary packages and call the appropriate methods to read a files input and convert it to an IceCream
// You should then validate IceCream if it fails validation print that out.
// If the icecream passes validation, store the IceCream in the db
import (
	 "fmt"
	// "encoding/json"
	// "io/ioutil"
	 "ReadFileProj/FileReader"
	"ReadFileProj/Repo"
)


func main() {

	// Initialize Struct that will call to the json database
	
















	// var test int = 1632256929

	// testInt64 := int64(test)


	data, _ := FileReader.ReadFile("input/ic2.json")
fmt.Println(Repo.CreateIceCream(data))

		// p := IceCreamRepo.SelectedIceCream[0]
		// fmt.Printf("%T", p)
		// fmt.Println(IceCreamRepo.UpdateIceCreamName("Rakke",IceCreamRepo.SelectedIceCream[0]))
		// fmt.Println(IceCreamRepo.SelectedIceCream[0])
		// fmt.Println("One or more of these toppings are invalid")


	// c := Gen(2, 3)
	// out := Sq(c)

	// fmt.Println(<-out)
	// fmt.Println(<-out)

}






// func  Gen(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			fmt.Println("this is n",n)
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func  Sq(in <- chan int) <- chan int {
// 	out:= make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n*n
// 		}
// 		close(out)
// 	}()
// 	return out
// }