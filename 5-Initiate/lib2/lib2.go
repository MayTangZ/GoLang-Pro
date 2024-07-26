package lib2

import "fmt"

//当前Lib2包提供的API
func Lib2Test() {
	fmt.Println("lib2.Lib1Test()")
}

func init() {
	fmt.Println("lib2.init()")

}
