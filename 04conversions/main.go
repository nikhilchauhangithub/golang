package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// here we learn about string conversion and trimming
func main() {
	fmt.Println("wlecome to our agency")
	fmt.Println("do rate us")

	
	reader := bufio.NewReader(os.Stdin)

	input, _ :=reader.ReadString('\n')
	fmt.Println("thanku for rating us: ",input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err!=nil{
		fmt.Println(err)
	}
	 {
		fmt.Println("add 1 to your rating:", numRating+1)
	}
	
}