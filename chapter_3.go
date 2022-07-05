package main

import "fmt"

// x := 6 // exists in the global scope

func main(){
	dogName := "martins" //automatic casting of type dogName is auto string
	
	fmt.Println("the name of my dog is", dogName)

	var count int8 = 5

	count +=1 

	fmt.Println("my count is incremented by 1", count)

}


// func dummy(){
// 	fmt.Println(x)
// }