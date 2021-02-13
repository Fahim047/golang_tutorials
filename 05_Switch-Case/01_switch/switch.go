package main
import "fmt"

func main() {
	var color = "red"
	switch color {
		case "black":
			fmt.Println("Monday")
		case "blue": 
			fmt.Println("Tuesday")
		case "pink": 
			fmt.Println("Wednesday")
		case "orange": 
			fmt.Println("Thursday")
		case "yellow": 
			fmt.Println("Friday")
		case "red": {
			fmt.Println("RED")
			fmt.Println("opssss!!! color of blood")
		}
		default: fmt.Println("color didn't match")
	}
}