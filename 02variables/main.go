package main

import "fmt"

const LoginToken = "hafhakhfaljflafljak" //public

func main() {
	var username string = "khadgachy07"
	fmt.Println(username)
	fmt.Printf("Type of variable is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable is %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of variable is %T \n", smallVal)

	var smallFloat float64 = 255.56565665
	fmt.Println(smallFloat)
	fmt.Printf("Type of variable is %T \n", smallFloat)

	// default values and alias
	var defaultVal int
	fmt.Println(defaultVal)
	fmt.Printf("Type of variable is %T \n", defaultVal)

	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("Type of variable is %T \n", defaultString)

	//implict type
	var website = "localhost:3000"
	fmt.Println(website)
	fmt.Printf("Type of variable is %T \n", website)

	// no var style
	numberOfUser := 6666
	fmt.Println(numberOfUser)
	fmt.Printf("Type of variable is %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Type of variable is %T \n", LoginToken)

}
