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
    diceNumber := r.Intn(6)
// diceNumber := rand.Intn(6)+1 //+1 eliminate zero from coming 
fmt.Println("the dice num is:",diceNumber)

}