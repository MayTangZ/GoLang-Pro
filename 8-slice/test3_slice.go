package main

import "fmt"

func main3() {
	//声明slice是一个切片 ，并且初始化 ，默认值是 1 2 3 4 5 。长度是5
	//slice1 := []int{1, 2, 3, 4, 5}

	//声明slice是一个切片 ，同时给slice分配空间，5个空间，初始化值是0
	// var slice1 []int
	// slice1 = make([]int, 5)

	//var slice1 = make([]int, 5)

	slice1 := make([]int, 5)

	fmt.Printf("len= %d ,slice= %v\n", len(slice1), slice1)

}
