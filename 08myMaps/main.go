package main

import (
	"fmt"
)

func main() {

	fmt.Println("maps in golang")

	languages := make(map[string]string)
	languages["JS"]="javascript"
	languages["py"]="python"
	languages["rb"]="ruby"

	fmt.Println("languages in map", languages)
	fmt.Println("languages in map", languages["JS"])

}