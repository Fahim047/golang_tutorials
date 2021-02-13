package main
import "fmt"

func main() {
	switch color:="white"; color {
		case "red", "blue", "yellow":
			fmt.Println("Colorful")
		case "black", "white":
			fmt.Println("Black & white")
		default:
			fmt.Println("didn't match")		
	}
}