package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book2 struct {
}

func (this *Book2) ReadBook() {
	fmt.Println("Read a Book")

}

func (this *Book2) WriteBook() {
	fmt.Println("Write a Book")
}

func main7() {

	//b:pair<type:Book2 ,value :"&Book2{}" book{}地址>
	b := &Book2{}
	//r:pair<type: ,value :>
	var r Reader

	//r:pair<type:Book2 ,value :"&Book2{}" book{}地址>
	r = b
	r.ReadBook()

	var w Writer
	//w:pair<type:Book2 ,value :"&Book2{}" book{}地址>

	//断言成功： 1. 两种类型是相同的 2. 断言接口值（也就是r）实现了目标接口的所有方法所以断言才能成功
	//y:=x.(T) T为非接口类型时，则T必须实现x
	w = r.(Writer)

	w.WriteBook()

}
