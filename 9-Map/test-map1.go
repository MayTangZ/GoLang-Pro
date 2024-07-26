package main

import "fmt"

func main1() {
	// 声明mymap1 是一种map 类型 key是string ， value 是string
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")

	}
	//使用map前 ，需要先用make给map分配人、数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	myMap1["four"] = "C#"
	fmt.Println(myMap1)

	//method 2
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	//method 3
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python",
	}

	fmt.Println(myMap3)

}
