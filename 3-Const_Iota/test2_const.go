package main

import "fmt"

//const 定义枚举类型
const (
	//可以添加iota ， 每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING  = 10 * iota //iota=0
	SHANGHAI             //iota=1
	SHENZHEN             //iota=2
)

const (
	a, b = iota + 1, iota + 2 //iota=0
	c, d                      //iota=1   c= iota + 1 d=iota + 2
	e, f                      //iota=2
	g, h = iota * 2, iota * 3 //iota=3   g=iota * 2 h= iota * 3
	i, k                      //iota=4
)

func main() {

	//只读属性
	const length int = 10

	fmt.Println("length=", length)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHENZHEN=", SHENZHEN)
	fmt.Println("a=", a, "，b=", b, "，c=", c, "，d=", d, "，e=", e, "，f=", f, "，g=", g, "，h=", h, "，i=", i, "，k=", k)

}
