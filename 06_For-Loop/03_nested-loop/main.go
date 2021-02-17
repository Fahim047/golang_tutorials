package main

import "fmt"

func main(){

	i := 1
	for ; i<5; i++ {
		for j:=0; j<i; j++ { 	//nested for loop
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}