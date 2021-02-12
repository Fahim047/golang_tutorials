// constant decration and initialization
package main

import "fmt"

// global declaration
const pi = "3.1416" 	// single const declaration

const( 					// multiple const declaration
	x = "Allah Is Almighty"
	y = "Man Is Mortal"
)

func main(){

	const z = 63 	// local declaration
	fmt.Println(pi)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
