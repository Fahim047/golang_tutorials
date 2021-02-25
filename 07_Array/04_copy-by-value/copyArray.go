package main

import "fmt"

func main() {
	//creating array using shorthand declaration
	my_arr1 := [5]string{"C","C++","Python","Java","Perl"}

	//copying the array to new array
	//here, the elements are passed by value
	my_arr2 := my_arr1

	fmt.Println("Array-1: ",my_arr1)
	fmt.Println("Array-2: ",my_arr2)

	my_arr1[0] = "GO"
	// Here, when we copy an array 
    // into another array by value 
    // then the changes made in original 
    // array do not reflect in the 
    // copy of that array
	fmt.Println("Array-1: ",my_arr1)
	fmt.Println("Array-2: ",my_arr2)
}
