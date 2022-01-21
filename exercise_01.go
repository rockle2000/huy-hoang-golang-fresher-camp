package main

import (
	"fmt"
	"time"
)

const NumberUrl = 1000
func output(c chan int,name string){
	for i:= range c{
		time.Sleep(time.Second/100)
		fmt.Println("go routine ",name," is running: ",i)
	}
}
func main() {
	in := make(chan int, NumberUrl)
	for i := 0; i < NumberUrl; i++ {
		in <- (i + 1)
	}
	close(in)
	go output(in,"1")
	go output(in,"2")
	go output(in,"3")
	go output(in,"4")
	go output(in,"5")

	time.Sleep(time.Second*5)
}
