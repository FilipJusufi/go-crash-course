package main

// import more packages with parantheses
// ! package will dissapear on save if not used

import (
	"fmt"
	"math"

	// using my own package
	"github.com/filipjusufi/go_crash_course/03_packages/strutil"
)

var greeting string = "Hello World"

func main() {
	fmt.Println(greeting)
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
