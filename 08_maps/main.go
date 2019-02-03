package main

import "fmt"

func main() {
	// //* define map
	// emails := make(map[string]string)

	// // assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	//* deklare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete from map
	delete(emails, "Bob")

	fmt.Println(emails)
}
