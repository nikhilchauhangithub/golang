package main

import (
	"fmt"
)

//if first letter is uppercase as "L" in this case, it means it is public and we can use this anywhere
const LoginToken string ="abc" 

func main() {
	var username string ="Nikhil"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool =false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 =255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.445545484   //we'll get more precise value if we use float64
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat) 


    //default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)


	//implicit types i.e we didnt mention it is string but it is
	var website ="go.dev"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)


	//no var style, as it is of type int
	// here we use := (walrus operator) which assign  value to a variable and returning that value
	//but it is not allowed to use walrus ope outside function, if you want you'll have to use var and type

	numberOfUser := 17
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}