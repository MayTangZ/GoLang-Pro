package main

import "fmt"

//int 的一个别名
type myint int

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {

	book.auth = "666"

}

func changeBook2(book *Book) {
	book.auth = "777"

}
func main1() {

	//var a myint = 10
	var book1 Book

	book1.title = "Golang"
	book1.auth = "zhang3"
	changeBook(book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
