package main

import (
	"fmt"
	// "runtime"
	"sync"
	"time"
)

const NumberUrl = 1000
func output(c chan int,name string){
	for i:= range c{
		time.Sleep(time.Second/100)
		fmt.Println("go routine ",name," is running: ",i)
	}
	wg.Done()
}
var wg sync.WaitGroup
func main() {
	
	in := make(chan int, NumberUrl)
	for i := 0; i < NumberUrl; i++ {
		in <- (i + 1)
	}
	close(in)
	
	wg.Add(5)
	go output(in,"1")
	go output(in,"2")
	go output(in,"3")
	go output(in,"4")
	go output(in,"5")
	// fmt.Println("number: ",runtime.NumGoroutine())
	wg.Wait()
}
