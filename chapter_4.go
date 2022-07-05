package main

import (
	"fmt"
)


func main(){
	var i = 10
	if i < 10{
		fmt.Println("hey its less than ten")
	}else {
		fmt.Println("its greater than ten or equal to 10")
	}

	for i :=1; i < 101; i++{
		if i % 3 == 0{
			fmt.Println(i)
		}
	}

	fmt.Println("\n\n",i)
}