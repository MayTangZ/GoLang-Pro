package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called.....")
	fmt.Println(arg)

	//类型断言
	vaule, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is  string type ,value is", vaule)

		fmt.Printf("value type is %T\n", vaule)
	}
}

type Book1 struct {
	auth string
}

func main5() {
	abook := Book1{"Go Lang"}
	myFunc(abook)
	myFunc(100)
	myFunc("qqq")
}
