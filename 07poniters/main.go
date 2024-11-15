package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers in Golang")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 24

	ptr = &myNumber
	fmt.Println("Address of pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value of pointer is", *ptr)
}
