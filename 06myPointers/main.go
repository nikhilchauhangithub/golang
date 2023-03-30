package main

import "fmt"

func main() {
	fmt.Println("Learning pointers in Go")

	// var ptr *int
    // fmt.Println("the value of this pointer is", ptr)
	//default value of pointer is nill


	//assisn value to pointer

	myNumber := 17
	var ptr =&myNumber  //we use & because here we are referencing value of myNumber to ptr
	fmt.Println("the value of this pointer is", ptr) //gives the actual memory address of pointer
	fmt.Println("the value of this pointer is", *ptr) //gives the value in pointer or u can say *ptr means value of ptr

	fmt.Println("what is the value of myNumber")

	*ptr = *ptr +2
//This line is adding 2 to the value stored at the memory location pointed to by ptr. Since ptr is pointing to the memory location where myNumber is stored, this line effectively adds 2 to the value of myNumber. So after this line is executed, the value of myNumber is no longer 17, it is now 19. This is because the value of myNumber is the value stored at the memory location pointed to by ptr, and that value has been modified by the previous line of code.


	fmt.Println("value of actual pointer is:", myNumber)



}