package main

import "fmt"

func main() {

	if n := 15; n%2 == 0 {			// short declaration statement
	fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}
	/*
	The variable declared in the short statement 
	 is only available inside the if block and
	 it’s else or else-if branches -
	*/
}