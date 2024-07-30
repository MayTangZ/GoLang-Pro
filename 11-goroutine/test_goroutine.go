package main

import (
	"fmt"
	"time"
)

// 主goroutine
func main() {
	// 匿名函数： 形参为空，返回值为空的一个函数
	/*
		go func() {
			defer fmt.Println("A.defer")

			func() {
				defer fmt.Println("B.defer")
				//退出当前goroutine
				runtime.Goexit()
				fmt.Println("B")
			}()
			fmt.Println("A")

		}()
	*/

	go func(a int, b int) bool {
		fmt.Println("a = ", a, "b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
