package main
import "fmt"

func main() {
	var color = "red"
	switch color {
		case "black":
			fmt.Println("Black")
		case "blue": 
			fmt.Println("Blue")
		case "pink": 
			fmt.Println("Pink")
		case "orange": 
			fmt.Println("Orange")
		case "yellow": 
			fmt.Println("Yellow")
		case "red": {
			fmt.Println("RED")
			fmt.Println("opssss!!! color of blood")
		}
		default: fmt.Println("color didn't match")
	}
}