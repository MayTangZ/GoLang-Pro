//package function

package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

//返回多个返回值匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	return 666, 777

}

//返回多个返回值，有形参名字
func foo3(a string, b int) (r1 int, r2 int) {

	fmt.Println("-------foo3-------")

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("-------foo4-------")

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func main1() {
	c := foo1("hello", 500)
	fmt.Println("c=", c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

	ret3, ret4 := foo3("haha", 333)
	fmt.Println("ret3=", ret3, "ret4=", ret4)

	ret5, ret6 := foo4("haha", 444)
	fmt.Println("ret5=", ret5, "ret6=", ret6)

}
