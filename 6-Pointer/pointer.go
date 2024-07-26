package main

import "fmt"

func swap(a, b int) {
	var temp int

	temp = a
	a = b
	b = temp

}

func swap2(a, b *int) {

	var temp int
	temp = *a
	*a = *b
	*b = temp

	fmt.Println("a=", a, "b=", b)

}

func main() {
	var a int = 10
	var b int = 20

	//swap(a,b)
	swap2(&a, &b)
	fmt.Println("a=", a, "b=", b)

	var p *int
	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)

}
