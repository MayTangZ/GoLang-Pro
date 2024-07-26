package lib1

import "fmt"

//当前Lib包提供的API
func Lib1Test() {
	fmt.Println("lib1.Lib1Test()")
}

func init() {
	fmt.Println("lib1.init()")

}
