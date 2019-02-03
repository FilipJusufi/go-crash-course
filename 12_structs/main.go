package main

import (
	"fmt"
	"strconv"
)

//* define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//* shorthand
	firstName, lastName, city, gender string
	age                               int
}

//* Greeting method (value reciever)
// p is similar to this
// p.age has to be converted in string
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

//* hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

//* getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//* init Peron using struct
	person1 := Person{firstName: "Filip", lastName: "Jusufi", city: "Vienna", gender: "male", age: 27}
	// Alternative
	person2 := Person{"Maria", "Jusufi", "Vienna", "female", 20}
	fmt.Println(person1)
	fmt.Println(person2)

	person1.age++
	person1.hasBirthday()
	person2.getMarried("William")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
