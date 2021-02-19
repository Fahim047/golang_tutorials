package main

import "fmt"

func main(){
	for i:=1; i<10; i++ {
		if i==5 || i==3 {
			continue  		// when i=3,5 loop will skip all other code in loop body 
							//but loop will not break
		}
		fmt.Println(i)
		
	}
}