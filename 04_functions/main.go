package main

import "fmt"

// func-declaration func-name (func-parameter-type) func-return-type
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

// if two params are int
func getPro(num1, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Println(greeting("Filip"))
	fmt.Println(getSum(3, 4))
	fmt.Println(getPro(2, 13))
}
