package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to Maths in Golang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("The sum of two number is: ", float64(myNumberOne)+(myNumberTwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// ramdom from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
