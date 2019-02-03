package main

import "fmt"

func main() {
	// //* if else
	// x := 11
	// y := 12
	// if x <= y {
	// 	//* %d is a placeholder for dezimal numbers less than 10, i guess
	// 	fmt.Printf("%d is less than or equal to  %d\n", x, y)
	// } else {
	// 	fmt.Printf("%d is less than %d\n", y, x)
	// }

	// //* else if
	// color := "green"

	// if color == "red" {
	// 	fmt.Println("color is red")
	// } else if color == "blue" {
	// 	fmt.Println("color is blue")
	// } else {
	// 	fmt.Println("color is not blue or red")
	// }

	//* switch
	var color string = "white"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "black":
		fmt.Println("color is black")
	default:
		fmt.Println("color is not black or red")
	}
}
