package main

import "fmt"

func main1() {
	//define a channel

	ch := make(chan int)

	go func() {
		defer fmt.Println("goroutine finished")
		fmt.Println("goroutine started")
		ch <- 666 //send 666 to ch

	}()

	num := <-ch //receive from ch
	fmt.Println("num = ", num)
	fmt.Println("main goroutine finished")

}
