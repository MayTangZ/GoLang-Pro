package main

import (
	// 匿名包  不可以调用包中方法 但是可以用Initiate 方法

	"GolangStudy/5-Initiate/lib1"
	"fmt"
)

// import  lib1
// import  lib2

func main() {
	fmt.Println("Hello, World!")
	lib1.Lib1Test()
	//Lib2Test()

}
