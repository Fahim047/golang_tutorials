// all about variables declaration
package main

import "fmt"

var g = "this is global variable"	//global declaration

//z := "error"		//this will not work; shorthand declaration must be in a function

func main() {

	var name string  		//declarae one variable
	name = "Fahimul Islam Fahad"	//initialization
	fmt.Println(name)

	var a, b, c int = 1, 2, 3 		//many at once (same type of data)
	fmt.Println(a, b, c)

	var (	//many at once (different types)
		d int = 6
		e float32 = 3.75
		f bool = true
	)
	fmt.Println(d, e, f)

	var age, height, address = 21, 5.7, "Dhaka, Bangladesh" // declare and initialize many at once(infer-types)
	fmt.Println(name, age, height, address)

	x := "Fahim"		//shorthand declaration
	y := 2 				//shorthand declaration
	pi := 3.1416		//shorthand declaration
	fmt.Println(x, y, pi)

	fmt.Println(g)
	//fmt.Println(z)    //shorthand can't be declarae without function
}