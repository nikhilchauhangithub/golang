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

//delete with id as "rb" is id in below case

// delete(languages,"rb")
// fmt.Println("languages in map", languages)


//loops in golang

for key,value := range languages{
	fmt.Printf("for key %v, value is %v\n", key, value)
}

}