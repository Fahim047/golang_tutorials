package main

import "fmt"

func main() {

	name := "Fahimul"
	switch {
		case name=="FIF" :
			fmt.Println("FIF")
		case name=="Fahimul" :
			fmt.Println("Fahimul")
			fallthrough
		case name=="Islam" :
			fmt.Println("Islam") //this will displayed even condition is false,using fallthrough
			fallthrough
		case name=="Fahad" :
			fmt.Println("Fahad") //this will displayed even condition is false, using fallthrough
		case name=="FIF" :
			fmt.Println("FIF") 		//this will not displayed , fallthrough not used
	}
}