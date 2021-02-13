package main

import "fmt"

func main() {
	var marks = 29
	switch {

		case marks>=80 :
			fmt.Println("A+")
		case marks>=70 :
			fmt.Println("A")
		case marks>=65 :
			fmt.Println("A-")
		case marks>=33 :
			fmt.Println("Passed")
		default :
			fmt.Println("Failed")
	}
}