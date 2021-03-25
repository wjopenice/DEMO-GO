package main

import (
	"fmt"
	"time"
)

var sum = 0

func say(s string,c chan int){
	for i := 0; i < 100; i++ {
		c <- sum
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		fmt.Println(sum)
		sum ++
	}
}

func main() {
	c := make(chan int,2)
	c <- 1
	c <- 2
	go say("world",c)
	say("hello",c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
