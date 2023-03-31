package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang, no super and parent

Nikhil := User{"Nikhil Chauhan", "thisisnikhilchauhan@gmail.com",true,21}
fmt.Println(Nikhil)
fmt.Printf("my details are :%+v\n",Nikhil) //+v gives output in structure format along with naming convention
fmt.Printf("my email is: %v and age is :%+v\n",Nikhil.email, Nikhil.Age)

}

type User struct{
	Name string
	email string
	status bool
	Age int
}

