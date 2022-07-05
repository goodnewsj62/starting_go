package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var x *int
	y :=5 

	x = &y

	fmt.Println("the memory value of y is: ", x)
	fmt.Println("while the actual value of y", *x)
	*x = 7
	fmt.Println(":", *x, y)

	rand.Seed(time.Now().UnixNano())

	var some [5]float64

	some[0] =1 
	some[1] = 2
	some[2] =3 
	some[3] = 4
	some[4] = 5

	for i:=0; i < 5; i++{
		fmt.Print(some[i])
	}

	another_array := [2]float64{1,2}
	fmt.Println("\n\n","another array:",another_array,"\n\n")

	longer_array := [8]int{1,2,3,4,5,6,67,89}

	for i, value := range longer_array{
		fmt.Println("index:", i,"value: ", value)
	}

	var b [5]float64

	for i :=0; i < 5; i++{
		b[i] = float64(i) * 3.1
	}


	array_1 := [...]int{1,2,3,4,5}

	fmt.Println(array_1)

	partial_array := [...]string{1:"yo partial"}
	fmt.Print(partial_array[0] == "")

	slice_1 := []int{}
	al := "apple"
	slice_2 := []string{"alpha","beta",string(al)}

	fmt.Println("capacity slice_1:",cap(slice_1),"\ncapacity slice_2:",cap(slice_2))

	slice_a := []int{1,2,3,4,5,6}
	slice_b := []string{5:"red"}
	slice_c := array_1[:]

	fmt.Println("\nslice_a:",slice_a,"slice_b:",slice_b, "slice_c:", slice_c)

	make_a := make([]int,2,5)

	for i , value := range make_a{
		make_a[i] = i + 3 + value
	}

	fmt.Println(make_a)

	slice_c = append(slice_c, slice_a...)
	slice_b = append(slice_b, "hey","yo","stupid")

	fmt.Println("\nslice_a:",slice_a,"slice_b:",slice_b, "slice_c:", slice_c)

	fmt.Println(variadic_b([]uint{1,2,3,4}...), variadic_a([]string{"hey","i am","goody"}...), generateEvenNumbers()())
	
}


func variadic_a(array ...string) (result string){
	for _ ,value := range array{
		result += " " + value
	}

	return
}


func variadic_b(array ...uint) (sum uint){
	sum = uint(0)
	
	for _, value := range array{		
		sum += value
	}

	return 
}


func generateEvenNumbers() func() []int {
	slice_a := []int{}
	return func() []int {
		for i:=1; i < 10;i++{
			if i % 2 == 0 {
				slice_a = append(slice_a, i)
			}
		}

		return slice_a
	}
}