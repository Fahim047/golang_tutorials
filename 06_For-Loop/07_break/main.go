package main

import "fmt"

func main(){
	for i:=1; i<10; i++ {
		fmt.Println(i)
		if i==5 {
			break //when i=5 loop will be break . we will see 1 to 5 in display.
		}
	}
}