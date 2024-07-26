/*
name
*/
package main

import "fmt"

//声明全局变量 :func1 ,2,3 is OK
var gA = 100
var gB = 200

//use fun4 声明全局变量
//gC:=200
//:= 只能够用在 函数体内声明

func main() {
	//function 1
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a =%T\n", a)

	//function 2

	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of b =%T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb =%s, type of bb=%T\n", bb, bb)

	//function 3

	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("type of c =%T\n", c)

	var cc = "abcd"
	fmt.Printf("cc =%s, type of cc=%T\n", cc, cc)

	//function 4
	e := 100
	fmt.Println("e=", e)
	fmt.Printf("type of e =%T\n", e)

	f := "abcd"
	fmt.Println("f=", f)
	fmt.Printf("type of f =%T\n", f)

	g := 3.14
	fmt.Println("g=", g)
	fmt.Printf("type of g =%T\n", g)

	//
	fmt.Println("gA=", gA, ", gB=", gB)
	//fmt.Println("gC=", gC)

	var xx, yy int = 100, 200
	fmt.Println("xx=", xx, ",yy=", yy)

	var kk, ll = 100, "qqqq"
	fmt.Println("kk=", kk, ",ll=", ll)

	//多变量声明
	var (
		vv int = 100
		jj     = true
	)

	fmt.Println("vv=", vv, ",jj=", jj)
}
