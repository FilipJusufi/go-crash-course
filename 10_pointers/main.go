package main

import "fmt"

func main() {
	var a int = 5
	// using pointer
	b := &a
	fmt.Println(a, b)
	// two different datatypes
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//* use * to read value from adress
	fmt.Println(*b)

	// change vlue with pointer
	*b = 10
	fmt.Println(a)
}
