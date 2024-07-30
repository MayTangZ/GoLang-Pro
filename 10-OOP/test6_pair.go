package main

import (
	"fmt"
	"io"
	"os"
)

func main6() {
	var a string
	// pair<staictype:string ,value :"Hello">
	a = "World"

	// pair<type:string ,value :"Hello">
	var allType interface{}
	allType = a

	value, _ := allType.(string)
	fmt.Println(value)

	//tty:pair<type:*os.File ,value :"/dev/tty">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}

	//r:pair<type: ,value :>
	var r io.Reader
	//r:pair<type:*os.File ,value :"/dev/tty">
	r = tty

	//w:pair<type: ,value :>
	var w io.Writer
	//w:pair<type:*os.File ,value :"/dev/tty">
	w = r.(io.Writer) //r强制转换成w

	w.Write([]byte("hello this is a tests"))

}
