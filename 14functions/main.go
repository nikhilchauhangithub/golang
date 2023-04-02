package main

import "fmt"

//we are not allowed to create func inside a func

func main() {
	hindi() //if we call hindi func 1st it will print content of hindi function 1st
	fmt.Println("Hello world")

	result := add("1","7")
	fmt.Println(result)

	proResult,msg := proAdder(1,1,1)
	fmt.Println(proResult)
	fmt.Println(msg)
}

func hindi()  {
	fmt.Println("Namaste world")
}

func add(val1 string, val2 string) string  {
	return val1+ val2
}

//important syntax
func proAdder(values ...int) (int, string) { //didnt care about no. of arguments
	total :=0

for _, val :=range values{
	total += val
}

	return total, "How you doin ?"
}

