package main

import "fmt"

func main(){
	//creating array using shorthand declaration
	my_arr1 := [7]int{1,5,7,10,13,47,58}

	//copying the array to new array
	//here, the elements are passed by reference
	my_arr2 := &my_arr1

	fmt.Println("Before change:")
	fmt.Println("Array-1: ",my_arr1)
	fmt.Println("Array-2: ",*my_arr2)

	my_arr1[2] = 99
	fmt.Println("After change:")
	// Here, when we copy an array 
    // into another array by value 
    // then the changes made in original 
    // array will reflect in the 
    // copy of that array
	fmt.Println("Array-1: ",my_arr1)
	fmt.Println("Array-2: ",*my_arr2)
}