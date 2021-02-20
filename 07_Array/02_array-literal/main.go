package main

import "fmt"

func main(){
	//predefined values
	arr := [5]int {1,2,3,4,5}
	fmt.Println("Elements of array:")
	for i:=0; i<5; i++ {
		fmt.Printf("%d ",arr[i])
	}
	fmt.Println()
	//we can use "..." as size . it will automatically measure the size of array
	b := [...]string {"Saturday", "Sunday", "Monday","Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println("Length of array b is: ",len(b))
	fmt.Println(b)
}