package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List's length is: ", len(fruitList))

	var vegList = []string{"Potato", "Tomato", "Mushroom"}
	fmt.Println("Vegetable List is: ", vegList)
	fmt.Println("Vegetable List's lenght is: ", len(vegList))
}
