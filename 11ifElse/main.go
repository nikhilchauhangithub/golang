package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("if else in golang")
	fmt.Println("Enter your age")
var result string
reader :=bufio.NewReader(os.Stdin) 

userAge,_ :=reader.ReadString('\n')
ageInNum,_ := strconv.ParseFloat(strings.TrimSpace(userAge),64)

if ageInNum < 18 {
	result ="can't apply for DL"
} else if ageInNum > 18 {
	result ="you can apply for DL"
} else{
	result ="You can only apply for learning license"
}

fmt.Println(result)
}