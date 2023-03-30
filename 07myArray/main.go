package main

import "fmt"

func main() {
	fmt.Println("array in golnag")

	var fruitsList [4]string

	fruitsList [0] ="Apple"
	fruitsList [1] ="Guavaa"
	fruitsList [2] ="Peach"

	fmt.Println("Here is the fruit: ", fruitsList[1])
	fmt.Println("Length of array: ", len(fruitsList[1])) //this gives the length of word guava

	//another way of writing an array

	var vegetable = [3] string {"potato", "tomato", "beans"}
	fmt.Println("Here is the veggies: ", vegetable)

}