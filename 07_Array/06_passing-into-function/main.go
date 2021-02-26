package main

import "fmt"

//This function accept 
// an array as an argument
func passArray(a [5]int, size int) int{
	var sum, i int
	for i=0; i<size; i++ {
		sum += a[i]
	}
	return sum
}

//Main function
func  main(){
	// creating an array size of 5
	var arr = [5]int{3,5,7,9,11}
	var res int

	// passing an array as argument
	res = passArray(arr,5)
	fmt.Println("The sum of array is: ",res)
}