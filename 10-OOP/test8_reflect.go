package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {

	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)

}

func DoFiledAndMethod(input interface{}) {

	inputType := reflect.TypeOf(input)
	//
	fmt.Println("inputType is :", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)
	//通过type ，获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fmt.Printf("%s : %v\n", field.Name, field.Type)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s : %v\n", field.Name, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)

	}

}

func reflectNum(arg interface{}) {

	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))

}

func main8() {
	var num float64 = 1.234
	reflectNum(num)

	user1 := User{1, "zhangsan", 18}
	DoFiledAndMethod(user1)

}
