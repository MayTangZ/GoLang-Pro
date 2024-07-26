package main

import "fmt"

func func1() {
	fmt.Println("func1:Hello World 1")

}

func func2() {
	fmt.Println("func2:Hello World 2")

}

func func3() {
	fmt.Println("func3:Hello World 3")

}

func deferfunc() int {
	fmt.Println("Call deferfunc:Hello World deferfunc")
	return 0

}

func returnfunc() int {
	fmt.Println("Call returnfunc:Hello World returnfunc")
	return 0

}

func returnAndDefer() int {
	fmt.Println("returnAndDefer:Hello World 4")
	defer deferfunc()
	return returnfunc()

}

func main() {

	// //写入defer关键字
	// defer fmt.Println("main end1")
	// defer fmt.Println("main end2")

	// fmt.Println("main:Hello World 1")

	// fmt.Println("main:Hello World 2")

	// defer func1()
	// defer func2()
	// defer func3()

	returnAndDefer()
}
