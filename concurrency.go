package main

import (
	"fmt"
	"time"
)


func main(){
	c := make(chan string)
	go count("sheep", c)

	
	for msg := range c {
		fmt.Println(msg)
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		for {
			c1 <- "slept for 500 secs"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func(){
		for {
			c1 <- "slept for 2 secs"
			time.Sleep(time.Millisecond * 2000)
		}
	}()


	for{
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)

		}
	}
}

func count(thing string, c chan string){
	for i := 1; i <= 7; i++{
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}