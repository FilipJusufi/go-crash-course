package main

import "fmt"

var nobody = true

func main() {
	// * MAIN TYPES (some of them)
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float64
	// complex64 complex128

	// * using var
	var name string = "Filip"
	// or - it knows that is a string, so string-type is unnecessary
	var name2 = "Filip"
	// if you declare a variable and not use it - you get an error
	var age int32 = 27
	// or - it knows that is a int, so int-type is unnecessary
	var age2 = 27
	var size2 float32 = 2.3

	// * using const
	const isTrue = true

	// using shorthand declaration - only inside functions possible
	name3 := "ansem"
	size := 1.3

	// using 2 declarations in one line
	user, userEmail := "Filip", "filip.jusufi@gmail.com"

	fmt.Println(name, name2, age, age2, isTrue, name3, size, size2)
	// check type name
	fmt.Printf("%T\n", name)
	// check type age
	fmt.Printf("%T\n", age)
	// check type isTrue
	fmt.Printf("%T\n", isTrue)
	// check type size
	fmt.Printf("%T\n", size)
	// check type size
	fmt.Printf("%T\n", size2)

	fmt.Println(user, userEmail)
	fmt.Printf("%T\n", user)
	fmt.Printf("%T\n", userEmail)


	fmt.Printf("%T\n", nobody)
}
