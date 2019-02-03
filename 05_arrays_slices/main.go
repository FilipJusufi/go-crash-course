package main

import "fmt"

func main() {
	//* Arrays
	// var fruitArr [2]string

	//* Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Banana"

	// Declase and assind - //? the same as above
	// fruitArr := [2]string{"Apple", "Banana"}

	fruitSlice := []string{"Apple", "Banana", "Ananas"}

	// fmt.Println(fruitArr)
	//* ask for array index
	fmt.Println(len(fruitSlice))
	//* get range
	fmt.Println(fruitSlice[1:2])
}
