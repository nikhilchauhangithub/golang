package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("do rate our service")

	//comma ok || err err

	userInput, _ :=reader.ReadString('\n')   //underscore sign represents we dont care about error
	fmt.Println("thanks for rating us ", userInput)
	fmt.Printf("type of this rating is %T", userInput)
}