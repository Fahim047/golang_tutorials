// if statement
package main

import "fmt"

func main() {
	// if statement
	if true {
		fmt.Println("You can see me!")
	}
	if false {
		fmt.Println("You can't see me!")
	}

	// exclametory makes the statement opposite
	if !true { 		// this became false so it will not be displayed
		fmt.Println("Yes, yes, yes")
	}
	if !false { 	// this became true so it will be displayed
		fmt.Println("No, no, no")
	}
}