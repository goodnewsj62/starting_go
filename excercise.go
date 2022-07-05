package main

import "fmt"


func main(){
	fmt.Println(half_is_even(2))
	fmt.Println(half_is_even(1))
	fmt.Println(max([]int{1,28,6,3,4,1,5}...))
	get_odd_num := makeOddGenerator()
	fmt.Println(get_odd_num())
	fmt.Println(get_odd_num())
	fmt.Println(get_odd_num())
	
	fib_slice := make([]int,0,6)
	for i :=0; i < 6; i++{
		fib_slice = append(fib_slice, fibonacci(i))
	}

	fmt.Println(fib_slice)
}


func half_is_even(num int) (result bool){
	num = int(num / 2)
	result = false

	if float64(num % 2) == 0 {
		result = true
		return
	}
	return
}

func  max(args... int) int{
	max := 0

	for _, value := range args{
		if value > max{
			max = value
		}
	} 
	return max
}

func makeOddGenerator() func() uint{
	stored_value := uint(0)

	return func () uint{
		if stored_value == 0 {
			stored_value += 1
			return stored_value
		}
		stored_value += 2
		return stored_value
	}
}


func fibonacci(n int) int{
	if n < 2{
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}