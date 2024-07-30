package main

import "fmt"

func main3() {
	c := make(chan int) //wu缓冲的 channel

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c) // if not close channel will pop fatal error: all goroutines are asleep - deadlock!
	}()

	/*
		for {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	// 使用 range 遍历 channel ，当 channel 关闭后会自动退出

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main 结束")

}
