package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i =%d\n ", i)
		time.Sleep(1 * time.Second)
	}

}

// 主goroutine
func main1() {
	//创建一个go 程去执行newTak（）流程
	go newTask()

	//fmt.Println("main goroutine exit")
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i =%d\n ", i)
		time.Sleep(1 * time.Second)
	}
}
