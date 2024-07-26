package main

import "fmt"

func printArray(myArray [4]int) {
	for index, value := range myArray {

		fmt.Println("index=", index, "value=", value)
	}

}

func main1() {
	fmt.Println("Hello World")
	//固定长度的数组
	var myArray [10]int
	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])

	}

	for index, value := range myArray2 {
		fmt.Println("index=", index, "value=", value)
	}

	//查看数组的数据类型

	fmt.Printf("myArray type is %T\n", myArray)
	fmt.Printf("myArray2 type is %T\n", myArray2)
	fmt.Printf("myArray3 type is %T\n", myArray3)
	printArray(myArray3)
}
