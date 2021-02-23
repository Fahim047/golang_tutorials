package main

import "fmt"

func main(){
	var a [10]int
	for i:=0;i<10;i++{
		a[i]=i
	}
	//without skipping variable
	for index,value := range a {
		fmt.Printf("a[%d]= %d \n", index, value)
	}

	fmt.Prinln()

	var daysOfWeek = [7]string{"Saturday","Sunday","Monday","Tuesday","Wednesday","Thursday","Friday"}

	//skipping one variable using underscore "_"
	for _, value := range daysOfWeek {
		fmt.Println(value)
	}
}