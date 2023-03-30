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

}