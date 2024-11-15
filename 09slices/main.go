package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Golang")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Println("FruitList is ", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println("New FruitList is ", fruitList)

	// fruitList = append(fruitList[1:3]) // starts with index 1 to index 3....ps: won't include index 3
	// fmt.Println("New FruitList is ", fruitList)

	highscore := make([]int, 4)

	highscore[0] = 111
	highscore[1] = 666
	highscore[2] = 444
	highscore[3] = 333

	fmt.Println(highscore)

	highscore = append(highscore, 333, 555)
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "rust"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
