package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday" ,"friday", "saturday"}
	fmt.Println(days)

	// for d :=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// } 

	// we prefer this for loop, looks more clear and consise

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days{ //when we use range it returns two values "index of current element and value of current element"
		fmt.Printf("index is %v, and day is %v\n",index,day)
	}
}