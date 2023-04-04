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
	DecodeJson()
}

func EncodeJson()  {
	myCourse := []course{
		
		{"Mahadev", 22, []string{"Shiv Shambhu","Bholenath", "Neelkanth"}},
		{"Nikhil Chauhan", 21, []string{"good guy","developer", "programmer"}},
	}

	//package this data as JSON data

finalJson,err:=	json.MarshalIndent(myCourse,"","\t")
if err!=nil {
	panic(err)
}
fmt.Printf("%s\n",finalJson)
}

func DecodeJson()  {
	jsonFromWeb := []byte(`
	{
		"Name": "Mahadev",
		"Age": 22,
		"Tags": [
			"Shiv Shambhu",
			"Bholenath",
			"Neelkanth"
		]
	}
	`)

	var myCourse course
	checkValid := json.Valid(jsonFromWeb)

	if checkValid {
		fmt.Println("Json data is valid")
		json.Unmarshal(jsonFromWeb, &myCourse)
		fmt.Printf("%#v\n",myCourse)
	} else{
		fmt.Println("some problem is json data")
	}
//some cases where you just want to add data to key value

var myOnlineData map[string]interface{}
json.Unmarshal(jsonFromWeb, &myOnlineData)
fmt.Printf("%#v\n",myOnlineData)

for k,v :=range myOnlineData{
	fmt.Printf("key is %v and value is %v and Type is :%T\n",k,v,v)
}
}

