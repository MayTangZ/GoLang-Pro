package main

import "fmt"

func printArray1(myArray []int) {
	//引用传递

	//——表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)

	}

	myArray[0] = 100
}

func main2() {

	myArray := []int{1, 2, 3, 4, 5} //动态数组 slice

	fmt.Printf("myArray type is %T\n", myArray)
	printArray1(myArray)

	fmt.Println("====")

	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}
}
