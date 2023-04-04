package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string
	Age int
	Tags []string
}

func main() {
	fmt.Println("More on json in golang")
	EncodeJson()
}

func EncodeJson()  {
	myCourse := []course{
		
		{"Mahadev", 22, []string{"Shiv Shambhu","Bholenath", "Neelkanth"}},
		{"Nikhil Chauhan", 21, []string{"good guy","developer", "programmer"}},
	}

	//package this data as JSON data

finalJson,err:=	json.MarshalIndent(myCourse,"Nikhil","\t")
if err!=nil {
	panic(err)
}
fmt.Printf("%s\n",finalJson)
}