package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in Golang")

	seed := time.Now().UnixNano()
    r := rand.New(rand.NewSource(seed))
    
    // Use the new random number generator to generate random numbers
    diceNumber := r.Intn(6)+1 //+1 eliminate zero from coming 
// diceNumber := rand.Intn(6)+1 //+1 eliminate zero from coming 
// fmt.Println("the dice num is:",diceNumber)

switch diceNumber {
case 1:
	fmt.Println("you can open/move but no bonus chance")

case 2:
	
	fmt.Println("move 2")
	fallthrough  //also print consecutive next output 
case 3:
	fmt.Println("move 3")
case 4:
	fmt.Println("move 4")
case 5:
	fmt.Println("move 5")
case 6:
	fmt.Println("you can open/move and has a bonus chance")
default:
	fmt.Println("what was that")
}

}