package main

import (
	"fmt"
	"time"
)

func hello(ch chan int){
	time.Sleep(5* time.Second)
	ch <- 42
	close(ch)
}

func main(){
	ch := make(chan int)


	go hello(ch)
	for {
	select{
		case _,ok := <-ch:
			if !ok {
				return 
			}
		default:
			fmt.Println("hello")
	}
}

}
