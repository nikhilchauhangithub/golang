package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("About slices")

	var fruitsList = []string{"apple","peach","banana"}
	fmt.Printf("Type of fruitsList is %T\n", fruitsList) //always use printf while using %T
	fmt.Println("Type of fruitsList is", len(fruitsList))

	fruitsList = append(fruitsList, "mango", "Litchi") //add more fruits to the fruitlist
	fmt.Println("Type of fruitsList is", fruitsList)

	fruitsList = append(fruitsList[2:4]) //remove 2 fruits from starting and 4 means it will stop counting at element n-1 which means 4-1=3, at 3rd element
	fmt.Println(fruitsList)

highScores := make([]int, 4)
highScores[0] =123
highScores[1] =5152
highScores[2] =12515
highScores[3] =12545

highScores = append(highScores, 0,11,22,33) //append reallocate all memories again

fmt.Println(highScores)


sort.Ints(highScores)

fmt.Println(highScores)
fmt.Println(sort.IntsAreSorted(highScores))

// Remove values from slices based on index

var courses=[]string{"java","html","js","php"}
var indexToRemove int =2
// courses = append(courses[:indexToRemove], courses[indexToRemove+1]) //easier syntax
courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...) //same work as above statement but with the ... notation, it spreads the elements of that slice as individual arguments to the append function.This also creates a new slice with the remaining elements after removing the element at the indexToRemove position, but it uses a slightly different syntax to achieve the same result.
 
fmt.Println(courses)

}