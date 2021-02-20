package main

import "fmt"

func main(){
	//creating array using var keyword
	var arr [3]string

	//elements are assigned using index
	arr[0]= "Fahimul"
	arr[1]= "Islam"
	arr[2]= "Fahad"

	//accessing the elements of array
	fmt.Println(arr)
	fmt.Println(arr[2])  //accessing specific element
}